package module

import (
	"log"

	taskdb "github.com/ericpai/please/module/task/db"
	taskmodel "github.com/ericpai/please/module/task/model"
	workermodel "github.com/ericpai/please/module/worker/model"
	"go.uber.org/dig"
)

func Init(c *dig.Container) {
	pv := []interface{}{
		taskdb.New,
		taskmodel.New,
		workermodel.NewWindowsModel,
	}
	for _, v := range pv {
		if err := c.Provide(v); err != nil {
			log.Fatalf("provide %v failed: %s", v, err.Error())
		}
	}
}
