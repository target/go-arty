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
	"github.com/target/go-arty/artifactory/fixtures/permissions"
)

func Test_PermissionsV2(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(permissions.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("PermissionsV2 Service", func() {
		target := &PermissionTargetV2{}
		users := make(map[string][]string)
		groups := make(map[string][]string)

		g.BeforeEach(func() {
			users["bob"] = []string{"read", "write", "manage"}
			users["alice"] = []string{"write", "annotate", "read"}

			groups["dev-leads"] = []string{"manage", "read", "annotate"}
			groups["readers"] = []string{"read"}

			target = &PermissionTargetV2{
				Name: String("java-developers"),
				Repo: &PermissionDetails{
					IncludePatterns: &[]string{"**"},
					ExcludePatterns: &[]string{""},
					Repositories:    &[]string{"local-rep1", "local-rep2", "remote-rep1", "virtual-rep2"},
					Actions: &Actions{
						Users:  &users,
						Groups: &groups,
					},
				},
				Build: &PermissionDetails{
					IncludePatterns: &[]string{""},
					ExcludePatterns: &[]string{""},
					Repositories:    &[]string{"artifactory-build-info"},
					Actions: &Actions{
						Users:  &users,
						Groups: &groups,
					},
				},
				ReleaseBundle: &PermissionDetails{
					IncludePatterns: &[]string{"**"},
					ExcludePatterns: &[]string{""},
					Repositories:    &[]string{"release-bundles"},
					Actions: &Actions{
						Users:  &users,
						Groups: &groups,
					},
				},
			}
		})

		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("PermissionsV2", func() {

			g.It("- should return true for existing permission target", func() {
				actual, err := c.PermissionsV2.Exists("valid")
				g.Assert(actual).Equal(true)
				g.Assert(err).Equal(nil)
			})

			g.It("- should return false for non existent permission target", func() {
				actual, err := c.PermissionsV2.Exists("foo")
				g.Assert(actual).Equal(false)
				g.Assert(err).Equal(nil)
			})

			g.It("- should return no error with update", func() {
				actual, resp, err := c.PermissionsV2.Update(target)
				g.Assert(resp.StatusCode).Equal(200)
				g.Assert(actual != nil).IsTrue()
				g.Assert(err).Equal(nil)
			})

			g.It("- should return error with update", func() {
				target.Name = String("invalid")
				actual, resp, err := c.PermissionsV2.Update(target)
				g.Assert(resp.StatusCode).Equal(400)
				g.Assert(actual != nil).IsTrue()
				g.Assert(err != nil).IsTrue()
			})

			g.It("- should return valid string for PermissionTargetV2 with String()", func() {
				data, _ := ioutil.ReadFile("fixtures/permissions/permissionv2.json")

				var expected PermissionTargetV2
				_ = json.Unmarshal(data, &expected)

				g.Assert(target.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with get", func() {
				actual, resp, err := c.PermissionsV2.Get("java-developers")
				g.Assert(resp.StatusCode).Equal(200)
				g.Assert(actual != nil).Equal(true)
				g.Assert(err).Equal(nil)
			})

			g.It("- should return error with get due to non existent permission", func() {
				actual, resp, err := c.PermissionsV2.Get("foobar")
				g.Assert(resp.StatusCode).Equal(404)
				g.Assert(actual != nil).IsTrue()
				g.Assert(err != nil).IsTrue()
			})

		})
	})
}
