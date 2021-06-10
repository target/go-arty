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
	Key                          *string   `json:"key,omitempty" xml:"key,omitempty"`
	RClass                       *string   `json:"rclass,omitempty" xml:"-"`
	PackageType                  *string   `json:"packageType,omitempty" xml:"type,omitempty"`
	Description                  *string   `json:"description,omitempty" xml:"description,omitempty"`
	Notes                        *string   `json:"notes,omitempty" xml:"notes,omitempty"`
	IncludesPattern              *string   `json:"includesPattern,omitempty" xml:"includesPattern,omitempty"`
	ExcludesPattern              *string   `json:"excludesPattern,omitempty" xml:"excludesPattern,omitempty"`
	LayoutRef                    *string   `json:"repoLayoutRef,omitempty" xml:"repoLayoutRef,omitempty"`
	HandleReleases               *bool     `json:"handleReleases,omitempty" xml:"handleReleases,omitempty"`
	HandleSnapshots              *bool     `json:"handleSnapshots,omitempty" xml:"handleSnapshots,omitempty"`
	MaxUniqueSnapshots           *int      `json:"maxUniqueSnapshots,omitempty" xml:"maxUniqueSnapshots,omitempty"`
	SuppressPomConsistencyChecks *bool     `json:"suppressPomConsistencyChecks,omitempty" xml:"suppressPomConsistencyChecks,omitempty"`
	BlackedOut                   *bool     `json:"blackedOut,omitempty" xml:"blackedOut,omitempty"`
	PropertySets                 *[]string `json:"propertySets,omitempty" xml:"propertySets>propertySetRef,omitempty"`
	ForceNugetAuthentication     *bool     `json:"forceNugetAuthentication,omitempty" xml:"forceNugetAuthentication,omitempty"`
}

func (g GenericRepository) String() string {
	return Stringify(g)
}

// LocalRepository represents a local repository in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON#RepositoryConfigurationJSON-application/vnd.org.jfrog.artifactory.repositories.LocalRepositoryConfiguration+json
type LocalRepository struct {
	*GenericRepository

	DebianTrivialLayout             *bool     `json:"debianTrivialLayout,omitempty" xml:"debianTrivialLayout,omitempty"`
	ChecksumPolicyType              *string   `json:"checksumPolicyType,omitempty" xml:"localRepoChecksumPolicyType,omitempty"`
	MaxUniqueTags                   *int      `json:"maxUniqueTags,omitempty" xml:"maxUniqueTags,omitempty"`
	SnapshotVersionBehavior         *string   `json:"snapshotVersionBehavior,omitempty" xml:"snapshotVersionBehavior,omitempty"`
	ArchiveBrowsingEnabled          *bool     `json:"archiveBrowsingEnabled,omitempty" xml:"archiveBrowsingEnabled,omitempty"`
	CalculateYumMetadata            *bool     `json:"calculateYumMetadata,omitempty" xml:"calculateYumMetadata,omitempty"`
	YumRootDepth                    *int      `json:"yumRootDepth,omitempty" xml:"yumRootDepth,omitempty"`
	DockerAPIVersion                *string   `json:"dockerApiVersion,omitempty" xml:"dockerApiVersion,omitempty"`
	BlockPushingSchema1             *bool     `json:"blockPushingSchema1,omitempty" xml:"blockPushingSchema1,omitempty"`
	EnableFileListsIndexing         *bool     `json:"enableFileListsIndexing,omitempty" xml:"enableFileListsIndexing,omitempty"`
	OptionalIndexCompressionFormats *[]string `json:"optionalIndexCompressionFormats,omitempty" xml:"optionalIndexCompressionFormats>debianFormat,omitempty"`
	XrayIndex                       *bool     `json:"xrayIndex,omitempty" xml:"xray>enabled,omitempty"`
	DownloadRedirect                *bool     `json:"downloadRedirect,omitempty" xml:"downloadRedirect>enabled,omitempty"`
}

func (l LocalRepository) String() string {
	return Stringify(l)
}

// RemoteRepository represents a remote repository in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON#RepositoryConfigurationJSON-application/vnd.org.jfrog.artifactory.repositories.RemoteRepositoryConfiguration+json
type RemoteRepository struct {
	*GenericRepository

	URL                               *string                 `json:"url,omitempty" xml:"url,omitempty"`
	Username                          *string                 `json:"username,omitempty" xml:"username,omitempty"`
	Password                          *string                 `json:"password,omitempty" xml:"password,omitempty"`
	Proxy                             *string                 `json:"proxy,omitempty" xml:"proxyRef,omitempty"`
	RemoteRepoChecksumPolicyType      *string                 `json:"remoteRepoChecksumPolicyType,omitempty" xml:"remoteRepoChecksumPolicyType,omitempty"`
	HardFail                          *bool                   `json:"hardFail,omitempty" xml:"hardFail,omitempty"`
	Offline                           *bool                   `json:"offline,omitempty" xml:"offline,omitempty"`
	StoreArtifactsLocally             *bool                   `json:"storeArtifactsLocally,omitempty" xml:"storeArtifactsLocally,omitempty"`
	SocketTimeoutMillis               *int                    `json:"socketTimeoutMillis,omitempty" xml:"socketTimeoutMillis,omitempty"`
	LocalAddress                      *string                 `json:"localAddress,omitempty" xml:"localAddress,omitempty"`
	RetrievalCachePeriodSecs          *int                    `json:"retrievalCachePeriodSecs,omitempty" xml:"retrievalCachePeriodSecs,omitempty"`
	FailedRetrievalCachePeriodSecs    *int                    `json:"failedRetrievalCachePeriodSecs,omitempty" xml:"failedRetrievalCachePeriodSecs,omitempty"`
	MissedRetrievalCachePeriodSecs    *int                    `json:"missedRetrievalCachePeriodSecs,omitempty" xml:"missedRetrievalCachePeriodSecs,omitempty"`
	MetadataRetrievalTimeoutSecs      *int                    `json:"metadataRetrievalTimeoutSecs,omitempty" xml:"metadataRetrievalTimeoutSecs,omitempty"`
	UnusedArtifactsCleanupEnabled     *bool                   `json:"unusedArtifactsCleanupEnabled,omitempty" xml:"unusedArtifactsCleanupEnabled,omitempty"`
	UnusedArtifactsCleanupPeriodHours *int                    `json:"unusedArtifactsCleanupPeriodHours,omitempty" xml:"unusedArtifactsCleanupPeriodHours,omitempty"`
	AssumedOfflinePeriodSecs          *int                    `json:"assumedOfflinePeriodSecs,omitempty" xml:"assumedOfflinePeriodSecs,omitempty"`
	FetchJarsEagerly                  *bool                   `json:"fetchJarsEagerly,omitempty" xml:"fetchJarsEagerly,omitempty"`
	FetchSourcesEagerly               *bool                   `json:"fetchSourcesEagerly,omitempty" xml:"fetchSourcesEagerly,omitempty"`
	ShareConfiguration                *bool                   `json:"shareConfiguration,omitempty" xml:"shareConfiguration,omitempty"`
	SynchronizeProperties             *bool                   `json:"synchronizeProperties,omitempty" xml:"synchronizeProperties,omitempty"`
	BlockMismatchingMimeTypes         *bool                   `json:"blockMismatchingMimeTypes,omitempty" xml:"blockMismatchingMimeTypes,omitempty"`
	AllowAnyHostAuth                  *bool                   `json:"allowAnyHostAuth,omitempty" xml:"allowAnyHostAuth,omitempty"`
	EnableCookieManagement            *bool                   `json:"enableCookieManagement,omitempty" xml:"enableCookieManagement,omitempty"`
	BowerRegistryURL                  *string                 `json:"bowerRegistryUrl,omitempty" xml:"bowerRegistryUrl,omitempty"`
	ComposerRegistryURL               *string                 `json:"composerRegistryUrl,omitempty" xml:"composerRegistryUrl,omitempty"`
	PyPIRegistryURL                   *string                 `json:"pyPIRegistryUrl,omitempty" xml:"pypi>pyPIRegistryUrl,omitempty"`
	PyPIRepositorySuffix              *string                 `json:"pyPIRepositorySuffix,omitempty" xml:"pyPIRepositorySuffix,omitempty"`
	VcsType                           *string                 `json:"vcsType,omitempty" xml:"vcs>type,omitempty"`
	VcsGitProvider                    *string                 `json:"vcsGitProvider,omitempty" xml:"vcs>git>provider,omitempty"`
	VcsGitDownloadUrl                 *string                 `json:"vcsGitDownloadUrl,omitempty" xml:"vcs>git>downloadUrl,omitempty"`
	BypassHeadRequests                *bool                   `json:"bypassHeadRequests,omitempty" xml:"bypassHeadRequests,omitempty"`
	ClientTLSCertificate              *string                 `json:"clientTlsCertificate,omitempty" xml:"clientTlsCertificate,omitempty"`
	ExternalDependenciesEnabled       *bool                   `json:"externalDependenciesEnabled,omitempty" xml:"externalDependencies>enabled,omitempty"`
	ExternalDependenciesPatterns      *[]string               `json:"externalDependenciesPatterns,omitempty" xml:"externalDependencies>patterns>pattern,omitempty"`
	DownloadRedirect                  *bool                   `json:"downloadRedirect,omitempty" xml:"downloadRedirect>enabled,omitempty"`
	FeedContextPath                   *string                 `json:"feedContextPath,omitempty" xml:"nuget>feedContextPath,omitempty"`
	DownloadContextPath               *string                 `json:"downloadContextPath,omitempty" xml:"nuget>downloadContextPath,omitempty"`
	V3FeedUrl                         *string                 `json:"v3FeedUrl,omitempty" xml:"nuget>v3FeedUrl,omitempty"`
	XrayIndex                         *bool                   `json:"xrayIndex,omitempty" xml:"xray>enabled,omitempty"`
	ListRemoteFolderItems             *bool                   `json:"listRemoteFolderItems,omitempty" xml:"listRemoteFolderItems,omitempty"`
	EnableTokenAuthentication         *bool                   `json:"enableTokenAuthentication,omitempty" xml:"enableTokenAuthentication,omitempty"`
	ContentSynchronisation            *ContentSynchronisation `json:"contentSynchronisation,omitempty" xml:"contentSynchronisation,omitempty"`
	BlockPushingSchema1               *bool                   `json:"blockPushingSchema1,omitempty" xml:"blockPushingSchema1,omitempty"`
	QueryParams                       *string                 `json:"queryParams,omitempty" xml:"queryParams,omitempty"`
	PropagateQueryParams              *bool                   `json:"propagateQueryParams,omitempty" xml:"propagateQueryParams,omitempty"`
}

// ContentSynchronisation represents smart remote repository configuration
type ContentSynchronisation struct {
	Enabled    *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Properties *struct {
		Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	} `json:"properties,omitempty" xml:"properties,omitempty"`
	Statistics *struct {
		Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	} `json:"statistics,omitempty" xml:"statistics,omitempty"`
	Source *struct {
		OriginAbsenceDetection *bool `json:"originAbsenceDetection,omitempty" xml:"originAbsenceDetection,omitempty"`
	} `json:"source,omitempty" xml:"source,omitempty"`
}

func (r RemoteRepository) String() string {
	return Stringify(r)
}

// VirtualRepository represents a virtual repository in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON#RepositoryConfigurationJSON-application/vnd.org.jfrog.artifactory.repositories.VirtualRepositoryConfiguration+json
type VirtualRepository struct {
	*GenericRepository

	Repositories                                  *[]string `json:"repositories,omitempty" xml:"repositories>repositoryRef,omitempty"`
	DebianTrivialLayout                           *bool     `json:"debianTrivialLayout,omitempty" xml:"debianTrivialLayout,omitempty"`
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool     `json:"artifactoryRequestsCanRetrieveRemoteArtifacts,omitempty" xml:"artifactoryRequestsCanRetrieveRemoteArtifacts,omitempty"`
	KeyPair                                       *string   `json:"keyPair,omitempty" xml:"keyPair,omitempty"`
	PomRepositoryReferencesCleanupPolicy          *string   `json:"pomRepositoryReferencesCleanupPolicy,omitempty" xml:"pomRepositoryReferencesCleanupPolicy,omitempty"`
	DefaultDeploymentRepo                         *string   `json:"defaultDeploymentRepo,omitempty" xml:"defaultDeploymentRepo,omitempty"`
	ForceMavenAuthentication                      *bool     `json:"forceMavenAuthentication,omitempty" xml:"forceMavenAuthentication,omitempty"`
	ExternalDependenciesEnabled                   *bool     `json:"externalDependenciesEnabled,omitempty" xml:"externalDependencies>enabled,omitempty"`
	ExternalDependenciesPatterns                  *[]string `json:"externalDependenciesPatterns,omitempty" xml:"externalDependencies>patterns>pattern,omitempty"`
	ExternalDependenciesRemoteRepo                *string   `json:"externalDependenciesRemoteRepo,omitempty" xml:"externalDependencies>remoteRepo,omitempty"`
	ResolveDockerTagsByTimestamp                  *bool     `json:"resolveDockerTagsByTimestamp,omitempty" xml:"resolveDockerTagsByTimestamp,omitempty"`
	VirtualRetrievalCachePeriodSecs               *int      `json:"virtualRetrievalCachePeriodSecs,omitempty" xml:"virtualCacheConfig>virtualRetrievalCachePeriodSecs,omitempty"`
	DebianDefaultArchitectures                    *string   `json:"debianDefaultArchitectures,omitempty" xml:"debianDefaultArchitectures,omitempty"`
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
