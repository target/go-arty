package repositories

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

	e.GET("/api/repositories", getRepositories)
	e.GET("/api/repositories/:repository", getRepository)
	e.PUT("/api/repositories/:repository", createRepository)
	e.POST("/api/repositories/:repository", createRepository)
	e.DELETE("/api/repositories/:repository", deleteRepository)

	return e
}

func getRepositories(c *gin.Context) {
	c.String(200, loadFixture("fixtures/repositories/repositories.json"))
}

func getRepository(c *gin.Context) {
	repository := c.Param("repository")

	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	if strings.Contains(repository, "local") {
		c.String(200, loadFixture("fixtures/repositories/local_repository.json"))
		return
	}

	if strings.Contains(repository, "remote") {
		c.String(200, loadFixture("fixtures/repositories/remote_repository.json"))
		return
	}

	if strings.Contains(repository, "virtual") {
		c.String(200, loadFixture("fixtures/repositories/virtual_repository.json"))
		return
	}

	c.String(200, loadFixture("fixtures/repositories/generic_repository.json"))
}

func createRepository(c *gin.Context) {
	repository := c.Param("repository")

	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.JSON(201, fmt.Sprintf("Successfully created repository '%s'", repository))
}

func deleteRepository(c *gin.Context) {
	repository := c.Param("repository")

	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.JSON(200, fmt.Sprintf("Repository %s and all its content have been removed successfully.", repository))
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
