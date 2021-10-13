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
	"github.com/target/go-arty/artifactory/fixtures/groups"
)

func Test_Groups(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(groups.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Groups Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Groups", func() {
			group := &Group{}

			g.BeforeEach(func() {
				group = &Group{
					Name:            String("dev-leads"),
					Description:     String("The development leads group"),
					AutoJoin:        Bool(false),
					AdminPrivileges: Bool(false),
					Realm:           String("ldap"),
					RealmAttributes: String("Realm attributes for use by LDAP"),
					UserNames:       &[]string{"user1", "user2", "user3"},
				}
			})

			g.It("- should return valid string for Group with String()", func() {
				data, _ := ioutil.ReadFile("fixtures/groups/group.json")

				var expected Group
				_ = json.Unmarshal(data, &expected)

				g.Assert(group.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with GetAll()", func() {
				actual, resp, err := c.Groups.GetAll()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Get() with IncludeUsers true", func() {
				groupRequest := &GetGroupRequest{
					Name:         String("dev-leads"),
					IncludeUsers: Bool(true),
				}

				actual, resp, err := c.Groups.Get(groupRequest)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Get() with IncludeUsers nil", func() {
				groupRequest := &GetGroupRequest{
					Name:         String("dev-leads"),
					IncludeUsers: Bool(false),
				}

				actual, resp, err := c.Groups.Get(groupRequest)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Create()", func() {
				actual, resp, err := c.Groups.Create(group)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Update()", func() {
				actual, resp, err := c.Groups.Update(group)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Delete()", func() {
				actual, resp, err := c.Groups.Delete("test")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
