package model

import (
	"context"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/ericpai/please/module/worker"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type windowsModel struct {
}

func NewWindowsModel() worker.Model {
	return new(windowsModel)
}

func (d *windowsModel) Copyfile(ctx context.Context, workCtx worker.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Hour)
	defer cancel()
	sourcePath := strings.Replace(workCtx.SourcePath, ":", "$", 1)
	out, err := exec.CommandContext(ctx, "copy.bat", workCtx.Address, workCtx.User, workCtx.Password, sourcePath, workCtx.DestPath).Output()
	if err != nil {
		return err
	}

	decoder := simplifiedchinese.GBK.NewDecoder()
	st, err := decoder.Bytes(out)
	if err != nil {
		return err
	}
	log.Println(string(st))
	return nil
}
