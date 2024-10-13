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

func SecureSourceManagerInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func SecureSourceManagerInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	return out
}

func SecureSourceManagerInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: PrivateConfig
	// MISSING: State
	// MISSING: StateNote
	// MISSING: KmsKey
	// MISSING: HostConfig
	return out
}
func SecureSourceManagerInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: PrivateConfig
	// MISSING: State
	// MISSING: StateNote
	// MISSING: KmsKey
	// MISSING: HostConfig
	return out
}
