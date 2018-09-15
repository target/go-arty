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
	"strings"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"github.com/target/go-arty/artifactory/fixtures/artifacts"
)

func Test_Artifacts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(artifacts.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Artifacts Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Artifacts", func() {

			g.It("- should return valid string for ArtifactMessage with String()", func() {
				actual := &ArtifactMessage{
					Level:   String("INFO"),
					Message: String("copying local-repo1:folder/foo.txt to local-repo1:test/foo.txt completed successfully, 1 artifacts and 0 folders were copied"),
				}

				data, _ := ioutil.ReadFile("fixtures/artifacts/copy.json")

				var body Artifacts
				_ = json.Unmarshal(data, &body)

				expected := *body.Messages

				g.Assert(actual.String() == expected[0].String()).IsTrue()
			})

			g.It("- should return valid string for Artifacts with String()", func() {
				actual := &Artifacts{
					Messages: &[]ArtifactMessage{
						ArtifactMessage{
							Level:   String("INFO"),
							Message: String("copying local-repo1:folder/foo.txt to local-repo1:test/foo.txt completed successfully, 1 artifacts and 0 folders were copied"),
						},
					},
				}

				data, _ := ioutil.ReadFile("fixtures/artifacts/copy.json")

				var expected Artifacts
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with Download()", func() {
				actual, resp, err := c.Artifacts.Download("local-repo1", "foo.txt")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Upload() using 1 property", func() {
				actual, resp, err := c.Artifacts.Upload("local-repo1", "folder", "fixtures/artifacts/foo.txt", map[string][]string{"key": []string{"value"}})
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(resp.Request.URL.Path).Equal("/local-repo1/folder/fixtures/artifacts/foo.txt;key=value")
				g.Assert(err == nil).IsTrue()
				g.Assert(*actual).Equal("")
			})

			g.It("- bad file should return an error", func() {
				actual, resp, err := c.Artifacts.Upload("local-repo1", "folder", "this/is/a/bad/path.txt", map[string][]string{"key": []string{"value"}})
				g.Assert(actual == nil).IsTrue()
				g.Assert(resp == nil).IsTrue()
				g.Assert(err != nil).IsTrue()
				g.Assert(err.Error()).Equal("open this/is/a/bad/path.txt: no such file or directory")
			})

			g.It("- should return no error with Upload() using multiple properties", func() {
				actual, resp, err := c.Artifacts.Upload("local-repo1", "folder", "fixtures/artifacts/foo.txt", map[string][]string{"key": []string{"value", "value2", "value3"}})
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(resp.Request.URL.Path).Equal("/local-repo1/folder/fixtures/artifacts/foo.txt;key=value,value2,value3")
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Upload() using multiple property keys", func() {
				actual, resp, err := c.Artifacts.Upload("local-repo1", "folder", "fixtures/artifacts/foo.txt", map[string][]string{"key": []string{"value", "value2", "value3"}, "key2": []string{"anothervalue"}})
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(strings.Contains(resp.Request.URL.Path, "key=value,value2,value3")).IsTrue()
				g.Assert(strings.Contains(resp.Request.URL.Path, "key2=anothervalue")).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Copy()", func() {
				actual, resp, err := c.Artifacts.Copy("local-repo1", "folder/foo.txt", "local-repo1", "test/foo.txt")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Move()", func() {
				actual, resp, err := c.Artifacts.Move("local-repo1", "folder/foo.txt", "local-repo1", "test/foo.txt")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Delete()", func() {
				actual, resp, err := c.Artifacts.Delete("local-repo1", "foo.txt")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
