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

package v1alpha1

// +kcc:proto=google.cloud.securesourcemanager.v1.Instance
type Instance struct {
	// Optional. A unique identifier for an instance. The name should be of the
	//  format:
	//  `projects/{project_number}/locations/{location_id}/instances/{instance_id}`
	//
	//  `project_number`: Maps to a unique int64 id assigned to each project.
	//
	//  `location_id`: Refers to the region where the instance will be deployed.
	//  Since Secure Source Manager is a regional service, it must be one of the
	//  valid GCP regions.
	//
	//  `instance_id`: User provided name for the instance, must be unique for a
	//  project_number and location_id combination.
	Name *string `json:"name,omitempty"`

	// Output only. Create timestamp.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update timestamp.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Labels as key value pairs.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Private settings for private instance.
	PrivateConfig *Instance_PrivateConfig `json:"privateConfig,omitempty"`

	// Output only. Current state of the instance.
	State *string `json:"state,omitempty"`

	// Output only. An optional field providing information about the current
	//  instance state.
	StateNote *string `json:"stateNote,omitempty"`

	// Optional. Immutable. Customer-managed encryption key name, in the format
	//  projects/*/locations/*/keyRings/*/cryptoKeys/*.
	KmsKey *string `json:"kmsKey,omitempty"`

	// Output only. A list of hostnames for this instance.
	HostConfig *Instance_HostConfig `json:"hostConfig,omitempty"`
}

// +kcc:proto=google.cloud.securesourcemanager.v1.Instance.HostConfig
type Instance_HostConfig struct {
	// Output only. HTML hostname.
	HTML *string `json:"html,omitempty"`

	// Output only. API hostname. This is the hostname to use for **Host: Data
	//  Plane** endpoints.
	Api *string `json:"api,omitempty"`

	// Output only. Git HTTP hostname.
	GitHTTP *string `json:"gitHTTP,omitempty"`

	// Output only. Git SSH hostname.
	GitSSH *string `json:"gitSSH,omitempty"`
}

// +kcc:proto=google.cloud.securesourcemanager.v1.Instance.PrivateConfig
type Instance_PrivateConfig struct {
	// Required. Immutable. Indicate if it's private instance.
	IsPrivate *bool `json:"isPrivate,omitempty"`

	// Required. Immutable. CA pool resource, resource must in the format of
	//  `projects/{project}/locations/{location}/caPools/{ca_pool}`.
	CaPool *string `json:"caPool,omitempty"`

	// Output only. Service Attachment for HTTP, resource is in the format of
	//  `projects/{project}/regions/{region}/serviceAttachments/{service_attachment}`.
	HTTPServiceAttachment *string `json:"httpServiceAttachment,omitempty"`

	// Output only. Service Attachment for SSH, resource is in the format of
	//  `projects/{project}/regions/{region}/serviceAttachments/{service_attachment}`.
	SSHServiceAttachment *string `json:"sshServiceAttachment,omitempty"`
}
