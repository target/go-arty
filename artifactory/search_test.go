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
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"github.com/target/go-arty/artifactory/fixtures/search"
)

func Test_Search(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(search.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Search Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		coords := &GAVCRequest{
			ArtifactID: String("folder"),
			Repos:      &[]string{"local-repo1"},
		}

		g.Describe("Search", func() {
			g.It("- should return valid string for GAVCResponse with String()", func() {
				actual := &GAVCResponse{
					Results: &[]File{
						File{
							URI: String("http://localhost:8081/artifactory/api/storage/local-repo1/folder/file.json"),
						},
						File{
							URI: String("http://localhost:8081/artifactory/api/storage/local-repo1/folder/foo.txt"),
						},
					},
				}

				data, _ := ioutil.ReadFile("fixtures/search/files.json")

				var expected GAVCResponse
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with GAVC()", func() {
				actual, resp, err := c.Search.GAVC(coords)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
