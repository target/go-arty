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
	"github.com/target/go-arty/artifactory/fixtures/builds"
)

func Test_Builds(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(builds.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Build Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Builds", func() {
			build := &Build{}

			g.BeforeEach(func() {
				build = &Build{
					BuildInfo: &BuildInfo{
						Properties: &map[string]string{
							"buildInfo.env.COLORTERM":           "truecolor",
							"buildInfo.env.LC_TERMINAL_VERSION": "3.3.0",
							"buildInfo.env.TERM":                "xterm-256color",
						},
						Version: String("1.1.1"),
						Name:    String("foo/bar"),
						Number:  String("0.1.0"),
						BuildAgent: &Agent{
							Name:    String("GENERIC"),
							Version: String("1.23.1"),
						},
						Agent: &Agent{
							Name:    String("jfrog-cli-go"),
							Version: String("1.23.1"),
						},
						Started:              String("2019-08-19T16:10:41.614-0500"),
						DurationMillis:       Int(0),
						ArtifactoryPrincipal: String("foo"),
						Modules: &[]Modules{
							Modules{
								Properties: &map[string]string{},
								ID:         String("foo:bar:0.1.0"),
								Artifacts: &[]BuildArtifacts{
									BuildArtifacts{
										Sha1:   String("088cb80eb3d651149c4ced1181ac9170d49d0069"),
										Sha256: String("8b77d882fecbabea512c10176b1fd0e117cc33a096640e6ee553aaa9eb70daa6"),
										Md5:    String("f2d9a2e6a3b4f988b66151c7a9747403"),
										Name:   String("foo-0.1.0.tgz"),
									},
								},
							},
						},
					},
					URI: String("https://artifactory.com/artifactory/api/build/foo/foo/0.1.0"),
				}
			})

			g.It("- should return valid string for Build with String()", func() {
				data, _ := ioutil.ReadFile("fixtures/builds/build.json")

				var expected Build
				_ = json.Unmarshal(data, &expected)

				g.Assert(build.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with GetInfo()", func() {
				actual, resp, err := c.Build.GetInfo("foo", "0.1.0")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})
	})
}
