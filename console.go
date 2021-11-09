package log

import (
	"os"
)

type Console struct {
}

func (w *Console) Name() string {
	return "console"
}

func (w *Console) Init() error {
	return nil
}

func (w *Console) Write(r *record) error {
	os.Stdout.Write(str2byte(r.String()))
	return nil
}

func (w *Console) Flush() error {
	return nil
}
