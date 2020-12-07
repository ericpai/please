package model

import (
	"context"
	"log"
	"sync"

	"github.com/ericpai/please/module/schedule"
	"github.com/ericpai/please/module/task"
	"github.com/ericpai/please/module/worker"
	"github.com/robfig/cron/v3"
)

type defaultModel struct {
	cronJobs    *cron.Cron
	jobEntity   map[int]cron.EntryID
	lock        *sync.Mutex
	taskModel   task.Model
	workerModel worker.Model
}

func (d *defaultModel) Update(id int, sche string) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if v, exists := d.jobEntity[id]; exists {
		d.cronJobs.Remove(v)
		delete(d.jobEntity, id)
		log.Printf("已删除定时任务, id: %d, eid: %d\n", id, v)
	}
	if sche == "" {
		return nil
	}
	eid, err := d.cronJobs.AddFunc(sche, func() {
		RemoteCopy(id, d.taskModel, d.workerModel)
	})
	if err != nil {
		return err
	}
	d.jobEntity[id] = eid
	log.Printf("已注册定时任务, id: %d, eid: %d, schedule: %s\n", id, eid, sche)
	return nil
}

func New(taskModel task.Model, workerModel worker.Model) (schedule.Model, error) {
	m := &defaultModel{
		cronJobs:    cron.New(),
		jobEntity:   make(map[int]cron.EntryID),
		lock:        new(sync.Mutex),
		taskModel:   taskModel,
		workerModel: workerModel,
	}
	ctx := context.Background()
	pos, err := taskModel.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, po := range pos {
		if err := m.Update(int(po.ID), po.Schedule); err != nil {
			return nil, nil
		}
	}
	m.cronJobs.Start()
	return m, nil
}
