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

		// actual Xray structs
		{
			Timestamp{time.Date(2006, time.January, 02, 15, 04, 05, 0, time.UTC)},
			`xray.Timestamp{2006-01-02 15:04:05 +0000 UTC}`,
		},
		{
			User{Name: String("test"), Email: String("test@company.com")},
			`xray.User{Email:"test@company.com", Name:"test"}`,
		},
		{
			Ping{Status: String("test")},
			`xray.Ping{Status:"test"}`,
		},
	}

	for i, tt := range tests {
		s := Stringify(tt.in)
		if s != tt.out {
			t.Errorf("%d. Stringify(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}

// Directly test the String() methods on various Xray types. We don't do an
// exhaustive test of all the various field types, since TestStringify() above
// takes care of that. Rather, we just make sure that Stringify() is being
// used to build the strings, which we do by verifying that pointers are
// stringified as their underlying value.
func TestString(t *testing.T) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		{ScanArtifactResponse{Info: String("test")}, `xray.ScanArtifactResponse{Info:"test"}`},
		{ScanBuildResponse{Summary: &ScanSummary{Message: String("test")}}, `xray.ScanBuildResponse{Summary:xray.ScanSummary{Message:"test"}}`},
		{SummaryResponse{Errors: &[]SummaryError{SummaryError{Error: String("test")}}}, `xray.SummaryResponse{Errors:[xray.SummaryError{Error:"test"}]}`},
		{Ping{Status: String("test")}, `xray.Ping{Status:"test"}`},
		{Versions{XrayVersion: String("test")}, `xray.Versions{XrayVersion:"test"}`},
		{Timestamp{time.Date(2006, time.January, 02, 15, 04, 05, 0, time.UTC)}, `2006-01-02 15:04:05 +0000 UTC`},
		{User{Name: String("test")}, `xray.User{Name:"test"}`},
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
