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

const (
	// HTTP Basic Authentication
	authTypeBasic = 1
	// Auth via API Token
	authTypeToken = 2
)

// AuthenticationService contains authentication related functions.
type AuthenticationService struct {
	client   *Client
	username string
	secret   string
	authType int
}

// SetBasicAuth sets the auth type as HTTP Basic auth.
func (s *AuthenticationService) SetBasicAuth(username, password string) {
	s.username = username
	s.secret = password
	s.authType = authTypeBasic
}

// SetTokenAuth sets the auth type as Token auth.
func (s *AuthenticationService) SetTokenAuth(token string) {
	s.secret = token
	s.authType = authTypeToken
}

// HasAuth checks if the auth type is set.
func (s *AuthenticationService) HasAuth() bool {
	return s.authType > 0
}

// HasBasicAuth checks if the auth type is HTTP Basic auth.
func (s *AuthenticationService) HasBasicAuth() bool {
	return s.authType == authTypeBasic
}

// HasTokenAuth checks if the auth type is Token auth.
func (s *AuthenticationService) HasTokenAuth() bool {
	return s.authType == authTypeToken
}
