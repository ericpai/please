package model

import (
	"context"
	"log"
	"os/exec"
	"time"

	"github.com/ericpai/please/module/worker"
)

type windowsModel struct {
}

func NewWindowsModel() worker.Model {
	return new(windowsModel)
}

func (d *windowsModel) Copyfile(ctx context.Context, workCtx worker.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Hour)
	defer cancel()
	out, err := exec.CommandContext(ctx, "copy.bat").Output()
	if err != nil {
		return err
	}
	log.Println(string(out))
	return nil
}
