package users

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling user
// related Xray API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/v1/users", getUsers)
	e.GET("/api/v1/users/:user", getUser)
	e.POST("/api/v1/users", createUser)
	e.PUT("/api/v1/users/:user", updateUser)
	e.DELETE("/api/v1/users/:user", deleteUser)

	return e
}

func getUsers(c *gin.Context) {
	c.String(200, loadFixture("fixtures/users/users.json"))
}

func getUser(c *gin.Context) {
	user := c.Param("user")

	if strings.Contains(user, "not-found") {
		c.JSON(404, fmt.Sprintf("User %s does not exist", user))
		return
	}

	c.String(200, loadFixture("fixtures/users/user.json"))
}

func createUser(c *gin.Context) {
	c.String(201, loadFixture("fixtures/users/user.json"))
}

func updateUser(c *gin.Context) {
	user := c.Param("user")

	if strings.Contains(user, "not-found") {
		c.JSON(404, fmt.Sprintf("User %s does not exist", user))
		return
	}

	c.String(200, "")
}

func deleteUser(c *gin.Context) {
	user := c.Param("user")

	if strings.Contains(user, "not-found") {
		c.JSON(404, fmt.Sprintf("User %s does not exist", user))
		return
	}

	c.String(200, "")
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
