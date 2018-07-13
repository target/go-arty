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

// DockerService handles communication with the Docker related
// methods of the Artifactory API.
type DockerService service

// Registry represents the list of Docker repositories in a registry in Artifactory.
type Registry struct {
	Repositories []string `json:"repositories,omitempty"`
}

// Tags represents the list of tags for a Docker repository in Artifactory.
type Tags struct {
	Name string   `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

// ImagePromotion represents the Docker image promotion request in Artifactory.
type ImagePromotion struct {
	TargetRepo             string `json:"targetRepo,omitempty"`             // The target repository for the move or copy
	DockerRepository       string `json:"dockerRepository,omitempty"`       // The docker repository name to promote
	TargetDockerRepository string `json:"targetDockerRepository,omitempty"` // An optional Docker repository name, if null, will use the same name as 'dockerRepository'
	Tag                    string `json:"tag,omitempty"`                    // An optional tag name to promote, if null - the entire docker repository will be promoted. Available from v4.10.
	TargetTag              string `json:"targetTag,omitempty"`              // An optional target tag to assign the image after promotion, if null - will use the same tag
	Copy                   bool   `json:"copy,omitempty"`                   // An optional value to set whether to copy instead of move. Default: false
}

// GetRepositories returns a list of all Docker repositories for the provided registry.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ListDockerRepositories
func (s *DockerService) GetRepositories(registry string) (*Registry, *Response, error) {
	u := fmt.Sprintf("/api/docker/%s/v2/_catalog", registry)
	v := new(Registry)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetTags returns a list of all tags for the provided Docker repository.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ListDockerTags
func (s *DockerService) GetTags(registry, repository string) (*Tags, *Response, error) {
	u := fmt.Sprintf("/api/docker/%s/v2/%s/tags/list", registry, repository)
	v := new(Tags)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// PromoteImage promotes the provided Docker image(s) from the provided source repository to the provided destination repository.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-PromoteDockerImage
func (s *DockerService) PromoteImage(registry string, promotion ImagePromotion) (*string, *Response, error) {
	u := fmt.Sprintf("/api/docker/%s/v2/promote", registry)
	v := new(string)

	resp, err := s.client.Call("POST", u, promotion, v)
	return v, resp, err
}
