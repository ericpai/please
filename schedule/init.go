package schedule

import (
	"log"

	"github.com/ericpai/please/module/task"
	"github.com/ericpai/please/module/worker"
	"github.com/robfig/cron/v3"
	"go.uber.org/dig"
)

var cronjob = cron.New()

func Init(c *dig.Container) error {
	return c.Invoke(func(taskModel task.Model, workerModel worker.Model) error {
		if _, err := cronjob.AddFunc("* * * * *", func() {
			RemoteCopy(taskModel, workerModel)
		}); err != nil {
			return err
		}
		cronjob.Start()
		log.Println("Cronjob started")
		return nil
	})
}
