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

	"github.com/target/go-arty/artifactory/fixtures/licenses"
	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
)

func Test_Licenses(t *testing.T) {
	gin.SetMode(gin.TestMode)

	g := goblin.Goblin(t)
	g.Describe("Licenses Service", func() {

		g.Describe("Licenses", func() {
			// Create http test server from our fake API handler
			s := httptest.NewServer(licenses.FakeHandler())

			// Create the client to interact with the http test server
			c, _ := NewClient(s.URL, nil)

			// Close http test server after we're done using it
			g.After(func() {
				s.Close()
			})

			license := LicenseRequest{
				LicenseKey: "179b7ea384d0c4655a00dfac7285a21d986a17923",
			}

			g.It("- should return no error with Get()", func() {
				actual, resp, err := c.Licenses.Get()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Install()", func() {
				actual, resp, err := c.Licenses.Install(license)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

		g.Describe("HA Licenses", func() {
			// Create http test server from our fake API handler
			s := httptest.NewServer(licenses.FakeHAHandler())

			// Create the client to interact with the http test server
			c, _ := NewClient(s.URL, nil)

			// Close http test server after we're done using it
			g.After(func() {
				s.Close()
			})

			licenses := []LicenseRequest{
				LicenseRequest{LicenseKey: "179b7ea384d0c4655a00dfac7285a21d986a17923"},
			}

			licenseHashes := LicenseRemoval{
				LicenseHashes: []string{"179b7ea384d0c4655a00dfac7285a21d986a17923"},
			}

			g.It("- should return no error with GetHA()", func() {
				actual, resp, err := c.Licenses.GetHA()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with InstallHA()", func() {
				actual, resp, err := c.Licenses.InstallHA(licenses)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with DeleteHA()", func() {
				actual, resp, err := c.Licenses.DeleteHA(licenseHashes)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
