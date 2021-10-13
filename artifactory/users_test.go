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
	"github.com/target/go-arty/v2/artifactory/fixtures/users"
)

func Test_Users(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(users.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Users Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Users", func() {
			user := &User{}
			suser := &SecurityUser{}

			g.BeforeEach(func() {
				suser = &SecurityUser{
					Name:                     String("admin"),
					Email:                    String("admin@company.com"),
					Password:                 String("somepass"),
					Admin:                    Bool(true),
					ProfileUpdatable:         Bool(true),
					DisableUIAccess:          Bool(false),
					InternalPasswordDisabled: Bool(false),
					Groups:                   &[]string{"administrators"},
					LastLoggedIn:             String("2015-08-11T14:04:11.472Z"),
					Realm:                    String("internal"),
				}
				user = &User{
					Name:                     String("admin"),
					Email:                    String("admin@company.com"),
					Admin:                    Bool(true),
					GroupAdmin:               Bool(false),
					ProfileUpdatable:         Bool(true),
					InternalPasswordDisabled: Bool(false),
					Groups:                   &[]string{"administrators"},
					LastLoggedIn:             String("2018-11-15 15:17:09 +00:00"),
					LastLoggedInMillis:       Int64(1542295029877),
					Realm:                    String("ldap"),
					OfflineMode:              Bool(false),
					DisableUIAccess:          Bool(false),
					ProWithoutLicense:        Bool(false),
					ExternalRealmLink:        String("Check external status"),
					ExistsInDB:               Bool(false),
					HideUploads:              Bool(false),
					RequireProfileUnlock:     Bool(false),
					RequireProfilePassword:   Bool(false),
					Locked:                   Bool(false),
					CredentialsExpired:       Bool(false),
					NumberOfGroups:           Int(1),
					NumberOfPermissions:      Int(0),
				}
			})

			g.It("- should return valid string for User with String()", func() {
				data, _ := ioutil.ReadFile("fixtures/users/user.json")

				var expected User
				_ = json.Unmarshal(data, &expected)

				g.Assert(user.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for SecurityUser with String()", func() {
				data, _ := ioutil.ReadFile("fixtures/users/security_user.json")

				var expected SecurityUser
				_ = json.Unmarshal(data, &expected)

				g.Assert(suser.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with GetAll()", func() {
				actual, resp, err := c.Users.GetAll()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetAllSecurity()", func() {
				actual, resp, err := c.Users.GetAllSecurity()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetSecurity()", func() {
				actual, resp, err := c.Users.GetSecurity("admin")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with CreateSecurity()", func() {
				actual, resp, err := c.Users.CreateSecurity(suser)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with UpdateSecurity()", func() {
				actual, resp, err := c.Users.UpdateSecurity(suser)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with DeleteSecurity()", func() {
				actual, resp, err := c.Users.DeleteSecurity("admin")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

		g.Describe("API Keys", func() {

			g.It("- should return valid string for APIKey with String()", func() {
				actual := &APIKey{
					APIKey: String("AreallyCOOLsuperSECRETapiKEYthatNOBODYknows"),
				}

				data, _ := ioutil.ReadFile("fixtures/users/get_api_key.json")

				var expected APIKey
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for DeleteAPIKey with String()", func() {
				actual := &DeleteAPIKey{
					Info: String("Api key for user: 'admin' has been successfully revoked"),
				}

				data, _ := ioutil.ReadFile("fixtures/users/delete_api_key.json")

				var expected DeleteAPIKey
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with GetAPIKey()", func() {
				actual, resp, err := c.Users.GetAPIKey()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with CreateAPIKey()", func() {
				actual, resp, err := c.Users.CreateAPIKey()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with RegenerateAPIKey()", func() {
				actual, resp, err := c.Users.RegenerateAPIKey()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with DeleteAPIKey()", func() {
				actual, resp, err := c.Users.DeleteAPIKey()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with DeleteUserAPIKey()", func() {
				actual, resp, err := c.Users.DeleteUserAPIKey("admin")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with DeleteAllAPIKeys()", func() {
				actual, resp, err := c.Users.DeleteAllAPIKeys()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetEncryptedPassword()", func() {
				actual, resp, err := c.Users.GetEncryptedPassword()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
