package controller

import (
	"github.com/ericpai/please/rest/adapter"
	"github.com/ericpai/please/rest/middleware"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context, req *adapter.Request, resp *adapter.Response)

func Wrapper(fn HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		respObj, _ := c.Get(middleware.KeyResponse)
		resp := respObj.(*adapter.Response)
		reqObj, _ := c.Get(middleware.KeyRequest)
		req := reqObj.(*adapter.Request)
		fn(c, req, resp)
	}
}
