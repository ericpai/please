package db

import (
	"context"

	"github.com/ericpai/please/module/task"
	"gorm.io/gorm"
)

type defaultDB struct {
	gormDB *gorm.DB
}

func New(gormDB *gorm.DB) task.DB {
	if err := gormDB.AutoMigrate(new(task.PO)); err != nil {
		panic(err.Error())
	}
	return &defaultDB{gormDB: gormDB}
}

func (d *defaultDB) Insert(ctx context.Context, v task.PO) (task.PO, error) {
	result := d.gormDB.WithContext(ctx).Create(&v)
	return v, result.Error
}

func (d *defaultDB) Update(ctx context.Context, v task.PO) (task.PO, error) {
	result := d.gormDB.WithContext(ctx).Select("*").Updates(v)
	return v, result.Error
}

func (d *defaultDB) SelectAll(ctx context.Context) ([]task.PO, error) {
	var pos []task.PO
	result := d.gormDB.WithContext(ctx).Find(&pos)
	return pos, result.Error
}

func (d *defaultDB) SelectByID(ctx context.Context, id uint) (task.PO, error) {
	var po task.PO
	result := d.gormDB.WithContext(ctx).First(&po, id)
	return po, result.Error
}

func (d *defaultDB) Delete(ctx context.Context, id uint) error {
	return d.gormDB.Unscoped().WithContext(ctx).Delete(new(task.PO), id).Error
}
