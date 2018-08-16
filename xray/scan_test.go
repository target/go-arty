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
	"time"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"github.com/target/go-arty/xray/fixtures/scan"
)

func Test_Scan(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(scan.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Scan Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Scan", func() {
			scanArtifact := &ScanArtifactRequest{
				ComponentID: String(""),
			}

			scanBuild := &ScanBuildRequest{
				ArtifactoryID: String(""),
				BuildName:     String(""),
				BuildNumber:   String(""),
			}

			g.It("- should return valid string for ScanArtifactResponse with String()", func() {
				actual := &ScanArtifactResponse{
					Info: String("Scan of artifact is in progress"),
				}

				data, _ := ioutil.ReadFile("fixtures/scan/artifact.json")

				var expected ScanArtifactResponse
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for ScanBuildResponse with String()", func() {
				actual := &ScanBuildResponse{
					Summary: &ScanSummary{
						FailBuild:      Bool(false),
						Message:        String("string"),
						MoreDetailsURL: String("string"),
						TotalAlerts:    String("int"),
					},
					Alerts: &[]ScanAlert{
						ScanAlert{
							Created:     &Timestamp{time.Date(2011, time.November, 11, 11, 11, 11, 0, time.UTC)},
							TopSeverity: String("string"),
							WatchName:   String("string"),
							Issues: &[]ScanIssue{
								ScanIssue{
									Created:     &Timestamp{time.Date(2012, time.December, 12, 12, 12, 12, 0, time.UTC)},
									Cve:         String("string"),
									Description: String("string"),
									Provider:    String("string"),
									Severity:    String("string"),
									Summary:     String("string"),
									Type:        String("string"),
									ImpactedArtifacts: &[]ScanImpactedArtifact{
										ScanImpactedArtifact{
											Depth:       String("int"),
											DisplayName: String("string"),
											Name:        String("string"),
											ParentSha:   String("string"),
											Path:        String("string"),
											PkgType:     String("string"),
											Sha1:        String("string"),
											Sha256:      String("string"),
											InfectedFiles: &[]ScanInfectedFile{
												ScanInfectedFile{
													ComponentID: String("string"),
													Depth:       String("int"),
													DisplayName: String("string"),
													Name:        String("string"),
													ParentSha:   String("string"),
													Path:        String("string"),
													PkgType:     String("string"),
													Sha1:        String("string"),
													Sha256:      String("string"),
													Details: &[]ScanDetail{
														ScanDetail{
															BannedLicenses: &[]ScanBannedLicense{
																ScanBannedLicense{
																	AlertType:   String("string"),
																	Description: String("string"),
																	ID:          &struct{}{},
																	Severity:    String("string"),
																	Summary:     String("string"),
																},
															},
															Child: String("string"),
															Vulnerabilities: &[]ScanVulnerability{
																ScanVulnerability{
																	AlertType:   String("string"),
																	Description: String("string"),
																	ID:          &struct{}{},
																	Severity:    String("string"),
																	Summary:     String("string"),
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					Licenses: &[]ScanLicense{
						ScanLicense{
							Name:        String("string"),
							Components:  &[]string{"string"},
							FullName:    String("string"),
							MoreInfoURL: &[]string{"string"},
						},
					},
				}

				data, _ := ioutil.ReadFile("fixtures/scan/build.json")

				var expected ScanBuildResponse
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with Artifact()", func() {
				actual, resp, err := c.Scan.Artifact(scanArtifact)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Build()", func() {
				actual, resp, err := c.Scan.Build(scanBuild)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
