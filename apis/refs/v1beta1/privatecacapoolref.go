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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PrivateCACAPoolRef struct {
	// +required
	/* The self-link of an existing PrivateCA CA Pool, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
}

type PrivateCACAPool struct {
	ProjectID  string
	Location   string
	ResourceID string
}

func (r *PrivateCACAPool) String() string {
	return "projects/" + r.ProjectID + "/locations/" + r.Location + "/caPools/" + r.ResourceID
}

func ResolvePrivateCACAPoolRef(ctx context.Context, reader client.Reader, obj client.Object, ref *PrivateCACAPoolRef) (*PrivateCACAPool, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		// External must be in form `projects/<projectID>/locations/<location>/caPools/<ResourceID>`.
		// see https://cloud.google.com/certificate-authority-service/docs/reference/rest/v1/projects.locations.caPools/get
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "caPools" {
			return &PrivateCACAPool{
				ProjectID:  tokens[1],
				Location:   tokens[3],
				ResourceID: tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of PrivateCACAPool external=%q was not known (use projects/<projectID>/locations/<location>/caPools/<ResourceID>)", ref.External)
	}

	return nil, fmt.Errorf("must specify external on PrivateCACAPoolRef")
}
