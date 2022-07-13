//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Kubesphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddPaths) DeepCopyInto(out *AddPaths) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(AddPathsConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddPaths.
func (in *AddPaths) DeepCopy() *AddPaths {
	if in == nil {
		return nil
	}
	out := new(AddPaths)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddPathsConfig) DeepCopyInto(out *AddPathsConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddPathsConfig.
func (in *AddPathsConfig) DeepCopy() *AddPathsConfig {
	if in == nil {
		return nil
	}
	out := new(AddPathsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AfiSafi) DeepCopyInto(out *AfiSafi) {
	*out = *in
	if in.MpGracefulRestart != nil {
		in, out := &in.MpGracefulRestart, &out.MpGracefulRestart
		*out = new(MpGracefulRestart)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(AfiSafiConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AddPaths != nil {
		in, out := &in.AddPaths, &out.AddPaths
		*out = new(AddPaths)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AfiSafi.
func (in *AfiSafi) DeepCopy() *AfiSafi {
	if in == nil {
		return nil
	}
	out := new(AfiSafi)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AfiSafiConfig) DeepCopyInto(out *AfiSafiConfig) {
	*out = *in
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(Family)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AfiSafiConfig.
func (in *AfiSafiConfig) DeepCopy() *AfiSafiConfig {
	if in == nil {
		return nil
	}
	out := new(AfiSafiConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BgpConf) DeepCopyInto(out *BgpConf) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BgpConf.
func (in *BgpConf) DeepCopy() *BgpConf {
	if in == nil {
		return nil
	}
	out := new(BgpConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BgpConf) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BgpConfList) DeepCopyInto(out *BgpConfList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BgpConf, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BgpConfList.
func (in *BgpConfList) DeepCopy() *BgpConfList {
	if in == nil {
		return nil
	}
	out := new(BgpConfList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BgpConfList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BgpConfSpec) DeepCopyInto(out *BgpConfSpec) {
	*out = *in
	if in.AsPerRack != nil {
		in, out := &in.AsPerRack, &out.AsPerRack
		*out = make(map[string]uint32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ListenAddresses != nil {
		in, out := &in.ListenAddresses, &out.ListenAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Families != nil {
		in, out := &in.Families, &out.Families
		*out = make([]uint32, len(*in))
		copy(*out, *in)
	}
	if in.GracefulRestart != nil {
		in, out := &in.GracefulRestart, &out.GracefulRestart
		*out = new(GracefulRestart)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BgpConfSpec.
func (in *BgpConfSpec) DeepCopy() *BgpConfSpec {
	if in == nil {
		return nil
	}
	out := new(BgpConfSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BgpConfStatus) DeepCopyInto(out *BgpConfStatus) {
	*out = *in
	if in.NodesConfStatus != nil {
		in, out := &in.NodesConfStatus, &out.NodesConfStatus
		*out = make(map[string]NodeConfStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BgpConfStatus.
func (in *BgpConfStatus) DeepCopy() *BgpConfStatus {
	if in == nil {
		return nil
	}
	out := new(BgpConfStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BgpPeer) DeepCopyInto(out *BgpPeer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BgpPeer.
func (in *BgpPeer) DeepCopy() *BgpPeer {
	if in == nil {
		return nil
	}
	out := new(BgpPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BgpPeer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BgpPeerList) DeepCopyInto(out *BgpPeerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BgpPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BgpPeerList.
func (in *BgpPeerList) DeepCopy() *BgpPeerList {
	if in == nil {
		return nil
	}
	out := new(BgpPeerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BgpPeerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BgpPeerSpec) DeepCopyInto(out *BgpPeerSpec) {
	*out = *in
	if in.Conf != nil {
		in, out := &in.Conf, &out.Conf
		*out = new(PeerConf)
		**out = **in
	}
	if in.EbgpMultihop != nil {
		in, out := &in.EbgpMultihop, &out.EbgpMultihop
		*out = new(EbgpMultihop)
		**out = **in
	}
	if in.Timers != nil {
		in, out := &in.Timers, &out.Timers
		*out = new(Timers)
		(*in).DeepCopyInto(*out)
	}
	if in.Transport != nil {
		in, out := &in.Transport, &out.Transport
		*out = new(Transport)
		**out = **in
	}
	if in.GracefulRestart != nil {
		in, out := &in.GracefulRestart, &out.GracefulRestart
		*out = new(GracefulRestart)
		**out = **in
	}
	if in.AfiSafis != nil {
		in, out := &in.AfiSafis, &out.AfiSafis
		*out = make([]*AfiSafi, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AfiSafi)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BgpPeerSpec.
func (in *BgpPeerSpec) DeepCopy() *BgpPeerSpec {
	if in == nil {
		return nil
	}
	out := new(BgpPeerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BgpPeerStatus) DeepCopyInto(out *BgpPeerStatus) {
	*out = *in
	if in.NodesPeerStatus != nil {
		in, out := &in.NodesPeerStatus, &out.NodesPeerStatus
		*out = make(map[string]NodePeerStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BgpPeerStatus.
func (in *BgpPeerStatus) DeepCopy() *BgpPeerStatus {
	if in == nil {
		return nil
	}
	out := new(BgpPeerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EbgpMultihop) DeepCopyInto(out *EbgpMultihop) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbgpMultihop.
func (in *EbgpMultihop) DeepCopy() *EbgpMultihop {
	if in == nil {
		return nil
	}
	out := new(EbgpMultihop)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Eip) DeepCopyInto(out *Eip) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Eip.
func (in *Eip) DeepCopy() *Eip {
	if in == nil {
		return nil
	}
	out := new(Eip)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Eip) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EipList) DeepCopyInto(out *EipList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Eip, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EipList.
func (in *EipList) DeepCopy() *EipList {
	if in == nil {
		return nil
	}
	out := new(EipList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EipList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EipSpec) DeepCopyInto(out *EipSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EipSpec.
func (in *EipSpec) DeepCopy() *EipSpec {
	if in == nil {
		return nil
	}
	out := new(EipSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EipStatus) DeepCopyInto(out *EipStatus) {
	*out = *in
	if in.Used != nil {
		in, out := &in.Used, &out.Used
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EipStatus.
func (in *EipStatus) DeepCopy() *EipStatus {
	if in == nil {
		return nil
	}
	out := new(EipStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Family) DeepCopyInto(out *Family) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Family.
func (in *Family) DeepCopy() *Family {
	if in == nil {
		return nil
	}
	out := new(Family)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GracefulRestart) DeepCopyInto(out *GracefulRestart) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GracefulRestart.
func (in *GracefulRestart) DeepCopy() *GracefulRestart {
	if in == nil {
		return nil
	}
	out := new(GracefulRestart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Message) DeepCopyInto(out *Message) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Message.
func (in *Message) DeepCopy() *Message {
	if in == nil {
		return nil
	}
	out := new(Message)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Messages) DeepCopyInto(out *Messages) {
	*out = *in
	if in.Received != nil {
		in, out := &in.Received, &out.Received
		*out = new(Message)
		**out = **in
	}
	if in.Sent != nil {
		in, out := &in.Sent, &out.Sent
		*out = new(Message)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Messages.
func (in *Messages) DeepCopy() *Messages {
	if in == nil {
		return nil
	}
	out := new(Messages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MpGracefulRestart) DeepCopyInto(out *MpGracefulRestart) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(MpGracefulRestartConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MpGracefulRestart.
func (in *MpGracefulRestart) DeepCopy() *MpGracefulRestart {
	if in == nil {
		return nil
	}
	out := new(MpGracefulRestart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MpGracefulRestartConfig) DeepCopyInto(out *MpGracefulRestartConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MpGracefulRestartConfig.
func (in *MpGracefulRestartConfig) DeepCopy() *MpGracefulRestartConfig {
	if in == nil {
		return nil
	}
	out := new(MpGracefulRestartConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfStatus) DeepCopyInto(out *NodeConfStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfStatus.
func (in *NodeConfStatus) DeepCopy() *NodeConfStatus {
	if in == nil {
		return nil
	}
	out := new(NodeConfStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePeerStatus) DeepCopyInto(out *NodePeerStatus) {
	*out = *in
	in.PeerState.DeepCopyInto(&out.PeerState)
	out.TimersState = in.TimersState
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePeerStatus.
func (in *NodePeerStatus) DeepCopy() *NodePeerStatus {
	if in == nil {
		return nil
	}
	out := new(NodePeerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerConf) DeepCopyInto(out *PeerConf) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerConf.
func (in *PeerConf) DeepCopy() *PeerConf {
	if in == nil {
		return nil
	}
	out := new(PeerConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerState) DeepCopyInto(out *PeerState) {
	*out = *in
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = new(Messages)
		(*in).DeepCopyInto(*out)
	}
	if in.Queues != nil {
		in, out := &in.Queues, &out.Queues
		*out = new(Queues)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerState.
func (in *PeerState) DeepCopy() *PeerState {
	if in == nil {
		return nil
	}
	out := new(PeerState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Queues) DeepCopyInto(out *Queues) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queues.
func (in *Queues) DeepCopy() *Queues {
	if in == nil {
		return nil
	}
	out := new(Queues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timers) DeepCopyInto(out *Timers) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(TimersConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timers.
func (in *Timers) DeepCopy() *Timers {
	if in == nil {
		return nil
	}
	out := new(Timers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimersConfig) DeepCopyInto(out *TimersConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimersConfig.
func (in *TimersConfig) DeepCopy() *TimersConfig {
	if in == nil {
		return nil
	}
	out := new(TimersConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimersState) DeepCopyInto(out *TimersState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimersState.
func (in *TimersState) DeepCopy() *TimersState {
	if in == nil {
		return nil
	}
	out := new(TimersState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transport) DeepCopyInto(out *Transport) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transport.
func (in *Transport) DeepCopy() *Transport {
	if in == nil {
		return nil
	}
	out := new(Transport)
	in.DeepCopyInto(out)
	return out
}
