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

import "github.com/coreos/go-semver/semver"

var (
	// VersionMajor is for an API incompatible changes
	VersionMajor int64 = 1
	// VersionMinor is for functionality in a backwards-compatible manner
	VersionMinor int64 = 12
	// VersionPatch is for backwards-compatible bug fixes
	VersionPatch int64 = 1
)

// Version represents the minimum version of the Artifactory API this library supports
var Version = semver.Version{
	Major: VersionMajor,
	Minor: VersionMinor,
	Patch: VersionPatch,
}
