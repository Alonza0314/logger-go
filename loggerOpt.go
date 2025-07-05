package loggergo

import "os"

type Option func(*options)

type options struct {
	flag  int
	perm  os.FileMode
	color bool
}

func WithFlag(flag int) Option {
	return func(o *options) {
		o.flag = flag
	}
}

func WithPerm(perm os.FileMode) Option {
	return func(o *options) {
		o.perm = perm
	}
}
