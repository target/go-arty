package licenses

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling license
// related Artifactory API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/system/licenses", getLicense)
	e.POST("/api/system/licenses", installLicense)

	return e
}

func getLicense(c *gin.Context) {
	c.String(200, loadFixture("fixtures/licenses/get_license.json"))
}

func installLicense(c *gin.Context) {
	c.String(200, loadFixture("fixtures/licenses/post_license.json"))
}

// FakeHAHandler returns an http.Handler that is capable of handling HA license
// related Artifactory API requests and returning mock responses.
func FakeHAHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/system/licenses", getHALicenses)
	e.POST("/api/system/licenses", installHALicenses)
	e.DELETE("/api/system/licenses", deleteHALicenses)

	return e
}

func getHALicenses(c *gin.Context) {
	c.String(200, loadFixture("fixtures/licenses/get_ha_license.json"))
}

func installHALicenses(c *gin.Context) {
	c.String(200, loadFixture("fixtures/licenses/post_ha_license.json"))
}

func deleteHALicenses(c *gin.Context) {
	c.String(200, loadFixture("fixtures/licenses/delete_ha_license.json"))
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
