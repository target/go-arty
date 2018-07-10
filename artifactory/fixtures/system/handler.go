package system

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling system
// related Artifactory API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/system/ping", ping)
	e.GET("/api/system", getSystemInfo)
	e.GET("/api/system/version", getSystemVersion)

	return e
}

func ping(c *gin.Context) {
	c.String(200, "OK")
}

func getSystemInfo(c *gin.Context) {
	c.String(200, loadFixture("fixtures/system/system.txt"))
}

func getSystemVersion(c *gin.Context) {
	c.String(200, loadFixture("fixtures/system/version.json"))
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
