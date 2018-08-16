// Copyright (c) 2013 The go-github AUTHORS. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Copyright (c) 2018 Target Brands, Inc.

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
