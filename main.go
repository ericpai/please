package main

import (
	"log"

	"github.com/ericpai/please/database"
	"github.com/ericpai/please/module"
	"github.com/ericpai/please/rest"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	c := dig.New()
	database.Init(c)
	module.Init(c)
	rest.Init(c)
	gin.SetMode(gin.ReleaseMode)
	if err := c.Invoke(func(r *gin.Engine) {
		if runErr := r.Run(":8080"); runErr != nil {
			log.Printf("Run engine failed: %s\n", runErr.Error())
		}
	}); err != nil {
		log.Fatalf("launch app failed: %s\n", err.Error())
	}
}
