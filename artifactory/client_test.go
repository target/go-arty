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
	"net/http"
	"testing"

	"github.com/franela/goblin"
)

func Test_client(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("NewClient", func() {

		g.It("- should be able to create new client", func() {
			actual, err := NewClient("https://some.company.com", nil)

			g.Assert(actual != nil).IsTrue()
			g.Assert(err == nil).IsTrue()
		})

		g.It("- should fail to create new client with empty url", func() {
			actual, err := NewClient("", nil)

			g.Assert(actual == nil).IsTrue()
			g.Assert(err != nil).IsTrue()
		})

		g.It("- should fail to create new client with bad url", func() {
			actual, err := NewClient("@$(:LKS24poihwekf1203", nil)

			g.Assert(actual == nil).IsTrue()
			g.Assert(err != nil).IsTrue()
		})
	})

	g.Describe("BuildURLForRequest", func() {
		client, _ := NewClient("https://some.company.com", nil)

		g.BeforeEach(func() {
			client, _ = NewClient("https://some.company.com", nil)
		})

		g.It("- should be able to build url for request", func() {
			actual, err := client.buildURLForRequest("test")

			g.Assert(actual != "").IsTrue()
			g.Assert(err == nil).IsTrue()
		})

		g.It("- should be able to build url for request", func() {
			actual, err := client.buildURLForRequest("/test")

			g.Assert(actual != "").IsTrue()
			g.Assert(err == nil).IsTrue()
		})

		g.It("- should be able to build url for request", func() {
			actual, err := client.buildURLForRequest("test/")

			g.Assert(actual != "").IsTrue()
			g.Assert(err == nil).IsTrue()
		})

		g.It("- should fail to build url for request", func() {
			actual, err := client.buildURLForRequest("@$(:LKS24poihwekf1203")

			g.Assert(actual == "").IsTrue()
			g.Assert(err != nil).IsTrue()
		})
	})

	g.Describe("AddAuthentication", func() {
		client, _ := NewClient("https://some.company.com", nil)
		request, _ := http.NewRequest("GET", "https://some.company.com/ping", nil)

		g.BeforeEach(func() {
			client, _ = NewClient("https://some.company.com", nil)
			request, _ = http.NewRequest("GET", "https://some.company.com/ping", nil)
		})

		g.It("- should be able to add authentication for request", func() {
			client.Authentication.SetBasicAuth("user", "pass")
			client.addAuthentication(request)

			user, pass, _ := request.BasicAuth()

			g.Assert(user).Equal("user")
			g.Assert(pass).Equal("pass")
		})

		g.It("- should be able to add token authentication for request", func() {
			client.Authentication.SetTokenAuth("someToken")
			client.addAuthentication(request)

			g.Assert(request.Header.Get("X-JFrog-Art-Api")).Equal("someToken")
		})
	})

	g.Describe("NewRequest", func() {
		client, _ := NewClient("https://some.company.com", nil)

		g.BeforeEach(func() {
			client, _ = NewClient("https://some.company.com", nil)
		})

		g.It("- should be able to create a new request", func() {
			actual, err := client.NewRequest("GET", "/ping", nil)

			g.Assert(actual != nil).IsTrue()
			g.Assert(err == nil).IsTrue()
		})

		g.It("- should fail to create a new request with bad url", func() {
			actual, err := client.NewRequest("GET", "@$(:LKS24poihwekf1203", nil)

			g.Assert(actual == nil).IsTrue()
			g.Assert(err != nil).IsTrue()
		})
	})

	g.Describe("addOptions", func() {
		type options struct {
			ShowAll bool `url:"all"`
			Page    int  `url:"page"`
		}
		g.It("- call addOptions - valid case", func() {
			options := options{ShowAll: true, Page: 1}
			actual, err := addOptions("http://www.somestring.com", options)

			g.Assert(actual).Equal("http://www.somestring.com?all=true&page=1")
			g.Assert(err == nil).IsTrue()
		})

		g.It("- call addOptions - valid url with bad options type", func() {
			actual, err := addOptions("http://www.somestring.com", 87)

			g.Assert(actual).Equal("http://www.somestring.com")
			g.Assert(err != nil).IsTrue()
			g.Assert(err.Error()).Equal("query: Values() expects struct input. Got int")
		})

		g.It("- call addOptions - bad url parse", func() {
			options := options{ShowAll: true, Page: 1}
			_, err := addOptions("!@*&^%%", options)
			g.Assert(err != nil).IsTrue()
			g.Assert(err.Error()).Equal("parse !@*&^%%: invalid URL escape \"%%\"")
		})

		g.It("- call addOptions - nil options passed", func() {
			options := (*options)(nil)
			actual, err := addOptions("http://www.somestring.com", options)
			g.Assert(actual).Equal("http://www.somestring.com")
			g.Assert(err == nil).IsTrue()
		})
	})

}
