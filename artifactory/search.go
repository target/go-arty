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

// SearchService handles communication with the search related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SEARCHES
type SearchService service

// GAVCRequest represents the GAVC request for searches in Artifactory.
type GAVCRequest struct {
	GroupID    *string   `url:"g,omitempty"`
	ArtifactID *string   `url:"a,omitempty"`
	Version    *string   `url:"v,omitempty"`
	Classifier *string   `url:"c,omitempty"`
	Repos      *[]string `url:"repos,omitempty"`
}

// GAVCResponse represents the GAVC response for searches in Artifactory.
type GAVCResponse struct {
	Results *[]File `json:"results,omitempty"`
}

// GAVC returns the list of artifacts from the Maven search.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GAVCSearch
func (s *SearchService) GAVC(coords *GAVCRequest) (*GAVCResponse, *Response, error) {
	u := "/api/search/gavc"
	v := new(GAVCResponse)

	u, err := addOptions(u, coords)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}
