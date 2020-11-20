package controller

import (
	"github.com/ericpai/please/rest/adapter"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context, envelope *adapter.Response)

func Wrapper(fn HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		envelope := new(adapter.Response)
		c.Set("response", envelope)
	}
}
