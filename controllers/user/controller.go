package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func Create(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
