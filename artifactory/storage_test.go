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
	"time"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"github.com/target/go-arty/artifactory/fixtures/storage"
)

func Test_Storage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(storage.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Storage Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Storage", func() {

			g.It("- should return valid string for Folder with String()", func() {
				actual := &Folder{
					URI:          String("http://localhost:8081/artifactory/api/storage/local-repo1/folder"),
					Repo:         String("local-repo1"),
					Path:         String("/folder"),
					Created:      &Timestamp{time.Date(2010, time.October, 10, 10, 10, 10, 0, time.UTC)},
					CreatedBy:    String("admin"),
					LastModified: &Timestamp{time.Date(2011, time.November, 11, 11, 11, 11, 0, time.UTC)},
					ModifiedBy:   String("admin"),
					LastUpdated:  &Timestamp{time.Date(2012, time.December, 12, 12, 12, 12, 0, time.UTC)},
					Children: &[]Child{
						Child{URI: String("/file.json"), Folder: String("true")},
						Child{URI: String("/foo.txt"), Folder: String("false")},
					},
				}

				data, _ := ioutil.ReadFile("fixtures/storage/folder.json")

				var expected Folder
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for File with String()", func() {
				actual := &File{
					URI:          String("http://localhost:8081/artifactory/api/storage/local-repo1/folder/file.json"),
					DownloadURI:  String("http://localhost:8081/artifactory/local-repo1/folder/file.json"),
					Repo:         String("local-repo1"),
					Path:         String("/folder/file.json"),
					Created:      &Timestamp{time.Date(2010, time.October, 10, 10, 10, 10, 0, time.UTC)},
					CreatedBy:    String("admin"),
					LastModified: &Timestamp{time.Date(2011, time.November, 11, 11, 11, 11, 0, time.UTC)},
					ModifiedBy:   String("admin"),
					LastUpdated:  &Timestamp{time.Date(2012, time.December, 12, 12, 12, 12, 0, time.UTC)},
					Size:         String("1024"),
					MimeType:     String("application/json"),
					Checksums: &Checksums{
						MD5:    String("B45CFFE084DD3D20D928BEE85E7B0F21"),
						SHA1:   String("ECB252044B5EA0F679EE78EC1A12904739E2904D"),
						SHA256: String("473287F8298DBA7163A897908958F7C0EAE733E25D2E027992EA2EDC9BED2FA8"),
					},
					OriginalChecksums: &Checksums{
						MD5:    String("6DDB57974C449A3BE93F3124211373C4"),
						SHA1:   String("B680C4A75B05C5AAB4C365D68D9FACF42482BC64"),
						SHA256: String("5FE075210D189874420BC9EDFBB6216FBCEAABE3B1792D2C53C391A96009CA55"),
					},
				}

				data, _ := ioutil.ReadFile("fixtures/storage/file.json")

				var expected File
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for StorageSummary with String()", func() {
				actual := &StorageSummary{
					BinariesSummary: &BinariesSummary{
						BinariesCount:  String("125,726"),
						BinariesSize:   String("3.48 GB"),
						ArtifactsSize:  String("59.77 GB"),
						Optimization:   String("5.82%"),
						ItemsCount:     String("2,176,580"),
						ArtifactsCount: String("2,084,408"),
					},
					FileStoreSummary: &FileStoreSummary{
						StorageType:      String("filesystem"),
						StorageDirectory: String("/home/.../artifactory/devenv/.artifactory/data/filestore"),
						TotalSpace:       String("204.28 GB"),
						UsedSpace:        String("32.22 GB (15.77%)"),
						FreeSpace:        String("172.06 GB (84.23%)"),
					},
					RepositoriesSummaryList: &[]RepositoriesSummary{
						RepositoriesSummary{
							RepoKey:      String("plugins-release"),
							RepoType:     String("VIRTUAL"),
							FoldersCount: Int(0),
							FilesCount:   Int(0),
							UsedSpace:    String("0 bytes"),
							ItemsCount:   Int(0),
							PackageType:  String("Maven"),
							Percentage:   String("0%"),
						},
						RepositoriesSummary{
							RepoKey:      String("repo"),
							RepoType:     String("VIRTUAL"),
							FoldersCount: Int(0),
							FilesCount:   Int(0),
							UsedSpace:    String("0 bytes"),
							ItemsCount:   Int(0),
							PackageType:  String("Generic"),
							Percentage:   String("0%"),
						},
						RepositoriesSummary{
							RepoKey:      String("TOTAL"),
							RepoType:     String("NA"),
							FoldersCount: Int(92172),
							FilesCount:   Int(2084408),
							UsedSpace:    String("59.77 GB"),
							ItemsCount:   Int(2176580),
						},
					},
				}

				data, _ := ioutil.ReadFile("fixtures/storage/storage_summary.json")

				var expected StorageSummary
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			// TODO: I hate the randomized order of maps
			// g.It("- should return valid string for EffectiveItemPermissions with String()", func() {
			// 	actual := &EffectiveItemPermissions{
			// 		URI: String("http://localhost:8081/artifactory/api/storage/local-repo1/file"),
			// 		Principals: &Principals{
			// 			Users:  &map[string][]string{"user1": []string{"r"}, "user2": []string{"r", "d", "w", "m", "n"}},
			// 			Groups: &map[string][]string{"readers": []string{"r"}},
			// 		},
			// 	}
			// 	data, _ := ioutil.ReadFile("fixtures/storage/effective_permission.json")

			// 	var expected EffectiveItemPermissions
			// 	_ = json.Unmarshal(data, &expected)

			// 	g.Assert(actual.String() == expected.String()).IsTrue()
			// })

			g.It("- should return valid string for ItemLastModified with String()", func() {
				actual := &ItemLastModified{
					URI:          String("http://localhost:8081/artifactory/api/storage/local-repo1/folder/file.json"),
					LastModified: &Timestamp{time.Date(2011, time.November, 11, 11, 11, 11, 0, time.UTC)},
				}

				data, _ := ioutil.ReadFile("fixtures/storage/last_modified.json")

				var expected ItemLastModified
				_ = json.Unmarshal(data, &expected)
				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for FileStatistics with String()", func() {
				actual := &FileStatistics{
					URI:              String("http://localhost:8081/artifactory/api/storage/local-repo1/folder/file.json"),
					LastDownloaded:   &Timestamp{time.Date(2012, time.December, 12, 12, 12, 12, 0, time.UTC)},
					DownloadCount:    Int(3),
					LastDownloadedBy: String("admin"),
				}

				data, _ := ioutil.ReadFile("fixtures/storage/statistics.json")

				var expected FileStatistics
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for ItemProperties with String()", func() {
				actual := &ItemProperties{
					URI:        String("http://localhost:8081/artifactory/api/storage/local-repo1/folder/file.json"),
					Properties: &map[string][]string{"p1": []string{"v1", "v2", "v3"}},
				}

				data, _ := ioutil.ReadFile("fixtures/storage/properties.json")

				var expected ItemProperties
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for FileList with String()", func() {
				actual := &FileList{
					URI:     String("http://localhost:8081/artifactory/api/storage/local-repo1/folder"),
					Created: &Timestamp{time.Date(2010, time.October, 10, 10, 10, 10, 0, time.UTC)},
					Files: &[]FileListItem{
						FileListItem{
							URI:          String("/file.json"),
							Size:         Int(253207),
							LastModified: &Timestamp{time.Date(2011, time.November, 11, 11, 11, 11, 0, time.UTC)},
							Folder:       Bool(false),
							SHA1:         String("ECB252044B5EA0F679EE78EC1A12904739E2904D"),
						},
						FileListItem{
							URI:          String("/foo.txt"),
							Size:         Int(253100),
							LastModified: &Timestamp{time.Date(2012, time.December, 12, 12, 12, 12, 0, time.UTC)},
							Folder:       Bool(false),
							SHA1:         String("B680C4A75B05C5AAB4C365D68D9FACF42482BC64"),
						},
					},
				}

				data, _ := ioutil.ReadFile("fixtures/storage/file_list.json")

				var expected FileList
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with GetFolder()", func() {
				actual, resp, err := c.Storage.GetFolder("local-repo1", "folder")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetFile()", func() {
				actual, resp, err := c.Storage.GetFile("local-repo1", "file")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetItemLastModified()", func() {
				actual, resp, err := c.Storage.GetItemLastModified("local-repo1", "file")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetFileStatistics()", func() {
				actual, resp, err := c.Storage.GetFileStatistics("local-repo1", "file")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetItemProperties()", func() {
				actual, resp, err := c.Storage.GetItemProperties("local-repo1", "file")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetEffectiveItemPermissions()", func() {
				actual, resp, err := c.Storage.GetEffectiveItemPermissions("local-repo1", "file")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with SetItemProperties()", func() {
				properties := make(map[string][]string)
				properties["p2"] = []string{"v1"}
				resp, err := c.Storage.SetItemProperties("local-repo1", "file", properties)
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with DeleteItemProperties()", func() {
				resp, err := c.Storage.DeleteItemProperties("local-repo1", "file", []string{"p2"})
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetFileList()", func() {
				actual, resp, err := c.Storage.GetFileList("local-repo1", "folder")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetStorageSummary()", func() {
				actual, resp, err := c.Storage.GetStorageSummary()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
