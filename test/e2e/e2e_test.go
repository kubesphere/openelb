package e2e_test

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	devopsv1alpha1 "github.com/kubesphere/porter/pkg/apis/network/v1alpha1"
	"github.com/kubesphere/porter/test/e2eutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
)

var _ = Describe("E2e", func() {
	It("Should work well when using samples", func() {
		eip := &devopsv1alpha1.EIP{}
		reader, err := os.Open(workspace + "/config/samples/network_v1alpha1_eip.yaml")
		Expect(err).NotTo(HaveOccurred(), "Cannot read sample yamls")
		err = yaml.NewYAMLOrJSONDecoder(reader, 10).Decode(eip)
		Expect(err).NotTo(HaveOccurred(), "Cannot unmarshal yamls")
		if eip.Namespace == "" {
			eip.Namespace = "default"
		}
		err = testClient.Create(context.TODO(), eip)
		Expect(err).NotTo(HaveOccurred())
		defer testClient.Delete(context.TODO(), eip)

		//apply service
		cmd := exec.Command("kubectl", "apply", "-f", workspace+"/config/samples/service.yaml")
		Expect(cmd.Run()).ShouldNot(HaveOccurred())
		defer func() {
			cmd := exec.Command("kubectl", "delete", "-f", workspace+"/config/samples/service.yaml")
			Expect(cmd.Run()).ShouldNot(HaveOccurred())
		}()
		serviceTypes := types.NamespacedName{Namespace: "default", Name: "mylbapp"}
		//Service should get its eip
		Eventually(func() error {
			service := &corev1.Service{}
			err := testClient.Get(context.TODO(), serviceTypes, service)
			if err != nil {
				return err
			}
			if len(service.Spec.ExternalIPs) > 0 && len(service.Status.LoadBalancer.Ingress) > 0 {
				if service.Spec.ExternalIPs[0] == eip.Spec.Address && service.Status.LoadBalancer.Ingress[0].IP == eip.Spec.Address {
					return nil
				}
			}
			return fmt.Errorf("Failed")
		}, 2*time.Minute, time.Second).Should(Succeed())
		//check route in bird
		bird_ip := os.Getenv("BIRD_IP")
		if bird_ip != "" {
			session, err := e2eutil.Connect("root", "", bird_ip, e2eutil.GetDefaultPrivateKeyFile(), 22, nil)
			Expect(err).NotTo(HaveOccurred(), "Connect Bird using private key FAILED")
			defer session.Close()
			stdinBuf, err := session.StdinPipe()
			var outbt, errbt bytes.Buffer
			session.Stdout = &outbt
			session.Stderr = &errbt
			err = session.Shell()
			Expect(err).ShouldNot(HaveOccurred(), "Failed to start ssh shell")
			Eventually(func() error {
				stdinBuf.Write([]byte("ip route\n"))
				ips, err := e2eutil.GetServiceNodesIP(testClient, serviceTypes.Namespace, serviceTypes.Name)
				if err != nil {
					fmt.Fprintln(GinkgoWriter, "Get service IPs Failed")
				}
				s := outbt.String() + errbt.String()
				for _, ip := range ips {
					if !strings.Contains(s, fmt.Sprintf("nexthop via %s", ip)) {
						return fmt.Errorf("No routes in Brid")
					}
				}
				if strings.Contains(s, fmt.Sprintf("%s  proto bird", eip.Spec.Address)) {
					return nil
				} else {
					return fmt.Errorf("No routes in Brid")
				}
			}, 30*time.Second, 2*time.Second).Should(Succeed())
		}
	})
	//install eip
})
