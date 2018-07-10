package summary

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling summary
// related Xray API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.POST("/api/v1/summary/artifact", getSummary)
	e.GET("/api/v1/summary/build", getSummary)

	return e
}

func getSummary(c *gin.Context) {
	c.String(200, loadFixture("fixtures/summary/summary.json"))
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
