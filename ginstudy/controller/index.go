package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexGetAction(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "hello gin get method")
}

func IndexPostAction(c *gin.Context) {
	c.String(http.StatusOK, "hello gin post method")
}

func IndexPutAction(c *gin.Context) {
	c.String(http.StatusOK, "hello gin put method")
}

func IndexDeleteAction(c *gin.Context) {
	c.String(http.StatusOK, "hello gin delete method")
}
