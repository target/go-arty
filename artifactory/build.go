package artifactory

import "fmt"

// BuildService handles communication with the builds related
// methods of the Artifactory API.
type BuildService service

// Agent represents the agent in the build
type Agent struct {
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// Modules contains information about modules within a build
type Modules struct {
	Properties *map[string]string `json:"properties,omitempty"`
	ID         *string            `json:"id,omitempty"`
	Artifacts  *[]BuildArtifacts  `json:"artifacts,omitempty"`
}

// BuildArtifacts contains information about build artifacts
type BuildArtifacts struct {
	Sha1   *string `json:"sha1,omitempty"`
	Sha256 *string `json:"sha256,omitempty"`
	Md5    *string `json:"md5,omitempty"`
	Name   *string `json:"name,omitempty"`
}

// BuildInfo represent the build payload in Artifactory
type BuildInfo struct {
	Properties           *map[string]string `json:"properties,omitempty"`
	Version              *string            `json:"version,omitempty"`
	Name                 *string            `json:"name,omitempty"`
	Number               *string            `json:"number,omitempty"`
	BuildAgent           *Agent             `json:"buildAgent,omitempty"`
	Agent                *Agent             `json:"agent,omitempty"`
	Started              *string            `json:"started,omitempty"`
	DurationMillis       *int               `json:"durationMillis,omitempty"`
	ArtifactoryPrincipal *string            `json:"artifactoryPrincipal,omitempty"`
	Modules              *[]Modules         `json:"modules,omitempty"`
}

// Build represents a build in Artifactory.
type Build struct {
	BuildInfo *BuildInfo `json:"buildInfo,omitempty"`
	URI       *string    `json:"uri,omitempty"`
}

func (b Build) String() string {
	return Stringify(b)
}

// GetInfo retrieves the provided build.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-BuildInfo
func (s *BuildService) GetInfo(name, number string) (*Build, *Response, error) {
	u := fmt.Sprintf("/api/build/%s/%s", name, number)
	v := new(Build)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}
