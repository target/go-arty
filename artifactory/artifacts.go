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
	"os"
	"strings"
)

// ArtifactsService handles communication with the artifact related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SECURITY
type ArtifactsService service

// ArtifactMessage represents an artifact message in Artifactory.
type ArtifactMessage struct {
	Level   *string `json:"level,omitempty"`
	Message *string `json:"message,omitempty"`
}

func (a ArtifactMessage) String() string {
	return Stringify(a)
}

// Artifacts represents artifacts in Artifactory.
type Artifacts struct {
	Messages *[]ArtifactMessage `json:"messages,omitempty"`
}

func (a Artifacts) String() string {
	return Stringify(a)
}

// Download retrieves the provided artifact.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveArtifact
func (s *ArtifactsService) Download(repo, path string) (*[]byte, *Response, error) {
	u := fmt.Sprintf("/%s/%s", repo, path)
	v := new([]byte)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Upload deploys the provided artifact to the provided repository.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-DeployArtifact
func (s *ArtifactsService) Upload(repo, path, file string, properties map[string][]string) (*string, *Response, error) {
	var propertyString string
	var index int
	for k, v := range properties {
		index++
		if len(v) == 1 {
			propertyString = propertyString + fmt.Sprintf("%s=%s", k, v[0])
		} else {
			propertyString = propertyString + fmt.Sprintf("%s=%s", k, strings.Join(v, ","))
		}

		if index != len(properties) {
			propertyString = propertyString + ";"
		}
	}

	u := fmt.Sprintf("/%s/%s;%s", repo, path, propertyString)
	v := new(string)

	data, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = data.Close() }()

	resp, err := s.client.Call("PUT", u, data, v)
	return v, resp, err
}

// Copy duplicates the provided artifact to the provided destination.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-CopyItem
func (s *ArtifactsService) Copy(sourceRepo, sourcePath, targetRepo, targetPath string) (*Artifacts, *Response, error) {
	u := fmt.Sprintf("/api/copy/%s/%s?to=/%s/%s", sourceRepo, sourcePath, targetRepo, targetPath)
	v := new(Artifacts)

	resp, err := s.client.Call("POST", u, nil, v)
	return v, resp, err
}

// Move migrates the provided artifact to the provided destination.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-MoveItem
func (s *ArtifactsService) Move(sourceRepo, sourcePath, targetRepo, targetPath string) (*Artifacts, *Response, error) {
	u := fmt.Sprintf("/api/move/%s/%s?to=/%s/%s", sourceRepo, sourcePath, targetRepo, targetPath)
	v := new(Artifacts)

	resp, err := s.client.Call("POST", u, nil, v)
	return v, resp, err
}

// Delete removes the provided artifact.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-DeleteItem
func (s *ArtifactsService) Delete(repo, path string) (*string, *Response, error) {
	u := fmt.Sprintf("/%s/%s", repo, path)
	v := new(string)

	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}
