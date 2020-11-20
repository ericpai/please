package middleware

import (
	"github.com/ericpai/please/rest/adapter"
	"github.com/gin-gonic/gin"
)

const (
	KeyResponse = "response"
)

func Response(c *gin.Context) {
	c.Next()
	envelopeObj, _ := c.Get(KeyResponse)
	envelope := envelopeObj.(*adapter.Response)
	c.JSON(envelope.Meta.Code, envelope)
}
