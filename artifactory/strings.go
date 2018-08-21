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
	"bytes"
	"fmt"
	"io"

	"reflect"
)

var timestampType = reflect.TypeOf(Timestamp{})

// Stringify attempts to create a reasonable string representation of types in
// the Artifactory library. It does things like resolve pointers to their values
// and omits struct fields with nil values.
func Stringify(message interface{}) string {
	var buf bytes.Buffer
	v := reflect.ValueOf(message)
	stringifyValue(&buf, v)
	return buf.String()
}

// stringifyValue was heavily inspired by the goprotobuf library.

func stringifyValue(w io.Writer, val reflect.Value) {
	if val.Kind() == reflect.Ptr && val.IsNil() {
		w.Write([]byte("<nil>"))
		return
	}

	v := reflect.Indirect(val)

	switch v.Kind() {
	case reflect.String:
		fmt.Fprintf(w, `"%s"`, v)
	case reflect.Slice:
		w.Write([]byte{'['})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				w.Write([]byte{' '})
			}

			stringifyValue(w, v.Index(i))
		}

		w.Write([]byte{']'})
		return
	case reflect.Struct:
		if v.Type().Name() != "" {
			w.Write([]byte(v.Type().String()))
		}

		// special handling of Timestamp values
		if v.Type() == timestampType {
			fmt.Fprintf(w, "{%s}", v.Interface())
			return
		}

		w.Write([]byte{'{'})

		var sep bool
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			if fv.Kind() == reflect.Ptr && fv.IsNil() {
				continue
			}
			if fv.Kind() == reflect.Slice && fv.IsNil() {
				continue
			}

			if sep {
				w.Write([]byte(", "))
			} else {
				sep = true
			}

			w.Write([]byte(v.Type().Field(i).Name))
			w.Write([]byte{':'})
			stringifyValue(w, fv)
		}

		w.Write([]byte{'}'})
	default:
		if v.CanInterface() {
			fmt.Fprint(w, v.Interface())
		}
	}
}
