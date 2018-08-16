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

import "fmt"

// UsersService handles communication with the user related
// methods of the Xray API.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-SUMMARY
type UsersService service

// User represents a user in Xray.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-USERMANAGEMENT
type User struct {
	Admin    *bool   `json:"admin,omitempty"`
	Email    *string `json:"email,omitempty"`
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

func (u User) String() string {
	return Stringify(u)
}

// GetAll returns a list of all users.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-GetUsers/GetUser
func (s *UsersService) GetAll() (*[]User, *Response, error) {
	u := "/api/v1/users"
	v := new([]User)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns the provided user.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-GetUsers/GetUser
func (s *UsersService) Get(user string) (*User, *Response, error) {
	u := fmt.Sprintf("/api/v1/users/%s", user)
	v := new(User)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Create constructs a new User with the provided details.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-CreateUser
func (s *UsersService) Create(user *User) (*User, *Response, error) {
	u := "/api/v1/users"
	v := new(User)

	resp, err := s.client.Call("POST", u, user, v)
	return v, resp, err
}

// Update modifies a user with the provided details.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-UpdateUser
func (s *UsersService) Update(user *User) (*string, *Response, error) {
	u := fmt.Sprintf("/api/v1/users/%s", *user.Name)
	v := new(string)

	resp, err := s.client.Call("PUT", u, user, v)
	return v, resp, err
}

// Delete removes the provided user.
//
// Docs: https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-DeleteUser
func (s *UsersService) Delete(user string) (*string, *Response, error) {
	u := fmt.Sprintf("/api/v1/users/%s", user)
	v := new(string)

	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}
