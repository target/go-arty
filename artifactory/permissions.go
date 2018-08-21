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
)

// PermissionsService handles communication with the permissions related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SECURITY
type PermissionsService service

// PermissionTarget represents a permission target in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Security+Configuration+JSON#SecurityConfigurationJSON-application/vnd.org.jfrog.artifactory.security.PermissionTarget+json
type PermissionTarget struct {
	Name            *string     `json:"name,omitempty"`
	URI             *string     `json:"uri,omitempty"`
	IncludesPattern *string     `json:"includesPattern,omitempty"`
	ExcludesPattern *string     `json:"excludesPattern,omitempty"`
	Repositories    *[]string   `json:"repositories,omitempty"`
	Principals      *Principals `json:"principals,omitempty"`
}

func (p PermissionTarget) String() string {
	return Stringify(p)
}

// Principals represents user and group permissions in Artifactory.
type Principals struct {
	Users  *map[string][]string `json:"users,omitempty"`
	Groups *map[string][]string `json:"groups,omitempty"`
}

// GetAll returns a list of all permission targets.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetPermissionTargets
func (s *PermissionsService) GetAll() (*[]PermissionTarget, *Response, error) {
	u := fmt.Sprintf("/api/security/permissions")
	v := new([]PermissionTarget)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns the provided permission target.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetPermissionTargetDetails
func (s *PermissionsService) Get(target string) (*PermissionTarget, *Response, error) {
	u := fmt.Sprintf("/api/security/permissions/%s", target)
	v := new(PermissionTarget)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Create constructs a permission target with the provided details.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplacePermissionTarget
func (s *PermissionsService) Create(target *PermissionTarget) (*string, *Response, error) {
	u := fmt.Sprintf("/api/security/permissions/%s", *target.Name)
	v := new(string)

	resp, err := s.client.Call("PUT", u, target, v)
	return v, resp, err
}

// Update modifies a permission target with the provided details.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplacePermissionTarget
func (s *PermissionsService) Update(target *PermissionTarget) (*string, *Response, error) {
	u := fmt.Sprintf("/api/security/permissions/%s", *target.Name)
	v := new(string)

	resp, err := s.client.Call("PUT", u, target, v)
	return v, resp, err
}

// Delete removes the provided permission target.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-DeletePermissionTarget
func (s *PermissionsService) Delete(target string) (*string, *Response, error) {
	u := fmt.Sprintf("/api/security/permissions/%s", target)
	v := new(string)

	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}
