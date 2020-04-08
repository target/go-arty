package permissions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling permission
// related Artifactory API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	e.GET("/api/security/permissions", getPermissions)
	e.GET("/api/security/permissions/:target", getPermission)
	e.GET("/api/v2/security/permissions/:target", getV2Permission)
	e.PUT("/api/security/permissions/:target", createPermission)
	e.PUT("/api/v2/security/permissions/:target", updatePermission)
	e.DELETE("/api/security/permissions/:target", deletePermission)
	e.HEAD("/api/v2/security/permissions/:target", getExistence)

	return e
}

func getPermissions(c *gin.Context) {
	c.String(200, loadFixture("fixtures/permissions/permissions.json"))
}

func getPermission(c *gin.Context) {
	target := c.Param("target")

	if strings.Contains(target, "not-found") {
		c.JSON(404, fmt.Sprintf("Permission Target %s does not exist", target))
		return
	}

	c.String(200, loadFixture("fixtures/permissions/permission.json"))
}

func createPermission(c *gin.Context) {
	target := c.Param("target")

	if strings.Contains(target, "not-found") {
		c.JSON(404, fmt.Sprintf("Permission Target %s does not exist", target))
		return
	}

	c.JSON(201, "")
}

func deletePermission(c *gin.Context) {
	target := c.Param("target")

	if strings.Contains(target, "not-found") {
		c.JSON(404, fmt.Sprintf("Permission Target %s does not exist", target))
		return
	}

	c.JSON(200, fmt.Sprintf("Successfully deleted permission Target '%s'", target))
}

func getExistence(c *gin.Context) {
	target := c.Param("target")

	switch target {
	case "valid":
		c.Status(200)
		return
	default:
		c.Status(404)
		return
	}
}

func updatePermission(c *gin.Context) {
	target := c.Param("target")

	switch target {
	case "valid":
		c.Status(201)
		return
	case "invalid":
		c.JSON(400, loadFixture("fixtures/permissions/invalid.json"))
		return
	}
}

func getV2Permission(c *gin.Context) {
	target := c.Param("target")

	switch target {
	case "java-developers":
		c.JSON(200, loadFixture("fixtures/permissions/permissionv2.json"))
		return
	default:
		c.JSON(404, loadFixture("fixtures/permissions/non-existent.json"))
		return
	}
}

func loadFixture(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}
