package app

import (
	"flag"
	"fmt"
	networkv1alpha2 "github.com/kubesphere/porter/api/v1alpha2"
	"github.com/kubesphere/porter/pkg/leader-elector"
	clientset "k8s.io/client-go/kubernetes"
	"net/http"
	"os"

	"github.com/kubesphere/porter/cmd/manager/app/options"
	"github.com/kubesphere/porter/pkg/constant"
	"github.com/kubesphere/porter/pkg/controllers/bgp"
	"github.com/kubesphere/porter/pkg/controllers/ipam"
	"github.com/kubesphere/porter/pkg/controllers/lb"
	"github.com/kubesphere/porter/pkg/log"
	"github.com/kubesphere/porter/pkg/manager"
	"github.com/kubesphere/porter/pkg/speaker"
	bgpd "github.com/kubesphere/porter/pkg/speaker/bgp"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/util/term"
	cliflag "k8s.io/component-base/cli/flag"
	ctrl "sigs.k8s.io/controller-runtime"
)

func NewPorterManagerCommand() *cobra.Command {
	s := options.NewPorterManagerOptions()

	cmd := &cobra.Command{
		Use:  "porter-manager",
		Long: `The porter manager is a daemon that `,
		Run: func(cmd *cobra.Command, args []string) {
			if errs := s.Validate(); len(errs) != 0 {
				fmt.Fprintf(os.Stderr, "%v\n", utilerrors.NewAggregate(errs))
				os.Exit(1)
			}

			if err := Run(s); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}

	fs := cmd.Flags()

	namedFlagSets := s.Flags()
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	fs.AddFlagSet(pflag.CommandLine)

	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStderr(), namedFlagSets, cols)
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
	})

	return cmd
}

func serveReadinessHandler(w http.ResponseWriter, r *http.Request) {
	if readinessProbe {
		w.WriteHeader(200)
		w.Write([]byte("Ready"))
	} else {
		w.WriteHeader(500)
		w.Write([]byte("Not Ready"))
	}
}

func Run(c *options.PorterManagerOptions) error {
	log.InitLog(c.LogOptions)

	setupLog := ctrl.Log.WithName("setup")

	mgr, err := manager.NewManager(ctrl.GetConfigOrDie(), c.GenericOptions)
	if err != nil {
		setupLog.Error(err, "unable to new manager")
		return err
	}

	bgpServer := bgpd.NewGoBgpd(c.Bgp)
	err = speaker.RegisteSpeaker(constant.PorterProtocolBGP, bgpServer)
	if err != nil {
		setupLog.Error(err, "unable to register bgp speaker")
		return err
	}

	// Setup all Controllers
	err = ipam.SetupIPAM(mgr)
	if err != nil {
		setupLog.Error(err, "unable to setup ipam")
		return err
	}
	networkv1alpha2.Eip{}.SetupWebhookWithManager(mgr)

	err = bgp.SetupBgpConfReconciler(bgpServer, mgr)
	if err != nil {
		setupLog.Error(err, "unable to setup bgpconf")
	}

	err = bgp.SetupBgpPeerReconciler(bgpServer, mgr)
	if err != nil {
		setupLog.Error(err, "unable to setup bgppeer")
	}

	if err = lb.SetupServiceReconciler(mgr); err != nil {
		setupLog.Error(err, "unable to setup lb controller")
		return err
	}

	k8sClient := clientset.NewForConfigOrDie(ctrl.GetConfigOrDie())
	leader.LeaderElector(k8sClient)

	serverMuxA := http.NewServeMux()
	serverMuxA.HandleFunc("/hello", serveReadinessHandler)
	go func() {
		err := http.ListenAndServe(c.ReadinessAddr, serverMuxA)
		if err != nil {
			setupLog.Error(err, "Failed to start readiness probe")
			os.Exit(1)
		}
	}()

	readinessProbe = true
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "unable to run the manager")
		return err
	}

	return nil
}

var (
	readinessProbe bool
)
