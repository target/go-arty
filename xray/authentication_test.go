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
	"testing"

	"github.com/franela/goblin"
)

func Test_Authentication(t *testing.T) {
	// Create the client to interact with the Authentication service
	c, _ := NewClient("http://localhost:8080", nil)

	g := goblin.Goblin(t)
	g.Describe("Authentication Service", func() {

		g.Describe("Authentication", func() {
			g.It("- should set HTTP Basic auth with SetBasicAuth()", func() {
				c.Authentication.SetBasicAuth("user", "pass")
				g.Assert(c.Authentication.HasAuth()).IsTrue()
				g.Assert(c.Authentication.HasBasicAuth()).IsTrue()
			})

			g.It("- should set Token auth with SetTokenAuth()", func() {
				c.Authentication.SetTokenAuth("someToken")
				g.Assert(c.Authentication.HasAuth()).IsTrue()
				g.Assert(c.Authentication.HasTokenAuth()).IsTrue()
			})
		})

	})

}
