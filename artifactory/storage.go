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
	"strings"
)

// StorageService handles communication with the storage related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ARTIFACTS&STORAGE
type StorageService service

// Child represents a child under a folder in Artifactory.
type Child struct {
	URI    *string `json:"uri,omitempty"`
	Folder *string `json:"folder,omitempty"`
}

// Folder represents a folder in Artifactory.
type Folder struct {
	URI          *string  `json:"uri,omitempty"`
	Repo         *string  `json:"repo,omitempty"`
	Path         *string  `json:"path,omitempty"`
	Created      *string  `json:"created,omitempty"`
	CreatedBy    *string  `json:"createdBy,omitempty"`
	LastModified *string  `json:"lastModified,omitempty"`
	ModifiedBy   *string  `json:"modifiedBy,omitempty"`
	LastUpdated  *string  `json:"lastUpdated,omitempty"`
	Children     *[]Child `json:"children,omitempty"`
}

// Checksums represents the checksums for a file in Artifactory.
type Checksums struct {
	MD5    *string `json:"md5,omitempty"`
	SHA1   *string `json:"sha1,omitempty"`
	SHA256 *string `json:"sha256,omitempty"`
}

// File represents a file in Artifactory.
type File struct {
	URI               *string    `json:"uri,omitempty"`
	DownloadURI       *string    `json:"downloadUri,omitempty"`
	Repo              *string    `json:"repo,omitempty"`
	Path              *string    `json:"path,omitempty"`
	RemoteURL         *string    `json:"remoteUrl,omitempty"`
	Created           *string    `json:"created,omitempty"`
	CreatedBy         *string    `json:"createdBy,omitempty"`
	LastModified      *string    `json:"lastModified,omitempty"`
	ModifiedBy        *string    `json:"modifiedBy,omitempty"`
	LastUpdated       *string    `json:"lastUpdated,omitempty"`
	Size              *string    `json:"size,omitempty"`
	MimeType          *string    `json:"mimeType,omitempty"`
	Checksums         *Checksums `json:"checksums,omitempty"`
	OriginalChecksums *Checksums `json:"originalChecksums,omitempty"`
}

// BinariesSummary represents the summary of binaries in Artifactory.
type BinariesSummary struct {
	BinariesCount  *string `json:"binariesCount,omitempty"`
	BinariesSize   *string `json:"binariesSize,omitempty"`
	ArtifactsSize  *string `json:"artifactsSize,omitempty"`
	Optimization   *string `json:"optimization,omitempty"`
	ItemsCount     *string `json:"itemsCount,omitempty"`
	ArtifactsCount *string `json:"artifactsCount,omitempty"`
}

// FileStoreSummary represents the summary of file storage in Artifactory.
type FileStoreSummary struct {
	StorageType      *string `json:"storageType,omitempty"`
	StorageDirectory *string `json:"storageDirectory,omitempty"`
	TotalSpace       *string `json:"totalSpace,omitempty"`
	UsedSpace        *string `json:"usedSpace,omitempty"`
	FreeSpace        *string `json:"freeSpace,omitempty"`
}

// RepositoriesSummary represents the summary of repositories in Artifactory.
type RepositoriesSummary struct {
	RepoKey      *string `json:"repoKey,omitempty"`
	RepoType     *string `json:"repoType,omitempty"`
	FoldersCount *int    `json:"foldersCount,omitempty"`
	FilesCount   *int    `json:"filesCount,omitempty"`
	UsedSpace    *string `json:"usedSpace,omitempty"`
	ItemsCount   *int    `json:"itemsCount,omitempty"`
	PackageType  *string `json:"packageType,omitempty"`
	Percentage   *string `json:"percentage,omitempty"`
}

// StorageSummary represents the summary of storage in Artifactory.
type StorageSummary struct {
	BinariesSummary         *BinariesSummary       `json:"binariesSummary,omitempty"`
	FileStoreSummary        *FileStoreSummary      `json:"fileStoreSummary,omitempty"`
	RepositoriesSummaryList *[]RepositoriesSummary `json:"repositoriesSummaryList,omitempty"`
}

// ItemLastModified represents the last modified date for a file in Artifactory.
type ItemLastModified struct {
	URI          *string `json:"uri,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
}

// FileStatistics represents statistics for a file in Artifactory.
type FileStatistics struct {
	URI              *string `json:"uri,omitempty"`
	LastDownloaded   *string `json:"lastDownloaded,omitempty"`
	DownloadCount    *int    `json:"downloadCount,omitempty"`
	LastDownloadedBy *string `json:"lastDownloadedBy,omitempty"`
}

// ItemProperties represents a set of properties for an item in Artifactory.
type ItemProperties struct {
	URI        *string              `json:"uri,omitempty"`
	Properties *map[string][]string `json:"properties,omitempty"`
}

// FileList represents a list of files in Artifactory.
type FileList struct {
	URI     *string         `json:"uri,omitempty"`
	Created *string         `json:"created,omitempty"`
	Files   *[]FileListItem `json:"files,omitempty"`
}

// FileListItem represents an item in a list of files in Artifactory.
type FileListItem struct {
	URI          *string `json:"uri,omitempty"`
	Size         *int    `json:"size,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	Folder       *bool   `json:"folder,omitempty"`
	SHA1         *string `json:"sha1,omitempty"`
}

// GetFolder returns the provided folder.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-FolderInfo
func (s *StorageService) GetFolder(repo, path string) (*Folder, *Response, error) {
	u := fmt.Sprintf("/api/storage/%s/%s", repo, path)
	v := new(Folder)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetFile returns the provided file.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-FileInfo
func (s *StorageService) GetFile(repo, path string) (*File, *Response, error) {
	u := fmt.Sprintf("/api/storage/%s/%s", repo, path)
	v := new(File)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetItemLastModified returns the ISO8601 timestamp of the provided item's last modified date.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ItemLastModified
func (s *StorageService) GetItemLastModified(repo, path string) (*ItemLastModified, *Response, error) {
	u := fmt.Sprintf("/api/storage/%s/%s?lastModified", repo, path)
	v := new(ItemLastModified)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetFileStatistics returns download statistics for the provided file.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-FileStatistics
func (s *StorageService) GetFileStatistics(repo, path string) (*FileStatistics, *Response, error) {
	u := fmt.Sprintf("/api/storage/%s/%s?stats", repo, path)
	v := new(FileStatistics)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetItemProperties returns properties on the provided item.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ItemProperties
func (s *StorageService) GetItemProperties(repo, path string) (*ItemProperties, *Response, error) {
	u := fmt.Sprintf("/api/storage/%s/%s?properties", repo, path)
	v := new(ItemProperties)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// SetItemProperties attaches the provided properties to the provided item.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SetItemProperties
func (s *StorageService) SetItemProperties(repo, path string, properties map[string][]string) (*Response, error) {
	var propertyString string
	var index int
	for k, v := range properties {
		index++
		if len(v) == 1 {
			propertyString = propertyString + fmt.Sprintf("%s=%s", k, v[0])
		} else {
			propertyString = propertyString + fmt.Sprintf("%s=[%s]", k, strings.Join(v, ","))
		}

		if index != len(properties) {
			propertyString = propertyString + ";"
		}
	}

	u := fmt.Sprintf("/api/storage/%s/%s?properties=%s&recursive=1", repo, path, propertyString)

	resp, err := s.client.Call("PUT", u, nil, nil)
	return resp, err
}

// DeleteItemProperties removes the provided properties from the provided item.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-DeleteItemProperties
func (s *StorageService) DeleteItemProperties(repo, path string, properties []string) (*Response, error) {
	var propertyString string
	for k, v := range properties {
		propertyString = propertyString + fmt.Sprintf("%s,", v)

		if k != len(properties) {
			propertyString = strings.TrimRight(propertyString, ",")
		}
	}

	u := fmt.Sprintf("/api/storage/%s/%s?properties=%s", repo, path, propertyString)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}

// GetFileList lists all files in the provided repo.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-FileList
func (s *StorageService) GetFileList(repo, path string) (*FileList, *Response, error) {
	u := fmt.Sprintf("/api/storage/%s/%s?list&deep=1", repo, path)
	v := new(FileList)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetStorageSummary returns the storage summary information.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetStorageSummaryInfo
func (s *StorageService) GetStorageSummary() (*StorageSummary, *Response, error) {
	u := "/api/storageinfo"
	v := new(StorageSummary)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}
