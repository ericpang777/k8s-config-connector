
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

package 

import (
	"fmt"
	"strings"
)

// The Identifier for ConfigConnector to track the SecureSourceManagerInstance resource from the GCP service.
type SecureSourceManagerInstanceIdentity struct {
	Parent *parent
	Instance string
}

type parent struct {
	Project string
	Location string
}

func (p *parent) String() string{
	return fmt.Sprintf("projects/%s/locations/%s", p.Project, p.Location)
}

// FullyQualifiedName returns both parent and resource ID in the full url format. 
func (c *SecureSourceManagerInstanceIdentity) FullyQualifiedName() string {
	// TODO(user): Edit the URL path
	return fmt.Sprintf("%s/instances/%s", c.Parent, c.Instance)
}

// AsExternalRef builds a externalRef from a SecureSourceManagerInstance
func (c *SecureSourceManagerInstanceIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a SecureSourceManagerInstanceIdentity from a `status.externalRef`
func asID(externalRef string) (*SecureSourceManagerInstanceIdentity, error) {
	// TODO(user): Build resource identity from external reference
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	// TODO(user): Confirm the format of your resources, and verify it like the example below
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "instances" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/instances/<Instance>, got %s",
			serviceDomain, externalRef)
	}
	return &SecureSourceManagerInstanceIdentity{
		Parent: &parent{Project: tokens[1], Location: tokens[3]},
		Instance: tokens[5],
	}, nil
}

// BuildID builds the ID for ConfigConnector to track the SecureSourceManagerInstance resource from the GCP service.
func BuildID(project, location, resourceID string) *SecureSourceManagerInstanceIdentity {
	// TODO(user): Build resource identity from resource components, i.e. project, location, resource id
	return &SecureSourceManagerInstanceIdentity{
		Parent: &parent{Project: project, Location: location},
		Instance: resourceID,
	}
}
