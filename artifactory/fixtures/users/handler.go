package users

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling user
// related Artifactory API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/users", getUsers)
	e.GET("/api/security/users", getSecurityUsers)
	e.GET("/api/security/users/:user", getSecurityUser)
	e.PUT("/api/security/users/:user", createSecurityUser)
	e.POST("/api/security/users/:user", updateSecurityUser)
	e.DELETE("/api/security/users/:user", deleteSecurityUser)

	e.GET("/api/security/apiKey", getAPIKey)
	e.POST("/api/security/apiKey", getAPIKey)
	e.PUT("/api/security/apiKey", getAPIKey)
	e.DELETE("/api/security/apiKey", deleteAPIKey)
	e.DELETE("/api/security/apiKey/:user", deleteSecurityUserAPIKey)

	e.GET("/api/security/encryptedPassword", getEncryptedPassword)

	return e
}

func getUsers(c *gin.Context) {
	c.String(200, loadFixture("fixtures/users/users.json"))
}

func getSecurityUsers(c *gin.Context) {
	c.String(200, loadFixture("fixtures/users/security_users.json"))
}

func getSecurityUser(c *gin.Context) {
	user := c.Param("user")

	if strings.Contains(user, "not-found") {
		c.JSON(404, fmt.Sprintf("User %s does not exist", user))
		return
	}

	c.String(200, loadFixture("fixtures/users/security_user.json"))
}

func createSecurityUser(c *gin.Context) {
	user := c.Param("user")

	if strings.Contains(user, "not-found") {
		c.JSON(404, fmt.Sprintf("User %s does not exist", user))
		return
	}

	c.JSON(201, "")
}

func updateSecurityUser(c *gin.Context) {
	user := c.Param("user")

	if strings.Contains(user, "not-found") {
		c.JSON(404, fmt.Sprintf("User %s does not exist", user))
		return
	}

	c.JSON(200, "")
}

func deleteSecurityUser(c *gin.Context) {
	user := c.Param("user")

	if strings.Contains(user, "not-found") {
		c.JSON(404, fmt.Sprintf("User %s does not exist", user))
		return
	}

	c.JSON(201, fmt.Sprintf("User %s has been removed successfully.", user))
}

func getAPIKey(c *gin.Context) {
	c.String(200, loadFixture("fixtures/users/get_api_key.json"))
}

func deleteAPIKey(c *gin.Context) {
	c.String(200, loadFixture("fixtures/users/delete_api_key.json"))
}

func deleteSecurityUserAPIKey(c *gin.Context) {
	user := c.Param("user")

	if strings.Contains(user, "not-found") {
		c.JSON(404, fmt.Sprintf("User %s does not exist", user))
		return
	}

	c.String(200, loadFixture("fixtures/users/delete_api_key.json"))
}

func getEncryptedPassword(c *gin.Context) {
	c.JSON(200, "AreallyCOOLsuperSECRETencryptedPasswordthatNOBODYknows")
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
