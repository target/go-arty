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

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"github.com/target/go-arty/xray/fixtures/summary"
)

func Test_Summary(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(summary.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Summary Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Summary", func() {
			artifactRequest := &SummaryArtifactRequest{
				Paths: &[]string{"art2/ext-release-local/artifactory-pro.zip"},
			}

			g.It("- should return valid string for SummaryResponse with String()", func() {
				actual := &SummaryResponse{
					Artifacts: &[]SummaryArtifact{
						SummaryArtifact{
							General: &SummaryGeneral{
								ComponentID: String("gav://org.artifactory.pro:artifactory-pro-war:4.14.0"),
								Name:        String("artifactory-pro.zip"),
								Path:        String("art2/ext-release-local/"),
								PkgType:     String("Generic"),
								Sha256:      String("d160c68ed8879ae42756e159daec1dd7ecfd53b6192321656b72715e20d46dd2"),
							},
							Issues: &[]SummaryIssue{
								SummaryIssue{
									Summary:     String("FileSystemBytecodeCache in Jinja2 2.7.2 does not properly create temporary directories"),
									Description: String("this is the description of the issue"),
									IssueType:   String("security"),
									Severity:    String("Major"),
									Provider:    String("JFrog"),
									Created:     String("2016-10-26T11:15:51.17Z"),
									ImpactPath:  &[]string{"xray-artifactory/maven-1000/com/atlassian/aui/auiplugin/0.0.5-9-0-snapshot-035-do-not-use/Jinja2-2.7.2"},
								},
							},
							Licenses: &[]SummaryLicense{
								SummaryLicense{
									Name:        String("MIT"),
									FullName:    String("The MIT License"),
									MoreInfoURL: &[]string{"https://opensource.org/licenses/MIT"},
									Components: &[]string{
										"some-component-1",
										"some-component-2",
										"some-component-3",
									},
								},
								SummaryLicense{
									Name:        String("AGPL-3.0"),
									FullName:    String("GNU AFFERO GENERAL PUBLIC LICENSE, Version 3"),
									MoreInfoURL: &[]string{"https://opensource.org/licenses/AGPL-3.0"},
									Components: &[]string{
										"some-component-4",
										"some-component-5",
									},
								},
								SummaryLicense{
									Name: String("unknown"),
									Components: &[]string{
										"some-component-6",
										"some-component-7",
									},
								},
							},
						},
					},
					Errors: &[]SummaryError{
						SummaryError{
							Error:      String("Artifact doesn't exist or not indexed in Xray"),
							Identifier: String("4e39f19212597312ee02db873847bcb12c17cc639898bd2fd9b6a4aff16690e5"),
						},
					},
				}

				data, _ := ioutil.ReadFile("fixtures/summary/summary.json")

				var expected SummaryResponse
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with SummaryArtifact()", func() {
				actual, resp, err := c.Summary.Artifact(artifactRequest)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with SummaryBuild()", func() {
				actual, resp, err := c.Summary.Build("test", 1)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
