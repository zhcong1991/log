package log

import (
	"os"
)

type Console struct {
	level Level
}

func (w *Console) Name() string {
	return "console"
}

func (w *Console) Write(r *record) error {
	if r.level < w.level {
		return nil
	}
	os.Stdout.Write(str2byte(r.String()))
	return nil
}

func (w *Console) Flush() error {
	return nil
}

func (w *Console) Split() error {
	return nil
}

func NewConsole(cfg *WriterConfig) *Console {
	return &Console{
		level: cfg.LogLevel,
	}
}
