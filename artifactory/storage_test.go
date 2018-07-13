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
	"net/http/httptest"
	"testing"

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
