// Copyright (c) 2018 Target Brands, Inc.
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

package xray

// SystemService handles communication with the system related
// methods of the Xray API.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-SUMMARY
type SystemService service

// Ping represents the ping response status from Xray
type Ping struct {
	Status string `json:"status,omitempty"`
}

// Versions represents the version information about Xray.
type Versions struct {
	XrayVersion  string `json:"xray_version,omitempty"`
	XrayRevision string `json:"xray_revision,omitempty"`
}

// Ping returns a simple status response.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-PingRequest
func (s *SystemService) Ping() (*Ping, *Response, error) {
	u := "/api/v1/system/ping"
	v := new(Ping)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Version returns information about the current version.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-GetVersion
func (s *SystemService) Version() (*Versions, *Response, error) {
	u := "/api/v1/system/version"
	v := new(Versions)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}
