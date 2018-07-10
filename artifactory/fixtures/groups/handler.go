package groups

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling group
// related Artifactory API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/security/groups", getGroups)
	e.GET("/api/security/groups/:group", getGroup)
	e.PUT("/api/security/groups/:group", createGroup)
	e.POST("/api/security/groups/:group", createGroup)
	e.DELETE("/api/security/groups/:group", deleteGroup)

	return e
}

func getGroups(c *gin.Context) {
	c.String(200, loadFixture("fixtures/groups/groups.json"))
}

func getGroup(c *gin.Context) {
	group := c.Param("group")

	if strings.Contains(group, "not-found") {
		c.JSON(404, fmt.Sprintf("Group %s does not exist", group))
		return
	}

	c.String(200, loadFixture("fixtures/groups/group.json"))
}

func createGroup(c *gin.Context) {
	group := c.Param("group")

	if strings.Contains(group, "not-found") {
		c.JSON(404, fmt.Sprintf("Group %s does not exist", group))
		return
	}

	c.JSON(201, "")
}

func deleteGroup(c *gin.Context) {
	group := c.Param("group")

	if strings.Contains(group, "not-found") {
		c.JSON(404, fmt.Sprintf("Group %s does not exist", group))
		return
	}

	c.JSON(200, fmt.Sprintf("Group '%s' has been removed successfully.", group))
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
