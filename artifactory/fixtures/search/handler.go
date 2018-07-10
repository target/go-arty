package search

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling search
// related Artifactory API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/search/gavc", getSearch)

	return e
}

func getSearch(c *gin.Context) {
	c.String(200, loadFixture("fixtures/search/files.json"))
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
