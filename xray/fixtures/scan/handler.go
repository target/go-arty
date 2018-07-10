package scan

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling scan
// related Xray API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.POST("/api/v1/scanArtifact", scanArtifact)
	e.POST("/api/v1/scanBuild", scanBuild)

	return e
}

func scanArtifact(c *gin.Context) {
	c.String(200, loadFixture("fixtures/scan/artifact.json"))
}

func scanBuild(c *gin.Context) {
	c.String(200, loadFixture("fixtures/system/version.json"))
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
