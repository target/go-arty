package storage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling storage
// related Artifactory API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/storage/:repository/folder", getFolder)
	e.GET("/api/storage/:repository/file", getFile)
	e.GET("/api/storageinfo", getStorageSummary)

	e.PUT("/api/storage/:repository/file", itemProperties)
	e.DELETE("/api/storage/:repository/file", itemProperties)

	return e
}

func getFolder(c *gin.Context) {
	repository := c.Param("repository")
	perms := c.Query("permissions")
	if len(perms) > 0 {
		if strings.Contains(repository, "not-found") {
			c.JSON(400, "This method can only be invoked on local/cached repositories.")
			return
		}

		c.String(200, loadFixture("fixtures/storage/effective_permissions.json"))
	}
	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.String(200, loadFixture("fixtures/storage/folder.json"))
}

func getFile(c *gin.Context) {
	lastMod := c.Query("lastModified")
	if len(lastMod) > 0 {
		c.String(200, loadFixture("fixtures/storage/last_modified.json"))
		return
	}

	stats := c.Query("stats")
	if len(stats) > 0 {
		c.String(200, loadFixture("fixtures/storage/statistics.json"))
		return
	}

	properties := c.Query("properties")
	if len(properties) > 0 {
		c.String(200, loadFixture("fixtures/storage/properties.json"))
		return
	}

	list := c.Query("list")
	if len(list) > 0 {
		c.String(200, loadFixture("fixtures/storage/file_list.json"))
		return
	}

	repository := c.Param("repository")
	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.String(200, loadFixture("fixtures/storage/file.json"))
}

func itemProperties(c *gin.Context) {
	repository := c.Param("repository")
	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.JSON(204, "")
}

func getStorageSummary(c *gin.Context) {
	c.String(200, loadFixture("fixtures/storage/storage_summary.json"))
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
