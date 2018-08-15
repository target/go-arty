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

// LicensesService handles communication with the license related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SECURITY
type LicensesService service

// License represents a license in Artifactory.
type License struct {
	Type         *string `json:"type,omitempty"`
	ValidThrough *string `json:"validThrough,omitempty"`
	LicensedTo   *string `json:"licensedTo,omitempty"`
}

func (l License) String() string {
	return Stringify(l)
}

// HALicense represents a HA license in Artifactory.
type HALicense struct {
	Type         *string `json:"type,omitempty"`
	ValidThrough *string `json:"validThrough,omitempty"`
	LicensedTo   *string `json:"licensedTo,omitempty"`
	LicenseHash  *string `json:"licenseHash,omitempty"`
	NodeID       *string `json:"nodeId,omitempty"`
	NodeURL      *string `json:"nodeUrl,omitempty"`
	Expired      *bool   `json:"expired,omitempty"`
}

func (h HALicense) String() string {
	return Stringify(h)
}

// HALicenses represents an array of HA licenses in Artifactory.
type HALicenses struct {
	Licenses *[]HALicense `json:"licenses,omitempty"`
}

func (h HALicenses) String() string {
	return Stringify(h)
}

// LicenseRequest represents the license request in Artifactory.
type LicenseRequest struct {
	LicenseKey *string `json:"licenseKey,omitempty"`
}

// LicenseResponse represents the response from installing a license in Artifactory.
type LicenseResponse struct {
	Status  *int    `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

func (l LicenseResponse) String() string {
	return Stringify(l)
}

// HALicenseResponse represents the response from installing a HA license(s) in Artifactory.
type HALicenseResponse struct {
	Status   *int               `json:"status,omitempty"`
	Messages *map[string]string `json:"messages,omitempty"`
}

func (h HALicenseResponse) String() string {
	return Stringify(h)
}

// LicenseRemoval is a list of license hashes for when removing licenses in Artifactory.
type LicenseRemoval struct {
	LicenseHashes *[]string `url:"licenseHash,omitempty"`
}

// Get returns a single license.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-LicenseInformation
func (s *LicensesService) Get() (*License, *Response, error) {
	u := "/api/system/licenses"
	v := new(License)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Install deploys the provided license to the instance.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-InstallLicense
func (s *LicensesService) Install(license *LicenseRequest) (*LicenseResponse, *Response, error) {
	u := "/api/system/licenses"
	v := new(LicenseResponse)

	resp, err := s.client.Call("POST", u, license, v)
	return v, resp, err
}

// GetHA returns a list of licenses for an HA cluster.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-HALicenseInformation
func (s *LicensesService) GetHA() (*HALicenses, *Response, error) {
	u := "/api/system/licenses"
	v := new(HALicenses)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// InstallHA deploys the provided license(s) to an HA cluster.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-InstallHAClusterLicenses
func (s *LicensesService) InstallHA(licenses *[]LicenseRequest) (*HALicenseResponse, *Response, error) {
	u := "/api/system/licenses"
	v := new(HALicenseResponse)

	resp, err := s.client.Call("POST", u, licenses, v)
	return v, resp, err
}

// DeleteHA removes the provided license key(s) from an HA cluster.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-DeleteHAClusterLicense
func (s *LicensesService) DeleteHA(hashes *LicenseRemoval) (*HALicenseResponse, *Response, error) {
	u := "/api/system/licenses"
	v := new(HALicenseResponse)

	u, err := addOptions(u, hashes)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}
