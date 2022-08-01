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
	if r.level == WARN || r.level == ERROR || r.level == FATAL {
		os.Stderr.WriteString(r.String())
	} else {
		os.Stdout.WriteString(r.String())
	}
	return nil
}

func (w *Console) Flush() error {
	return nil
}

func NewConsole(cfg *WriterConfig) *Console {
	return &Console{
		level: logNameToLevel[cfg.LogLevel],
	}
}
