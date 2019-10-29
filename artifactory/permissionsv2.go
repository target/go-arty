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
	"fmt"
	"net/http"
)

// PermissionsServiceV2 handles communication with the permissions related
// methods of the Artifactory API v2.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API+V2
type PermissionsServiceV2 service

// Exists validates if the specific permission target exists.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API+V2#ArtifactoryRESTAPIV2-PermissionTargetexistencecheck
func (s *PermissionsServiceV2) Exists(target string) (bool, error) {
	u := fmt.Sprintf("/api/v2/security/permissions/%s", target)
	resp, err := s.client.Call("HEAD", u, nil, nil)
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
