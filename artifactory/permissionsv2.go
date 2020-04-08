// Copyright (c) 2016 John E. Vincent
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Copyright (c) 2018 Target Brands, Inc.

package artifactory

import (
	"fmt"
	"net/http"
)

// PermissionsServiceV2 handles communication with the permissions related
// methods of the Artifactory API v2.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API+V2
type PermissionsServiceV2 service

// Exists validates if the specific permission target exists.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API+V2#ArtifactoryRESTAPIV2-PermissionTargetexistencecheck
func (s *PermissionsServiceV2) Exists(target string) (bool, error) {
	u := fmt.Sprintf("/api/v2/security/permissions/%s", target)
	resp, err := s.client.Call("HEAD", u, nil, nil)
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// PermissionDetails represents the information about the repo, build, or releasebundle within the permission target.
type PermissionDetails struct {
	IncludePatterns *[]string `json:"include-patterns,omitempty"`
	ExcludePatterns *[]string `json:"exclude-patterns,omitempty"`
	Repositories    *[]string `json:"repositories,omitempty"`
	Actions         *Actions  `json:"actions,omitempty"`
}

// Actions represents user and group permissions.
type Actions struct {
	Users  *map[string][]string `json:"users,omitempty"`
	Groups *map[string][]string `json:"groups,omitempty"`
}

// PermissionTargetV2 represents a v2 permission target.
//
// Docs: https://www.jfrog.com/confluence/display/JFROG/Security+Configuration+JSON#SecurityConfigurationJSON-application/vnd.org.jfrog.artifactory.security.PermissionTargetV2+json
type PermissionTargetV2 struct {
	Name          *string            `json:"name,omitempty"`
	Repo          *PermissionDetails `json:"repo,omitempty"`
	Build         *PermissionDetails `json:"build,omitempty"`
	ReleaseBundle *PermissionDetails `json:"releaseBundle,omitempty"`
}

func (p PermissionTargetV2) String() string {
	return Stringify(p)
}

// Update creates a new permission target or replaces an existing permission target.
//
// Docs: https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API+V2#ArtifactoryRESTAPIV2-UpdatePermissionTarget
func (s *PermissionsServiceV2) Update(target *PermissionTargetV2) (*string, *Response, error) {
	u := fmt.Sprintf("/api/v2/security/permissions/%s", *target.Name)
	v := new(string)

	resp, err := s.client.Call("PUT", u, target, v)
	return v, resp, err
}

// Get returns the provided permission target.
//
// Docs: https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API+V2#ArtifactoryRESTAPIV2-GetPermissionTargetDetails
func (s *PermissionsServiceV2) Get(target string) (*string, *Response, error) {
	u := fmt.Sprintf("/api/v2/security/permissions/%s", target)
	v := new(string)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}
