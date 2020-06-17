package system

import (
	"fmt"
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
	e.GET("/api/system/configuration", getSystemConfig)
	e.GET("/api/system/version", getSystemVersion)
	e.PATCH("/api/system/configuration", updateSystemConfig)

	return e
}

func ping(c *gin.Context) {
	c.String(200, "OK")
}

func getSystemConfig(c *gin.Context) {
	c.String(200, loadFixture("fixtures/system/artifactory.config.xml"))
}

func getSystemInfo(c *gin.Context) {
	c.String(200, loadFixture("fixtures/system/system.txt"))
}

func getSystemVersion(c *gin.Context) {
	c.String(200, loadFixture("fixtures/system/version.json"))
}

func updateSystemConfig(c *gin.Context) {
	req, _ := ioutil.ReadAll(c.Request.Body)

	if loadFixture("fixtures/system/patch.yaml") == string(req) {
		c.String(200, fmt.Sprint("11 changes to config merged successfully"))
	} else {
		c.JSON(400, loadFixture("fixtures/system/invalid.json"))
	}
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
