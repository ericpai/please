package middleware

import (
	"github.com/ericpai/please/rest/adapter"
	"github.com/gin-gonic/gin"
)

const (
	KeyResponse = "response"
	KeyRequest  = "request"
)

func Response(c *gin.Context) {
	resp := new(adapter.Response)
	c.Set(KeyResponse, resp)
	c.Next()
	c.JSON(resp.Meta.Code, resp)
}
