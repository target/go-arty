package artifacts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling artifact
// related Artifactory API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/:repository/foo.txt", downloadFile)
	e.PUT("/:repository/folder/*path", uploadFile)
	e.POST("/api/copy/:repository/*path", copyFile)
	e.POST("/api/move/:repository/*path", moveFile)
	e.DELETE("/:repository/foo.txt", deleteFile)

	return e
}

func downloadFile(c *gin.Context) {
	repository := c.Param("repository")

	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.Data(200, "application/text", loadFixture("fixtures/artifacts/foo.txt"))
}

func uploadFile(c *gin.Context) {
	repository := c.Param("repository")

	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.String(200, "")
}

func copyFile(c *gin.Context) {
	repository := c.Param("repository")

	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.String(200, string(loadFixture("fixtures/artifacts/copy.json")))
}

func moveFile(c *gin.Context) {
	repository := c.Param("repository")

	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.String(200, string(loadFixture("fixtures/artifacts/move.json")))
}

func deleteFile(c *gin.Context) {
	repository := c.Param("repository")

	if strings.Contains(repository, "not-found") {
		c.JSON(404, fmt.Sprintf("Repository %s does not exist", repository))
		return
	}

	c.String(204, "")
}

func loadFixture(file string) []byte {
	data, _ := ioutil.ReadFile(file)

	return data
}
