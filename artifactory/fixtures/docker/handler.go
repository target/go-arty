package docker

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling docker
// related Artifactory API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/docker/:registry/v2/_catalog", getRepositories)
	e.GET("/api/docker/:registry/v2/docker-dev/tags/list", getTags)
	e.POST("/api/docker/:registry/v2/promote", promoteImage)

	return e
}

func getRepositories(c *gin.Context) {
	registry := c.Param("registry")

	if strings.Contains(registry, "not-found") {
		c.JSON(404, fmt.Sprintf("Registry %s does not exist", registry))
		return
	}

	c.String(200, loadFixture("fixtures/docker/repositories.json"))
}

func getTags(c *gin.Context) {
	registry := c.Param("registry")

	if strings.Contains(registry, "not-found") {
		c.JSON(404, fmt.Sprintf("Registry %s does not exist", registry))
		return
	}

	c.String(200, loadFixture("fixtures/docker/tags.json"))
}

func promoteImage(c *gin.Context) {
	registry := c.Param("registry")

	if strings.Contains(registry, "not-found") {
		c.JSON(404, fmt.Sprintf("Registry %s does not exist", registry))
		return
	}

	c.JSON(200, "Promotion ended successfully")
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
