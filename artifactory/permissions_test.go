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

	"github.com/target/go-arty/artifactory/fixtures/permissions"
	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
)

func Test_Permissions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(permissions.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Permissions Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Permissions", func() {
			users := make(map[string][]string)
			users["test"] = []string{"r", "w", "n"}
			target := PermissionTarget{
				Name:            "test",
				IncludesPattern: "**",
				ExcludesPattern: "",
				Repositories:    []string{"test"},
				Principals: Principals{
					Users: users,
				},
			}

			g.It("- should return no error with GetAll()", func() {
				actual, resp, err := c.Permissions.GetAll()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Get()", func() {
				actual, resp, err := c.Permissions.Get("populateCaches")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Create()", func() {
				actual, resp, err := c.Permissions.Create(target)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Update()", func() {
				actual, resp, err := c.Permissions.Update(target)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Delete()", func() {
				actual, resp, err := c.Permissions.Delete("populateCaches")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
