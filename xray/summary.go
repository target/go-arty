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

import "fmt"

// SummaryService handles communication with the summary related
// methods of the Xray API.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-SUMMARY
type SummaryService service

// SummaryArtifactRequest represents the SummaryRequest request we send to Xray.
type SummaryArtifactRequest struct {
	Checksums *[]string `json:"checksums,omitempty"`
	Paths     *[]string `json:"paths,omitempty"`
}

// SummaryResponse represents the response from a summary in Xray.
type SummaryResponse struct {
	Artifacts *[]SummaryArtifact `json:"artifacts,omitempty"`
	Errors    *[]SummaryError    `json:"errors,omitempty"`
}

// SummaryArtifact represents a artifact within the summary in Xray.
type SummaryArtifact struct {
	General  *SummaryGeneral   `json:"general,omitempty"`
	Issues   *[]SummaryIssue   `json:"issues,omitempty"`
	Licenses *[]SummaryLicense `json:"licenses,omitempty"`
}

// SummaryIssue represents a issue within the summary in Xray.
type SummaryIssue struct {
	Created     *string   `json:"created,omitempty"`
	Description *string   `json:"description,omitempty"`
	ImpactPath  *[]string `json:"impact_path,omitempty"`
	IssueType   *string   `json:"issue_type,omitempty"`
	Provider    *string   `json:"provider,omitempty"`
	Severity    *string   `json:"severity,omitempty"`
	Summary     *string   `json:"summary,omitempty"`
}

// SummaryGeneral represents the general information of a summary in Xray.
type SummaryGeneral struct {
	ComponentID *string `json:"component_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Path        *string `json:"path,omitempty"`
	PkgType     *string `json:"pkg_type,omitempty"`
	Sha256      *string `json:"sha256,omitempty"`
}

// SummaryLicense represents a licence found from a summary in Xray.
type SummaryLicense struct {
	Components  *[]string `json:"components,omitempty"`
	FullName    *string   `json:"full_name,omitempty"`
	MoreInfoURL *[]string `json:"more_info_url,omitempty"`
	Name        *string   `json:"name,omitempty"`
}

// SummaryError represents an error from a summery in Xray.
type SummaryError struct {
	Error      *string `json:"error,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
}

// Artifact provides details about any artifact specified by path identifiers or checksum.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-ArtifactSummary
func (s *SummaryService) Artifact(summary *SummaryArtifactRequest) (*SummaryResponse, *Response, error) {
	u := "/api/v1/summary/artifact"
	v := new(SummaryResponse)

	resp, err := s.client.Call("POST", u, summary, v)
	return v, resp, err
}

// Build provides details about any build specified by path identifiers or checksum.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-BuildSummary
func (s *SummaryService) Build(buildName string, buildNumber int) (*SummaryResponse, *Response, error) {
	u := fmt.Sprintf("/api/v1/summary/build?build_name=%s&build_number=%d", buildName, buildNumber)
	v := new(SummaryResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}
