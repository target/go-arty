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

// SystemService handles communication with the system related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SYSTEM&CONFIGURATION
type SystemService service

// Versions represents the version information about Artifactory.
type Versions struct {
	Version  *string   `json:"version,omitempty"`
	Revision *string   `json:"revision,omitempty"`
	Addons   *[]string `json:"addons,omitempty"`
}

func (v Versions) String() string {
	return Stringify(v)
}

// Ping returns a simple status response.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SystemHealthPing
func (s *SystemService) Ping() (*string, *Response, error) {
	u := "/api/system/ping"
	v := new(string)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns the general system information.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SystemInfo
func (s *SystemService) Get() (*string, *Response, error) {
	u := "/api/system"
	v := new(string)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetVersionAndAddOns returns information about the current version, revision, and installed add-ons.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-VersionandAdd-onsinformation
func (s *SystemService) GetVersionAndAddOns() (*Versions, *Response, error) {
	u := "/api/system/version"
	v := new(Versions)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}
