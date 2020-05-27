package artifactory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ReplicationsService handles communication with the replications related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetRepositoryReplicationConfiguration
type ReplicationsService service

// Replication represents possible fields across all replication types in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Replication+Configuration+JSON
type Replication struct {
	Username                        *string `json:"username,omitempty" xml:"username,omitempty"`
	Password                        *string `json:"password,omitempty" xml:"password,omitempty"`
	Url                             *string `json:"url,omitempty" xml:"url,omitempty"`
	SocketTimeoutMillis             *int    `json:"socketTimeoutMillis,omitempty" xml:"socketTimeoutMillis,omitempty"`
	CronExp                         *string `json:"cronExp,omitempty" xml:"cronExp,omitempty"`
	RepoKey                         *string `json:"repoKey,omitempty" xml:"repoKey,omitempty"`
	EnableEventReplication          *bool   `json:"enableEventReplication,omitempty" xml:"enableEventReplication,omitempty"`
	Enabled                         *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	SyncDeletes                     *bool   `json:"syncDeletes,omitempty" xml:"syncDeletes,omitempty"`
	SyncProperties                  *bool   `json:"syncProperties,omitempty" xml:"syncProperties,omitempty"`
	SyncStatistics                  *bool   `json:"syncStatistics,omitempty" xml:"syncStatistics,omitempty"`
	PathPrefix                      *string `json:"pathPrefix,omitempty" xml:"pathPrefix,omitempty"`
	CheckBinaryExistenceInFilestore *bool   `json:"checkBinaryExistenceInFilestore,omitempty" xml:"checkBinaryExistenceInFilestore,omitempty"`
}

func (r Replication) String() string {
	return Stringify(r)
}

// Replications represents a replication returned by the undocumented replications endpoint.
//
// Docs: This struct is currently undocumented by JFrog
type Replications struct {
	*Replication
	ReplicationType *string `json:"replicationType,omitempty"`
}

func (r Replications) String() string {
	return Stringify(r)
}

// MultiPushReplication represents a Local Multi-push replication in Artifactory
type MultiPushReplication struct {
	CronExp                *string        `json:"cronExp,omitempty"`
	EnableEventReplication *bool          `json:"enableEventReplication,omitempty"`
	Replications           *[]Replication `json:"replications,omitempty"`
}

func (r MultiPushReplication) String() string {
	return Stringify(r)
}

// GetAll returns a list of all replications.
//
// Docs: This endpoint is currently undocumented by JFrog
func (r *ReplicationsService) GetAll() (*[]Replications, *Response, error) {
	u := fmt.Sprintf("/api/replications")
	v := new([]Replications)

	resp, err := r.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns replications for the provided repository.
// Artifactory returns a JSON array for a local replication or a JSON object for a remote replication.
// This method returns a slice to maintain consistency.
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetRepositoryReplicationConfiguration
func (r *ReplicationsService) Get(repo string) (*[]Replication, *Response, error) {
	u := fmt.Sprintf("/api/replications/%s", repo)
	v := new(bytes.Buffer)

	replications := new([]Replication)

	resp, err := r.client.Call("GET", u, nil, v)
	if err != nil {
		if resp.StatusCode == 404 {
			return replications, resp, err
		}
		return nil, resp, err
	}

	body, err := ioutil.ReadAll(v)
	if err != nil {
		return nil, resp, err
	}

	switch body[0] {
	case '[':
		err = json.Unmarshal(body, replications)
	case '{':
		replication := new(Replication)
		if err = json.Unmarshal(body, replication); err == nil {
			replications = &[]Replication{*replication}
		}
	}
	return replications, resp, err
}

// Create constructs a single replication for the provided repository.
// If multiple push replications are required CreateMultiPush needs to be used
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-CreateRepository
func (r *ReplicationsService) Create(repo string, replication *Replication) (*string, *Response, error) {
	u := fmt.Sprintf("/api/replications/%s", repo)
	v := new(string)

	resp, err := r.client.Call("PUT", u, replication, v)
	return v, resp, err
}

// Update updates a single replication for the provided repository. If updates are required
// for a local repository with multiple push replications UpdateMultiPush needs to be used
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateRepositoryReplicationConfiguration
func (r *ReplicationsService) Update(repo string, replication *Replication) (*string, *Response, error) {
	u := fmt.Sprintf("/api/replications/%s", repo)
	v := new(string)

	resp, err := r.client.Call("POST", u, replication, v)
	return v, resp, err
}

// Delete deletes the existing replication configuration for the provided repository.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-DeleteRepositoryReplicationConfiguration
func (r *ReplicationsService) Delete(repo string) (*string, *Response, error) {
	u := fmt.Sprintf("/api/replications/%s", repo)
	v := new(string)

	resp, err := r.client.Call("DELETE", u, nil, v)
	return v, resp, err
}

// CreateMultiPush constructs a Local Multi-push replication for the provided repository.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceLocalMulti-pushReplication
func (r *ReplicationsService) CreateMultiPush(repo string, replications *MultiPushReplication) (*string, *Response, error) {
	u := fmt.Sprintf("/api/replications/multiple/%s", repo)
	v := new(string)

	resp, err := r.client.Call("PUT", u, replications, v)
	return v, resp, err
}

// UpdateMultiPush updates a Local Multi-push replication for the provided repository
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateLocalMulti-pushReplication
func (r *ReplicationsService) UpdateMultiPush(repo string, replications *MultiPushReplication) (*string, *Response, error) {
	u := fmt.Sprintf("/api/replications/multiple/%s", repo)
	v := new(string)

	resp, err := r.client.Call("POST", u, replications, v)
	return v, resp, err
}

// DeleteMultiPush deletes replication configuration at the provided URL for the provided repository.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-DeleteRepositoryReplicationConfiguration
func (r *ReplicationsService) DeleteMultiPush(repo string, url string) (*string, *Response, error) {
	u := fmt.Sprintf("/api/replications/%s?url=%s", repo, url)
	v := new(string)

	resp, err := r.client.Call("DELETE", u, nil, v)
	return v, resp, err
}
