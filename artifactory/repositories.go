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
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tidwall/gjson"
)

// RepositoriesService handles communication with the repository related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-REPOSITORIES
type RepositoriesService service

// Repository represents a repository in Artifactory.
type Repository struct {
	Key         *string `json:"key,omitempty"`
	Type        *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// GenericRepository represents the common json across all repository types from Artifactory.
type GenericRepository struct {
	Key                          *string   `json:"key,omitempty"`
	RClass                       *string   `json:"rclass,omitempty"`
	PackageType                  *string   `json:"packageType,omitempty"`
	Description                  *string   `json:"description,omitempty"`
	Notes                        *string   `json:"notes,omitempty"`
	IncludesPattern              *string   `json:"includesPattern,omitempty"`
	ExcludesPattern              *string   `json:"excludesPattern,omitempty"`
	LayoutRef                    *string   `json:"repoLayoutRef,omitempty"`
	HandleReleases               *bool     `json:"handleReleases,omitempty"`
	HandleSnapshots              *bool     `json:"handleSnapshots,omitempty"`
	MaxUniqueSnapshots           *int      `json:"maxUniqueSnapshots,omitempty"`
	SuppressPomConsistencyChecks *bool     `json:"suppressPomConsistencyChecks,omitempty"`
	BlackedOut                   *bool     `json:"blackedOut,omitempty"`
	PropertySets                 *[]string `json:"propertySets,omitempty"`
}

// LocalRepository represents a local repository in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON#RepositoryConfigurationJSON-application/vnd.org.jfrog.artifactory.repositories.LocalRepositoryConfiguration+json
type LocalRepository struct {
	*GenericRepository

	DebianTrivialLayout     *bool   `json:"debianTrivialLayout,omitempty"`
	ChecksumPolicyType      *string `json:"checksumPolicyType,omitempty"`
	MaxUniqueTags           *int    `json:"maxUniqueTags,omitempty"`
	SnapshotVersionBehavior *string `json:"snapshotVersionBehavior,omitempty"`
	ArchiveBrowsingEnabled  *bool   `json:"archiveBrowsingEnabled,omitempty"`
	CalculateYumMetadata    *bool   `json:"calculateYumMetadata,omitempty"`
	YumRootDepth            *int    `json:"yumRootDepth,omitempty"`
	DockerAPIVersion        *string `json:"dockerApiVersion,omitempty"`
	EnableFileListsIndexing *bool   `json:"enableFileListsIndexing,omitempty"`
}

// RemoteRepository represents a remote repository in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON#RepositoryConfigurationJSON-application/vnd.org.jfrog.artifactory.repositories.RemoteRepositoryConfiguration+json
type RemoteRepository struct {
	*GenericRepository

	URL                               *string `json:"url,omitempty"`
	Username                          *string `json:"username,omitempty"`
	Password                          *string `json:"password,omitempty"`
	Proxy                             *string `json:"proxy,omitempty"`
	RemoteRepoChecksumPolicyType      *string `json:"remoteRepoChecksumPolicyType,omitempty"`
	HardFail                          *bool   `json:"hardFail,omitempty"`
	Offline                           *bool   `json:"offline,omitempty"`
	StoreArtifactsLocally             *bool   `json:"storeArtifactsLocally,omitempty"`
	SocketTimeoutMillis               *int    `json:"socketTimeoutMillis,omitempty"`
	LocalAddress                      *string `json:"localAddress,omitempty"`
	RetrievalCachePeriodSecs          *int    `json:"retrievalCachePeriodSecs,omitempty"`
	FailedRetrievalCachePeriodSecs    *int    `json:"failedRetrievalCachePeriodSecs,omitempty"`
	MissedRetrievalCachePeriodSecs    *int    `json:"missedRetrievalCachePeriodSecs,omitempty"`
	UnusedArtifactsCleanupEnabled     *bool   `json:"unusedArtifactCleanupEnabled,omitempty"`
	UnusedArtifactsCleanupPeriodHours *int    `json:"unusedArtifactCleanupPeriodHours,omitempty"`
	FetchJarsEagerly                  *bool   `json:"fetchJarsEagerly,omitempty"`
	FetchSourcesEagerly               *bool   `json:"fetchSourcesEagerly,omitempty"`
	ShareConfiguration                *bool   `json:"shareConfiguration,omitempty"`
	SynchronizeProperties             *bool   `json:"synchronizeProperties,omitempty"`
	BlockMismatchingMimeTypes         *bool   `json:"blockMismatchingMimeTypes,omitempty"`
	AllowAnyHostAuth                  *bool   `json:"allowAnyHostAuth,omitempty"`
	EnableCookieManagement            *bool   `json:"enableCookieManagement,omitempty"`
	BowerRegistryURL                  *string `json:"bowerRegistryUrl,omitempty"`
	ComposerRegistryURL               *string `json:"composerRegistryUrl,omitempty"`
	PyPIRegistryURL                   *string `json:"pyPIRegistryUrl,omitempty"`
	VcsType                           *string `json:"vcsType,omitempty"`
	VcsGitProvider                    *string `json:"vcsGitProvider,omitempty"`
	VcsGitDownloader                  *string `json:"vcsGitDownloader,omitempty"`
	BypassHeadRequest                 *bool   `json:"bypassHeadRequest,omitempty"`
	ClientTLSCertificate              *string `json:"clientTlsCertificate,omitempty"`
}

// VirtualRepository represents a virtual repository in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON#RepositoryConfigurationJSON-application/vnd.org.jfrog.artifactory.repositories.VirtualRepositoryConfiguration+json
type VirtualRepository struct {
	*GenericRepository

	Repositories                                  *[]string `json:"repositories,omitempty"`
	DebianTrivialLayout                           *bool     `json:"debianTrivialLayout,omitempty"`
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool     `json:"artifactoryRequestsCanRetrieveRemoteArtifacts,omitempty"`
	KeyPair                                       *string   `json:"keyPair,omitempty"`
	PomRepositoryReferencesCleanupPolicy          *string   `json:"pomRepositoryReferencesCleanupPolicy,omitempty"`
	DefaultDeploymentRepo                         *string   `json:"defaultDeploymentRepo,omitempty"`
}

// GetAll returns a list of all repositories.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetRepositories
func (s *RepositoriesService) GetAll() (*[]Repository, *Response, error) {
	u := "/api/repositories"
	v := new([]Repository)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns the provided repository.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-RepositoryConfiguration
func (s *RepositoriesService) Get(repo string) (interface{}, *Response, error) {
	u := fmt.Sprintf("/api/repositories/%s", repo)
	v := new(GenericRepository)

	resp, err := s.client.Call("GET", u, nil, v)
	if err != nil {
		return v, resp, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	switch gjson.GetBytes(body, "rclass").Str {
	case "local":
		v := new(LocalRepository)
		_ = json.Unmarshal(body, v)
		return v, resp, err
	case "remote":
		v := new(RemoteRepository)
		_ = json.Unmarshal(body, v)
		return v, resp, err
	case "virtual":
		v := new(VirtualRepository)
		_ = json.Unmarshal(body, v)
		return v, resp, err
	default:
		return v, resp, err
	}
}

// Create constructs a repository with the provided details.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-CreateRepository
func (s *RepositoriesService) Create(repo string, body interface{}) (*string, *Response, error) {
	u := fmt.Sprintf("/api/repositories/%s", repo)
	v := new(string)

	resp, err := s.client.Call("PUT", u, body, v)
	return v, resp, err
}

// Update modifies a repository with the provided details.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateRepositoryConfiguration
func (s *RepositoriesService) Update(repo string, body interface{}) (*string, *Response, error) {
	u := fmt.Sprintf("/api/repositories/%s", repo)
	v := new(string)

	resp, err := s.client.Call("POST", u, body, v)
	return v, resp, err
}

// Delete removes the provided repository.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-DeleteRepository
func (s *RepositoriesService) Delete(repo string) (*string, *Response, error) {
	u := fmt.Sprintf("/api/repositories/%s", repo)
	v := new(string)

	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}
