package worker

import "context"

type Model interface {
	Copyfile(ctx context.Context, workCtx Context) error
}

type Context struct {
	Address    string
	User       string
	Password   string
	SourcePath string
	DestPath   string
}
