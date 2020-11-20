package task

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	BackendWindows = "windows"
	BackendLinux   = "linux"
)

type PO struct {
	gorm.Model
	Address     string
	User        string
	Password    string
	SourcePath  string
	DestPath    string
	Backend     string
	Succeed     bool
	Enabled     bool
	CreatedTime time.Time
	UpdatedTime time.Time
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
}
