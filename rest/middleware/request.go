package middleware

import (
	"fmt"
	"net/http"

	"github.com/ericpai/please/rest/adapter"
	"github.com/gin-gonic/gin"
)

func Request(c *gin.Context) {
	var req adapter.Request
	if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPatch {
		if err := c.ShouldBindJSON(&req); err != nil {
			resp := new(adapter.Response)
			resp.Meta.Code = http.StatusBadRequest
			resp.Meta.Message = fmt.Sprintf("请求数据格式错误: %s", err.Error())
			c.JSON(resp.Meta.Code, resp)
			c.Abort()
			return
		}
	}
	c.Set(KeyRequest, &req)
}
