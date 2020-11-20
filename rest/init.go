package rest

import (
	"log"

	"github.com/ericpai/please/rest/controller"
	"go.uber.org/dig"
)

func Init(c *dig.Container) {
	pv := []interface{}{
		controller.NewTaskController,
		NewRouter,
	}
	for _, v := range pv {
		if err := c.Provide(v); err != nil {
			log.Fatalf("provide %v failed: %s", v, err.Error())
		}
	}
}
