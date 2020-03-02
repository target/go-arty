package replications

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling repository
// related Artifactory API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/replications", getReplications)
	e.GET("/api/replications/:repository", getReplication)
	e.DELETE("/api/replications/:repository", deleteReplication)
	e.PUT("/api/replications/:repository", createReplication)
	e.POST("/api/replications/:repository", updateReplication)
	e.PUT("/api/replications/:repository/:multiRepository", createReplication)
	e.POST("/api/replications/:repository/:multiRepository", updateReplication)

	return e
}

func getReplications(c *gin.Context) {
	c.String(200, loadFixture("fixtures/replications/replications.json"))
}

func getReplication(c *gin.Context) {
	repository := c.Param("repository")

	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	if strings.Contains(repository, "local") {
		c.String(200, loadFixture("fixtures/replications/local_replication.json"))
		return
	}

	if strings.Contains(repository, "remote") {
		c.String(200, loadFixture("fixtures/replications/remote_replication.json"))
		return
	}

	c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
}

func createReplication(c *gin.Context) {
	repository := c.Param("multiRepository")
	if repository == "" {
		repository = c.Param("repository")
	}
	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.JSON(201, fmt.Sprintf("Successfully created replication for repository '%s'", repository))
}

func updateReplication(c *gin.Context) {
	repository := c.Param("multiRepository")
	if repository == "" {
		repository = c.Param("repository")
	}
	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.JSON(200, fmt.Sprintf("Successfully updated replication for repository '%s'", repository))
}

func deleteReplication(c *gin.Context) {
	repository := c.Param("repository")
	url := c.Query("url")

	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	if len(url) > 0 {
		if strings.Contains(url, "http://host:port/target-repo") {
			c.JSON(200, fmt.Sprintf("Successfully deleted replication url '%s' for repository '%s'", url, repository))
			return
		}

		c.JSON(400, fmt.Sprintln("Invalid replication url"))
	}

	c.JSON(200, fmt.Sprintf("Successfully deleted replications for repository '%s'", repository))
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
