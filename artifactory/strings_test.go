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

package artifactory

import (
	"fmt"
	"testing"
	"time"
)

func TestStringify(t *testing.T) {
	var nilPointer *string

	var tests = []struct {
		in  interface{}
		out string
	}{
		// basic types
		{"foo", `"foo"`},
		{123, `123`},
		{1.5, `1.5`},
		{false, `false`},
		{
			[]string{"a", "b"},
			`["a" "b"]`,
		},
		{
			struct {
				A []string
			}{nil},
			// nil slice is skipped
			`{}`,
		},
		{
			struct {
				A string
			}{"foo"},
			// structs not of a named type get no prefix
			`{A:"foo"}`,
		},

		// pointers
		{nilPointer, `<nil>`},
		{String("foo"), `"foo"`},
		{Int(123), `123`},
		{Bool(false), `false`},
		{
			[]*string{String("a"), String("b")},
			`["a" "b"]`,
		},

		// actual Artifactory structs
		{
			Timestamp{time.Date(2006, time.January, 02, 15, 04, 05, 0, time.UTC)},
			`artifactory.Timestamp{2006-01-02 15:04:05 +0000 UTC}`,
		},
		{
			User{Name: String("test"), Email: String("test@company.com")},
			`artifactory.User{Name:"test", Email:"test@company.com"}`,
		},
		{
			GenericRepository{Key: String("test"), RClass: String("")},
			`artifactory.GenericRepository{Key:"test", RClass:""}`,
		},
	}

	for i, tt := range tests {
		s := Stringify(tt.in)
		if s != tt.out {
			t.Errorf("%d. Stringify(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}

// Directly test the String() methods on various Artifactory types. We don't do an
// exaustive test of all the various field types, since TestStringify() above
// takes care of that. Rather, we just make sure that Stringify() is being
// used to build the strings, which we do by verifying that pointers are
// stringified as their underlying value.
func TestString(t *testing.T) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		{ArtifactMessage{Level: String("test")}, `artifactory.ArtifactMessage{Level:"test"}`},
		{Registry{Repositories: &[]string{"test"}}, `artifactory.Registry{Repositories:["test"]}`},
		{Tags{Name: String("test")}, `artifactory.Tags{Name:"test"}`},
		{Group{Name: String("test")}, `artifactory.Group{Name:"test"}`},
		{License{Type: String("Commercial")}, `artifactory.License{Type:"Commercial"}`},
		{HALicense{Type: String("Commercial")}, `artifactory.HALicense{Type:"Commercial"}`},
		{LicenseResponse{Status: Int(200)}, `artifactory.LicenseResponse{Status:200}`},
		{HALicenseResponse{Status: Int(200)}, `artifactory.HALicenseResponse{Status:200}`},
		{PermissionTarget{Name: String("test")}, `artifactory.PermissionTarget{Name:"test"}`},
		{Repository{Key: String("test")}, `artifactory.Repository{Key:"test"}`},
		{GenericRepository{Key: String("test")}, `artifactory.GenericRepository{Key:"test"}`},
		{LocalRepository{ChecksumPolicyType: String("test")}, `artifactory.LocalRepository{ChecksumPolicyType:"test"}`},
		{RemoteRepository{URL: String("test")}, `artifactory.RemoteRepository{URL:"test"}`},
		{VirtualRepository{KeyPair: String("test")}, `artifactory.VirtualRepository{KeyPair:"test"}`},
		{Folder{URI: String("test")}, `artifactory.Folder{URI:"test"}`},
		{File{URI: String("test")}, `artifactory.File{URI:"test"}`},
		{ItemLastModified{URI: String("test")}, `artifactory.ItemLastModified{URI:"test"}`},
		{FileStatistics{URI: String("test")}, `artifactory.FileStatistics{URI:"test"}`},
		{ItemProperties{URI: String("test")}, `artifactory.ItemProperties{URI:"test"}`},
		{FileList{URI: String("test")}, `artifactory.FileList{URI:"test"}`},
		{Versions{Version: String("test")}, `artifactory.Versions{Version:"test"}`},
		{Timestamp{time.Date(2006, time.January, 02, 15, 04, 05, 0, time.UTC)}, `2006-01-02 15:04:05 +0000 UTC`},
		{User{Name: String("test")}, `artifactory.User{Name:"test"}`},
		{APIKey{APIKey: String("test")}, `artifactory.APIKey{APIKey:"test"}`},
		{DeleteAPIKey{Info: String("test")}, `artifactory.DeleteAPIKey{Info:"test"}`},
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
