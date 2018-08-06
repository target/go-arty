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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	userAgent = "go-arty"
)

// Client is a client that manages communication with the Xray API.
type Client struct {
	// HTTP client used to communicate with the Xray API.
	client *http.Client

	// Base URL for Xray API requests.
	baseURL *url.URL

	// User agent used when communicating with the Xray API.
	UserAgent string

	// Xray service for authentication.
	Authentication *AuthenticationService
	Scan           *ScanService
	Summary        *SummaryService
	System         *SystemService
	Users          *UsersService
}

type service struct {
	client *Client
}

// NewClient returns a new Xray API client.
// baseUrl has to be the HTTP endpoint of the Xray API.
// If no httpClient is provided, then the http.DefaultClient will be used.
func NewClient(baseUrl string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("No Xray baseUrl provided")
	}
	baseURL, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:    httpClient,
		baseURL:   baseURL,
		UserAgent: userAgent,
	}

	c.Authentication = &AuthenticationService{client: c}
	c.Scan = &ScanService{client: c}
	c.Summary = &SummaryService{client: c}
	c.System = &SystemService{client: c}
	c.Users = &UsersService{client: c}

	return c, nil
}

// buildURLForRequest will build the URL (as a string) that will be called.
// It does several cleaning tasks for us.
func (c *Client) buildURLForRequest(urlStr string) (string, error) {
	u := c.baseURL.String()

	// If there is no / at the end, add one.
	if strings.HasSuffix(u, "/") == false {
		u += "/"
	}

	// If there is a "/" at the start, remove it.
	if strings.HasPrefix(urlStr, "/") == true {
		urlStr = urlStr[1:]
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	u += rel.String()

	return u, nil
}

// addAuthentication adds the necessary authentication to the request.
func (c *Client) addAuthentication(req *http.Request) {
	// Apply HTTP Basic Authentication.
	if c.Authentication.HasBasicAuth() {
		req.SetBasicAuth(*c.Authentication.username, *c.Authentication.secret)
	}

	if c.Authentication.HasTokenAuth() {
		q := req.URL.Query()
		q.Add("token", *c.Authentication.secret)
		req.URL.RawQuery = q.Encode()
	}
}

// NewRequest creates an API request.
// A relative URL can be provided in urlStr,
// in which case it is resolved relative to the baseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.buildURLForRequest(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	// Apply Authentication.
	if c.Authentication.HasAuth() {
		c.addAuthentication(req)
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// Response represents an Xray API response.
// This wraps the standard http.Response returned from Xray.
type Response struct {
	*http.Response
}

// Call is a combine function for Client.NewRequest and Client.Do.
//
// Most API methods are quite the same.
// Get the URL, apply options, make a request, and get the response.
// Without adding special headers or something.
// To avoid a big amount of code duplication you can Client.Call.
//
// method is the HTTP method you want to call.
// u is the URL you want to call.
// body is the HTTP body.
// v is the HTTP response.
//
// For more information read https://github.com/google/go-github/issues/234
func (c *Client) Call(method, u string, body interface{}, v interface{}) (*Response, error) {
	req, err := c.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, v)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v,
// or returned as an error if an API error has occurred.
// If v implements the io.Writer interface, the raw response body will be written to v,
// without attempting to first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Wrap response
	response := &Response{Response: resp}

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			var body []byte
			body, err = ioutil.ReadAll(resp.Body)
			// This ensures the response body is not empty in the event the user
			// wants to inspect the response body further
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			if err != nil {
				// even though there was an error, we still return the response
				// in case the caller wants to inspect it further
				return response, err
			}
			// Since the API we integrate with doesn't always return JSON
			// we no longer explicitly return an error if we can't unmarshal the body
			_ = json.Unmarshal(body, v)
		}
	}
	return response, err
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if it has a status code outside the 200 range.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return fmt.Errorf("API call to %s failed: %s", r.Request.URL.String(), r.Status)
}
