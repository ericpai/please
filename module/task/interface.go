package task

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

const (
	BackendWindows = "windows"
	BackendLinux   = "linux"
)

var (
	ErrInvalidParam = errors.New("非法字段")
	ErrNotFound     = errors.New("未找到")
)

type PO struct {
	gorm.Model
	Address     string `gorm:"unique"`
	User        string
	Password    string
	SourcePath  string
	DestPath    string
	Backend     string
	Schedule    string
	Succeed     bool
	Enabled     bool
	CreatedTime time.Time
	UpdatedTime time.Time
}

func (PO) TableName() string {
	return "tasks"
}

type DB interface {
	Insert(ctx context.Context, v PO) (PO, error)
	Update(ctx context.Context, newV PO) (PO, error)
	SelectAll(ctx context.Context) ([]PO, error)
	SelectByID(ctx context.Context, id uint) (PO, error)
	Delete(ctx context.Context, id uint) error
}

type Model interface {
	GetAll(ctx context.Context) ([]PO, error)
	GetByID(ctx context.Context, id int) (PO, error)
	Delete(ctx context.Context, id int) error
	Insert(ctx context.Context, v PO) (PO, error)
	Update(ctx context.Context, id int, v PO) (PO, error)
}
