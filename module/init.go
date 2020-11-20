package module

import (
	"log"

	"github.com/ericpai/please/module/task/db"
	"github.com/ericpai/please/module/task/model"
	"go.uber.org/dig"
)

func Init(c *dig.Container) {
	pv := []interface{}{
		db.New,
		model.New,
	}
	for _, v := range pv {
		if err := c.Provide(v); err != nil {
			log.Fatalf("provide %v failed: %s", v, err.Error())
		}
	}
}
