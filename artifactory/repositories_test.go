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
	"github.com/target/go-arty/artifactory/fixtures/repositories"
)

func Test_Repositories(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(repositories.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Repositories Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Repositories", func() {
			repo := GenericRepository{
				Key:         String("test"),
				RClass:      String("local"),
				PackageType: String("generic"),
				Description: String("test repository"),
			}

			g.It("- should return valid string for Repository with String()", func() {
				actual := &Repository{
					Key:         String("libs-releases-local"),
					Type:        String("LOCAL"),
					Description: String("Local repository for in-house libraries"),
					URL:         String("http://localhost:8081/artifactory/libs-releases-local"),
					PackageType: String("NuGet"),
				}

				data, _ := ioutil.ReadFile("fixtures/repositories/repositories.json")

				var expected []Repository
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected[0].String()).IsTrue()
			})

			g.It("- should return valid string for GenericRepository with String()", func() {
				actual := &GenericRepository{
					Key:                          String("generic-repo1"),
					RClass:                       String(""),
					PackageType:                  String("maven"),
					Description:                  String("The generic repository public description"),
					Notes:                        String("Some internal notes"),
					IncludesPattern:              String("**/*"),
					ExcludesPattern:              String(""),
					LayoutRef:                    String("maven-2-default"),
					HandleReleases:               Bool(true),
					HandleSnapshots:              Bool(false),
					MaxUniqueSnapshots:           Int(0),
					SuppressPomConsistencyChecks: Bool(true),
					BlackedOut:                   Bool(false),
					PropertySets:                 &[]string{"artifactory"},
				}

				data, _ := ioutil.ReadFile("fixtures/repositories/generic_repository.json")

				var expected GenericRepository
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for LocalRepository with String()", func() {
				actual := &LocalRepository{
					GenericRepository: &GenericRepository{
						Key:                          String("local-repo1"),
						RClass:                       String("local"),
						PackageType:                  String("generic"),
						Description:                  String("The local repository public description"),
						Notes:                        String("Some internal notes"),
						IncludesPattern:              String("**/*"),
						ExcludesPattern:              String(""),
						LayoutRef:                    String("maven-2-default"),
						HandleReleases:               Bool(true),
						HandleSnapshots:              Bool(true),
						MaxUniqueSnapshots:           Int(0),
						SuppressPomConsistencyChecks: Bool(false),
						BlackedOut:                   Bool(false),
						PropertySets:                 &[]string{"ps1", "ps2"},
					},
					DebianTrivialLayout:             Bool(false),
					ChecksumPolicyType:              String("client-checksums"),
					MaxUniqueTags:                   Int(0),
					SnapshotVersionBehavior:         String("unique"),
					ArchiveBrowsingEnabled:          Bool(false),
					CalculateYumMetadata:            Bool(false),
					YumRootDepth:                    Int(0),
					DockerAPIVersion:                String("V2"),
					EnableFileListsIndexing:         Bool(false),
					OptionalIndexCompressionFormats: &[]string{"bz2", "lzma", "xz"},
				}

				data, _ := ioutil.ReadFile("fixtures/repositories/local_repository.json")

				var expected LocalRepository
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for RemoteRepository with String()", func() {
				actual := &RemoteRepository{
					GenericRepository: &GenericRepository{
						Key:                          String("remote-repo1"),
						RClass:                       String("remote"),
						PackageType:                  String("generic"),
						Description:                  String("The remote repository public description"),
						Notes:                        String("Some internal notes"),
						IncludesPattern:              String("**/*"),
						ExcludesPattern:              String(""),
						LayoutRef:                    String("maven-2-default"),
						HandleReleases:               Bool(true),
						HandleSnapshots:              Bool(true),
						MaxUniqueSnapshots:           Int(0),
						SuppressPomConsistencyChecks: Bool(false),
						BlackedOut:                   Bool(false),
						PropertySets:                 &[]string{"ps1", "ps2"},
					},
					URL:                               String("http://host:port/some-repo"),
					Username:                          String("remote-repo-user"),
					Password:                          String("pass"),
					Proxy:                             String("proxy1"),
					RemoteRepoChecksumPolicyType:      String("generate-if-absent"),
					HardFail:                          Bool(false),
					Offline:                           Bool(false),
					StoreArtifactsLocally:             Bool(true),
					SocketTimeoutMillis:               Int(15000),
					LocalAddress:                      String("212.150.139.167"),
					RetrievalCachePeriodSecs:          Int(43200),
					FailedRetrievalCachePeriodSecs:    Int(30),
					MissedRetrievalCachePeriodSecs:    Int(7200),
					UnusedArtifactsCleanupEnabled:     Bool(false),
					UnusedArtifactsCleanupPeriodHours: Int(0),
					AssumedOfflinePeriodSecs:          Int(300),
					FetchJarsEagerly:                  Bool(false),
					FetchSourcesEagerly:               Bool(false),
					ShareConfiguration:                Bool(false),
					SynchronizeProperties:             Bool(false),
					BlockMismatchingMimeTypes:         Bool(true),
					AllowAnyHostAuth:                  Bool(false),
					EnableCookieManagement:            Bool(false),
					BowerRegistryURL:                  String("https://registry.bower.io"),
					ComposerRegistryURL:               String("https://packagist.org"),
					PyPIRegistryURL:                   String("https://pypi.org"),
					VcsType:                           String("GIT"),
					VcsGitProvider:                    String("CUSTOM"),
					VcsGitDownloadUrl:                 String(""),
					BypassHeadRequest:                 Bool(false),
					ClientTLSCertificate:              String(""),
				}

				data, _ := ioutil.ReadFile("fixtures/repositories/remote_repository.json")

				var expected RemoteRepository
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for VirtualRepository with String()", func() {
				actual := &VirtualRepository{
					GenericRepository: &GenericRepository{
						Key:             String("virtual-repo1"),
						RClass:          String("virtual"),
						PackageType:     String("generic"),
						Description:     String("The virtual repository public description"),
						Notes:           String("Some internal notes"),
						IncludesPattern: String("**/*"),
						ExcludesPattern: String(""),
					},
					Repositories:        &[]string{"local-repo1", "local-repo2", "remote-repo1", "virtual-repo2"},
					DebianTrivialLayout: Bool(false),
					ArtifactoryRequestsCanRetrieveRemoteArtifacts: Bool(false),
					KeyPair:                              String("keypair1"),
					PomRepositoryReferencesCleanupPolicy: String("discard_active_reference"),
					DefaultDeploymentRepo:                String("local-repo1"),
				}

				data, _ := ioutil.ReadFile("fixtures/repositories/virtual_repository.json")

				var expected VirtualRepository
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with GetAll()", func() {
				actual, resp, err := c.Repositories.GetAll()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Get()", func() {
				actual, resp, err := c.Repositories.Get("local-repo1")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Get()", func() {
				actual, resp, err := c.Repositories.Get("remote-repo1")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Get()", func() {
				actual, resp, err := c.Repositories.Get("virtual-repo1")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Get()", func() {
				actual, resp, err := c.Repositories.Get("generic-repo1")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Create()", func() {
				actual, resp, err := c.Repositories.Create("test", repo)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Update()", func() {
				actual, resp, err := c.Repositories.Update("test", repo)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Delete()", func() {
				actual, resp, err := c.Repositories.Delete("test")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
