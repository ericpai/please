package adapter

import (
	"strconv"

	"github.com/ericpai/please/module/task"
)

type Task struct {
	ID          string `json:"id"`
	Address     string `json:"address"`
	User        string `json:"user"`
	Password    string `json:"password"`
	SourcePath  string `json:"sourcePath"`
	DestPath    string `json:"destPath"`
	Backend     string `json:"backend"`
	Succeed     bool   `json:"succeed"`
	Enabled     bool   `json:"enabled"`
	CreatedTime int64  `json:"createdTime"`
	UpdatedTime int64  `json:"updatedTime"`
}

func TaskModelToExternal(p task.PO) Task {
	return Task{
		ID:          strconv.Itoa(int(p.ID)),
		Address:     p.Address,
		User:        p.User,
		Password:    p.Password,
		SourcePath:  p.SourcePath,
		DestPath:    p.DestPath,
		Backend:     p.Backend,
		Succeed:     p.Succeed,
		Enabled:     p.Enabled,
		CreatedTime: p.CreatedTime.Unix(),
		UpdatedTime: p.UpdatedTime.Unix(),
	}
}

func TaskExternalToModel(p Task) task.PO {
	return task.PO{
		Address:    p.Address,
		User:       p.User,
		Password:   p.Password,
		SourcePath: p.SourcePath,
		DestPath:   p.DestPath,
		Backend:    p.Backend,
		Succeed:    p.Succeed,
		Enabled:    p.Enabled,
	}
}
