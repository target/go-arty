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

// ScanService handles communication with the scan related
// methods of the Xray API.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-SCAN
type ScanService service

// ScanArtifactRequest represents the ScanArtifact request we send to Xray.
type ScanArtifactRequest struct {
	Checksum    ScanChecksum `json:"checksum,omitempty"`
	ComponentID string       `json:"componentId,omitempty"`
	Summary     string       `json:"summary,omitempty"`
}

// ScanChecksum represents the checksums used in the scan artifact request in Xray.
type ScanChecksum struct {
	Md5    string `json:"md5,omitempty"`
	Sha1   string `json:"sha1,omitempty"`
	Sha256 string `json:"sha256,omitempty"`
}

// ScanArtifactResponse represents the response from scanning an artifact in Xray.
type ScanArtifactResponse struct {
	Info string `json:"info,omitempty"`
}

// ScanBuildRequest represents the scan build request we send to Xray.
type ScanBuildRequest struct {
	ArtifactoryID string `json:"artifactoryId,omitempty"`
	BuildName     string `json:"buildName,omitempty"`
	BuildNumber   string `json:"buildNumber,omitempty"`
}

// ScanBuildResponse represents the response from scanning a build in Xray.
type ScanBuildResponse struct {
	Summary  ScanSummary   `json:"summary,omitempty"`
	Alerts   []ScanAlert   `json:"alerts,omitempty"`
	Licenses []ScanLicense `json:"licenses,omitempty"`
}

// ScanSummary represents a summary of the scan in Xray.
type ScanSummary struct {
	FailBuild      string `json:"fail_build,omitempty"`
	Message        string `json:"message,omitempty"`
	MoreDetailsURL string `json:"more_details_url,omitempty"`
	TotalAlerts    string `json:"total_alerts,omitempty"`
}

// ScanAlert represents an alert of the scan in Xray.
type ScanAlert struct {
	Created     string      `json:"created,omitempty"`
	Issues      []ScanIssue `json:"issues,omitempty"`
	TopSeverity string      `json:"top_severity,omitempty"`
	WatchName   string      `json:"watch_name,omitempty"`
}

// ScanIssue represents a issue of the scan in Xray.
type ScanIssue struct {
	Created           string                 `json:"created,omitempty"`
	Cve               string                 `json:"cve,omitempty"`
	Description       string                 `json:"description,omitempty"`
	ImpactedArtifacts []ScanImpactedArtifact `json:"impacted_artifacts,omitempty"`
	Provider          string                 `json:"provider,omitempty"`
	Severity          string                 `json:"severity,omitempty"`
	Summary           string                 `json:"summary,omitempty"`
	Type              string                 `json:"type,omitempty"`
}

// ScanImpactedArtifact represents the impacted artifact of a scan in Xray.
type ScanImpactedArtifact struct {
	Depth         string             `json:"depth,omitempty"`
	DisplayName   string             `json:"display_name,omitempty"`
	InfectedFiles []ScanInfectedFile `json:"infected_files,omitempty"`
	Name          string             `json:"name,omitempty"`
	ParentSha     string             `json:"parent_sha,omitempty"`
	Path          string             `json:"path,omitempty"`
	PkgType       string             `json:"pkg_type,omitempty"`
	Sha1          string             `json:"sha1,omitempty"`
	Sha256        string             `json:"sha256,omitempty"`
}

// ScanInfectedFile represents the infected file of a scan in Xray.
type ScanInfectedFile struct {
	ComponentID string       `json:"component_id,omitempty"`
	Depth       string       `json:"depth,omitempty"`
	Details     []ScanDetail `json:"details,omitempty"`
	DisplayName string       `json:"display_name,omitempty"`
	Name        string       `json:"name,omitempty"`
	ParentSha   string       `json:"parent_sha,omitempty"`
	Path        string       `json:"path,omitempty"`
	PkgType     string       `json:"pkg_type,omitempty"`
	Sha1        string       `json:"sha1,omitempty"`
	Sha256      string       `json:"sha256,omitempty"`
}

// ScanDetail represents the detail of a scan in Xray.
type ScanDetail struct {
	BannedLicenses  []ScanBannedLicense `json:"banned_licenses,omitempty"`
	Child           string              `json:"child,omitempty"`
	Vulnerabilities []ScanVulnerability `json:"vulnerabilities,omitempty"`
}

// ScanBannedLicense represents the banned license of a scan in Xray.
type ScanBannedLicense struct {
	AlertType   string   `json:"alert_type,omitempty"`
	Description string   `json:"description,omitempty"`
	ID          struct{} `json:"id,omitempty"` // We don't have any build examples
	Severity    string   `json:"severity,omitempty"`
	Summary     string   `json:"summary,omitempty"`
}

// ScanVulnerability represents the vulnerability found from a scan in Xray.
type ScanVulnerability struct {
	AlertType   string   `json:"alert_type,omitempty"`
	Description string   `json:"description,omitempty"`
	ID          struct{} `json:"id,omitempty"` // We don't have any build examples
	Severity    string   `json:"severity,omitempty"`
	Summary     string   `json:"summary,omitempty"`
}

// ScanLicense represents the licence found from a scan in Xray.
type ScanLicense struct {
	Components  []string `json:"components,omitempty"`
	FullName    string   `json:"full_name,omitempty"`
	MoreInfoURL []string `json:"more_info_url,omitempty"`
	Name        string   `json:"name,omitempty"`
}

// Artifact invokes scanning of an artifact.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-ScanArtifact
func (s *ScanService) Artifact(scan *ScanArtifactRequest) (*ScanArtifactResponse, *Response, error) {
	u := "/api/v1/scanArtifact"
	v := new(ScanArtifactResponse)

	resp, err := s.client.Call("POST", u, scan, v)
	return v, resp, err
}

// Build invokes scanning of a build that was uploaded to Artifactory as requested by a CI server.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-ScanBuild
func (s *ScanService) Build(scan *ScanBuildRequest) (*ScanBuildResponse, *Response, error) {
	u := "/api/v1/scanBuild"
	v := new(ScanBuildResponse)

	resp, err := s.client.Call("POST", u, scan, v)
	return v, resp, err
}
