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

// GroupsService handles communication with the group related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SECURITY
type GroupsService service

// Group represents a group in Artifactory.
//
// Doc: https://www.jfrog.com/confluence/display/RTF/Security+Configuration+JSON#SecurityConfigurationJSON-application/vnd.org.jfrog.artifactory.security.Group+json
type Group struct {
	Name            *string   `json:"name,omitempty"`
	URI             *string   `json:"uri,omitempty"`
	Description     *string   `json:"description,omitempty"`
	AutoJoin        *bool     `json:"autoJoin,omitempty"`
	AdminPrivileges *bool     `json:"adminPrivileges,omitempty"`
	Realm           *string   `json:"realm,omitempty"`
	RealmAttributes *string   `json:"realmAttributes,omitempty"`
	UserNames       *[]string `json:"userNames,omitempty"`
}

func (g Group) String() string {
	return Stringify(g)
}

// GetAll returns a list of all groups.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetGroups
func (s *GroupsService) GetAll() (*[]Group, *Response, error) {
	u := "/api/security/groups"
	v := new([]Group)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns the provided group.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetGroupDetails
func (s *GroupsService) Get(group string) (*Group, *Response, error) {
	u := fmt.Sprintf("/api/security/groups/%s", group)
	v := new(Group)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetIncludeUsers returns the provided group including its user membership (Artifactory >= 6.13)
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetGroupDetails
func (s *GroupsService) GetIncludeUsers(group string) (*Group, *Response, error) {
	u := fmt.Sprintf("/api/security/groups/%s?includeUsers=true", group)
	v := new(Group)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Create constructs a group with the provided details.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceGroup
func (s *GroupsService) Create(group *Group) (*string, *Response, error) {
	u := fmt.Sprintf("/api/security/groups/%s", *group.Name)
	v := new(string)

	resp, err := s.client.Call("PUT", u, group, v)
	return v, resp, err
}

// Update modifies a group with the provided details.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateGroup
func (s *GroupsService) Update(group *Group) (*string, *Response, error) {
	u := fmt.Sprintf("/api/security/groups/%s", *group.Name)
	v := new(string)

	resp, err := s.client.Call("POST", u, group, v)
	return v, resp, err
}

// Delete removes the provided group.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-DeleteGroup
func (s *GroupsService) Delete(group string) (*string, *Response, error) {
	u := fmt.Sprintf("/api/security/groups/%s", group)
	v := new(string)

	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}
