package main

import (
	"log"

	"github.com/ericpai/please/database"
	"github.com/ericpai/please/module"
	"github.com/ericpai/please/rest"
	"github.com/ericpai/please/schedule"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	c := dig.New()
	database.Init(c)
	module.Init(c)
	rest.Init(c)
	if err := schedule.Init(c); err != nil {
		log.Fatalf("launch schedule task failed: %s\n", err.Error())
	}
	gin.SetMode(gin.ReleaseMode)
	if err := c.Invoke(func(r *gin.Engine) {
		log.Println("Start server at :8080")
		if runErr := r.Run(":8080"); runErr != nil {
			log.Printf("Run engine failed: %s\n", runErr.Error())
		}
	}); err != nil {
		log.Fatalf("launch app failed: %s\n", err.Error())
	}
}
