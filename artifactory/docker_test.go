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

	"github.com/target/go-arty/artifactory/fixtures/docker"
	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
)

func Test_Docker(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(docker.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Docker Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Docker", func() {

			promotion := ImagePromotion{
				TargetRepo:             "docker",
				DockerRepository:       "docker-dev",
				TargetDockerRepository: "docker-test",
				Tag:       "latest",
				TargetTag: "latest",
				Copy:      true,
			}

			g.It("- should return no error with GetRepositories()", func() {
				actual, resp, err := c.Docker.GetRepositories("docker")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetTags()", func() {
				actual, resp, err := c.Docker.GetTags("docker", "docker-dev")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with PromoteImage()", func() {
				actual, resp, err := c.Docker.PromoteImage("docker", promotion)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})
		})

	})

}
