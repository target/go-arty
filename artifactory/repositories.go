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
	PackageType *string `json:"packageType,omitempty"`
}

func (r Repository) String() string {
	return Stringify(r)
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

func (g GenericRepository) String() string {
	return Stringify(g)
}

// LocalRepository represents a local repository in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON#RepositoryConfigurationJSON-application/vnd.org.jfrog.artifactory.repositories.LocalRepositoryConfiguration+json
type LocalRepository struct {
	*GenericRepository

	DebianTrivialLayout             *bool     `json:"debianTrivialLayout,omitempty"`
	ChecksumPolicyType              *string   `json:"checksumPolicyType,omitempty"`
	MaxUniqueTags                   *int      `json:"maxUniqueTags,omitempty"`
	SnapshotVersionBehavior         *string   `json:"snapshotVersionBehavior,omitempty"`
	ArchiveBrowsingEnabled          *bool     `json:"archiveBrowsingEnabled,omitempty"`
	CalculateYumMetadata            *bool     `json:"calculateYumMetadata,omitempty"`
	YumRootDepth                    *int      `json:"yumRootDepth,omitempty"`
	DockerAPIVersion                *string   `json:"dockerApiVersion,omitempty"`
	EnableFileListsIndexing         *bool     `json:"enableFileListsIndexing,omitempty"`
	OptionalIndexCompressionFormats *[]string `json:"optionalIndexCompressionFormats,omitempty"`
	XrayIndex                       *bool     `json:"xrayIndex,omitempty"`
	DownloadRedirect                *bool     `json:"downloadRedirect,omitempty"`
}

func (l LocalRepository) String() string {
	return Stringify(l)
}

// RemoteRepository represents a remote repository in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON#RepositoryConfigurationJSON-application/vnd.org.jfrog.artifactory.repositories.RemoteRepositoryConfiguration+json
type RemoteRepository struct {
	*GenericRepository

	URL                               *string                 `json:"url,omitempty"`
	Username                          *string                 `json:"username,omitempty"`
	Password                          *string                 `json:"password,omitempty"`
	Proxy                             *string                 `json:"proxy,omitempty"`
	RemoteRepoChecksumPolicyType      *string                 `json:"remoteRepoChecksumPolicyType,omitempty"`
	HardFail                          *bool                   `json:"hardFail,omitempty"`
	Offline                           *bool                   `json:"offline,omitempty"`
	StoreArtifactsLocally             *bool                   `json:"storeArtifactsLocally,omitempty"`
	SocketTimeoutMillis               *int                    `json:"socketTimeoutMillis,omitempty"`
	LocalAddress                      *string                 `json:"localAddress,omitempty"`
	RetrievalCachePeriodSecs          *int                    `json:"retrievalCachePeriodSecs,omitempty"`
	FailedRetrievalCachePeriodSecs    *int                    `json:"failedRetrievalCachePeriodSecs,omitempty"`
	MissedRetrievalCachePeriodSecs    *int                    `json:"missedRetrievalCachePeriodSecs,omitempty"`
	UnusedArtifactsCleanupEnabled     *bool                   `json:"unusedArtifactsCleanupEnabled,omitempty"`
	UnusedArtifactsCleanupPeriodHours *int                    `json:"unusedArtifactsCleanupPeriodHours,omitempty"`
	AssumedOfflinePeriodSecs          *int                    `json:"assumedOfflinePeriodSecs,omitempty"`
	FetchJarsEagerly                  *bool                   `json:"fetchJarsEagerly,omitempty"`
	FetchSourcesEagerly               *bool                   `json:"fetchSourcesEagerly,omitempty"`
	ShareConfiguration                *bool                   `json:"shareConfiguration,omitempty"`
	SynchronizeProperties             *bool                   `json:"synchronizeProperties,omitempty"`
	BlockMismatchingMimeTypes         *bool                   `json:"blockMismatchingMimeTypes,omitempty"`
	AllowAnyHostAuth                  *bool                   `json:"allowAnyHostAuth,omitempty"`
	EnableCookieManagement            *bool                   `json:"enableCookieManagement,omitempty"`
	BowerRegistryURL                  *string                 `json:"bowerRegistryUrl,omitempty"`
	ComposerRegistryURL               *string                 `json:"composerRegistryUrl,omitempty"`
	PyPIRegistryURL                   *string                 `json:"pyPIRegistryUrl,omitempty"`
	VcsType                           *string                 `json:"vcsType,omitempty"`
	VcsGitProvider                    *string                 `json:"vcsGitProvider,omitempty"`
	VcsGitDownloadUrl                 *string                 `json:"VcsGitDownloadUrl,omitempty"`
	BypassHeadRequests                *bool                   `json:"bypassHeadRequests,omitempty"`
	ClientTLSCertificate              *string                 `json:"clientTlsCertificate,omitempty"`
	ExternalDependenciesEnabled       *bool                   `json:"externalDependenciesEnabled,omitempty"`
	ExternalDependenciesPatterns      *[]string               `json:"externalDependenciesPatterns,omitempty"`
	DownloadRedirect                  *bool                   `json:"downloadRedirect,omitempty"`
	FeedContextPath                   *string                 `json:"feedContextPath,omitempty"`
	DownloadContextPath               *string                 `json:"downloadContextPath,omitempty"`
	V3FeedUrl                         *string                 `json:"v3FeedUrl,omitempty"`
	XrayIndex                         *bool                   `json:"xrayIndex,omitempty"`
	ListRemoteFolderItems             *bool                   `json:"listRemoteFolderItems,omitempty"`
	EnableTokenAuthentication         *bool                   `json:"enableTokenAuthentication,omitempty"`
	ContentSynchronisation            *ContentSynchronisation `json:"contentSynchronisation,omitempty"`
}

// ContentSynchronisation represents smart remote repository configuration
type ContentSynchronisation struct {
	Enabled    *bool `json:"enabled,omitempty"`
	Properties *struct {
		Enabled *bool `json:"enabled,omitempty"`
	} `json:"properties,omitempty"`
	Statistics *struct {
		Enabled *bool `json:"enabled,omitempty"`
	} `json:"statistics,omitempty"`
	Source *struct {
		OriginAbsenceDetection *bool `json:"originAbsenceDetection,omitempty"`
	} `json:"source,omitempty"`
}

func (r RemoteRepository) String() string {
	return Stringify(r)
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
	ForceMavenAuthentication                      *bool     `json:"forceMavenAuthentication,omitempty"`
	ExternalDependenciesEnabled                   *bool     `json:"externalDependenciesEnabled,omitempty"`
	ExternalDependenciesPatterns                  *[]string `json:"externalDependenciesPatterns,omitempty"`
	ExternalDependenciesRemoteRepo                *string   `json:"externalDependenciesRemoteRepo,omitempty"`
}

func (v VirtualRepository) String() string {
	return Stringify(v)
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
