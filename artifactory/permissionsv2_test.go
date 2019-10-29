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
		})
	})
}
