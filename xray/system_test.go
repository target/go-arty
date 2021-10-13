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
	"github.com/target/go-arty/v2/xray/fixtures/system"
)

func Test_System(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(system.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("System Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("System", func() {

			g.It("- should return valid string for Ping with String()", func() {
				actual := &Ping{
					Status: String("pong"),
				}

				data, _ := ioutil.ReadFile("fixtures/system/ping.json")

				var expected Ping
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return valid string for Versions with String()", func() {
				actual := &Versions{
					XrayVersion:  String("1.4"),
					XrayRevision: String("b3034"),
				}

				data, _ := ioutil.ReadFile("fixtures/system/version.json")

				var expected Versions
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with Ping()", func() {
				actual, resp, err := c.System.Ping()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Version()", func() {
				actual, resp, err := c.System.Version()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
