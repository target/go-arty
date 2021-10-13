package artifactory

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"github.com/target/go-arty/v2/artifactory/fixtures/replications"
)

func Test_Replications(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(replications.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("Replications Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("Replication", func() {
			replication := &Replication{
				Username:                        String("replication-user"),
				Password:                        String("pass"),
				Url:                             String("http://host:port/some-repo"),
				SocketTimeoutMillis:             Int(15000),
				CronExp:                         String("0 0 12 * * ?"),
				RepoKey:                         String("test"),
				EnableEventReplication:          Bool(true),
				Enabled:                         Bool(true),
				SyncDeletes:                     Bool(true),
				SyncProperties:                  Bool(true),
				SyncStatistics:                  Bool(true),
				PathPrefix:                      String("path/to/replicate"),
				CheckBinaryExistenceInFilestore: Bool(false),
			}

			replications := &Replications{
				ReplicationType: String("PUSH"),
				Replication: &Replication{
					Enabled:                         Bool(true),
					CronExp:                         String("0 0 12 * * ?"),
					SyncDeletes:                     Bool(true),
					SyncProperties:                  Bool(true),
					PathPrefix:                      String("path/to/replicate"),
					RepoKey:                         String("local-repo1"),
					EnableEventReplication:          Bool(true),
					CheckBinaryExistenceInFilestore: Bool(false),
					Url:                             String("http://host:port/target-repo"),
					SyncStatistics:                  Bool(false),
				},
			}

			localReplication := &Replication{
				Username:                        String("replication-user"),
				Password:                        String("pass"),
				Url:                             String("http://host:port/target-repo"),
				SocketTimeoutMillis:             Int(15000),
				CronExp:                         String("0 0 12 * * ?"),
				RepoKey:                         String("local-repo1"),
				EnableEventReplication:          Bool(true),
				Enabled:                         Bool(true),
				SyncDeletes:                     Bool(true),
				SyncProperties:                  Bool(true),
				SyncStatistics:                  Bool(false),
				PathPrefix:                      String("path/to/replicate"),
				CheckBinaryExistenceInFilestore: Bool(false),
			}

			remoteReplication := &Replication{
				Enabled:                         Bool(true),
				CronExp:                         String("0 0 12 * * ?"),
				SyncDeletes:                     Bool(true),
				SyncProperties:                  Bool(true),
				PathPrefix:                      String(""),
				RepoKey:                         String("remote-repo1"),
				EnableEventReplication:          Bool(true),
				CheckBinaryExistenceInFilestore: Bool(false),
			}

			multiPushReplication := &MultiPushReplication{
				CronExp:                String("0 0 12 * * ?"),
				EnableEventReplication: Bool(true),
				Replications: &[]Replication{
					*localReplication,
					*replication,
				},
			}

			g.It("- should return valid string for replications with String()", func() {
				actual := []Replications{
					*replications,
				}

				data, _ := ioutil.ReadFile("fixtures/replications/replications.json")

				var expected []Replications
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual[0].String() == expected[0].String()).IsTrue()
			})

			g.It("- should return valid string for Local (push) replication with String()", func() {
				actual := []Replication{
					*localReplication,
				}

				data, _ := ioutil.ReadFile("fixtures/replications/local_replication.json")

				var expected []Replication
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual[0].String() == expected[0].String()).IsTrue()
			})

			g.It("- should return valid string for Remote (pull) replication with String()", func() {
				actual := remoteReplication

				data, _ := ioutil.ReadFile("fixtures/replications/remote_replication.json")

				var expected Replication
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with GetAll()", func() {
				actual, resp, err := c.Replications.GetAll()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Get() on local replication", func() {
				actual, resp, err := c.Replications.Get("local-repo1")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Get() on remote replication", func() {
				actual, resp, err := c.Replications.Get("remote-repo1")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				if err != nil {
					g.Fail(err.Error())
				}
			})

			g.It("- should return no error with Create()", func() {
				actual, resp, err := c.Replications.Create("test", replication)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				if err != nil {
					g.Fail(err.Error())
				}
			})

			g.It("- should return no error with Update()", func() {
				actual, resp, err := c.Replications.Update("test", replication)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				if err != nil {
					g.Fail(err.Error())
				}
			})

			g.It("- should return no error with Delete()", func() {
				actual, resp, err := c.Replications.Delete("test")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				if err != nil {
					g.Fail(err.Error())
				}
			})

			g.It("- should return no error with CreateMultiPush()", func() {
				actual, resp, err := c.Replications.CreateMultiPush("test", multiPushReplication)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				if err != nil {
					g.Fail(err.Error())
				}
			})

			g.It("- should return no error with UpdateMultiPush()", func() {
				actual, resp, err := c.Replications.UpdateMultiPush("test", multiPushReplication)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				if err != nil {
					g.Fail(err.Error())
				}
			})

			g.It("- should return no error with DeleteMultiPush()", func() {
				actual, resp, err := c.Replications.DeleteMultiPush("test", "http://host:port/target-repo")
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				if err != nil {
					g.Fail(err.Error())
				}
			})
		})
	})
}
