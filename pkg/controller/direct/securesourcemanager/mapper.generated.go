// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package securesourcemanager

import (
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Labels = in.Labels
	out.PrivateConfig = Instance_PrivateConfig_FromProto(mapCtx, in.GetPrivateConfig())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateNote = direct.Enum_FromProto(mapCtx, in.GetStateNote())
	out.KmsKey = direct.LazyPtr(in.GetKmsKey())
	out.HostConfig = Instance_HostConfig_FromProto(mapCtx, in.GetHostConfig())
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Labels = in.Labels
	out.PrivateConfig = Instance_PrivateConfig_ToProto(mapCtx, in.PrivateConfig)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateNote = direct.Enum_ToProto[pb.Instance_StateNote](mapCtx, in.StateNote)
	out.KmsKey = direct.ValueOf(in.KmsKey)
	out.HostConfig = Instance_HostConfig_ToProto(mapCtx, in.HostConfig)
	return out
}
func Instance_HostConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_HostConfig) *krm.Instance_HostConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_HostConfig{}
	out.HTML = direct.LazyPtr(in.GetHtml())
	out.Api = direct.LazyPtr(in.GetApi())
	out.GitHTTP = direct.LazyPtr(in.GetGitHttp())
	out.GitSSH = direct.LazyPtr(in.GetGitSsh())
	return out
}
func Instance_HostConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_HostConfig) *pb.Instance_HostConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_HostConfig{}
	out.Html = direct.ValueOf(in.HTML)
	out.Api = direct.ValueOf(in.Api)
	out.GitHttp = direct.ValueOf(in.GitHTTP)
	out.GitSsh = direct.ValueOf(in.GitSSH)
	return out
}
func Instance_PrivateConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PrivateConfig) *krm.Instance_PrivateConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PrivateConfig{}
	out.IsPrivate = direct.LazyPtr(in.GetIsPrivate())
	out.CaPool = direct.LazyPtr(in.GetCaPool())
	out.HTTPServiceAttachment = direct.LazyPtr(in.GetHttpServiceAttachment())
	out.SSHServiceAttachment = direct.LazyPtr(in.GetSshServiceAttachment())
	return out
}
func Instance_PrivateConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PrivateConfig) *pb.Instance_PrivateConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PrivateConfig{}
	out.IsPrivate = direct.ValueOf(in.IsPrivate)
	out.CaPool = direct.ValueOf(in.CaPool)
	out.HttpServiceAttachment = direct.ValueOf(in.HTTPServiceAttachment)
	out.SshServiceAttachment = direct.ValueOf(in.SSHServiceAttachment)
	return out
}
