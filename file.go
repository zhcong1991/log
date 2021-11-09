package log

import (
	"os"
)

type FileWriter struct {
}

func (w *FileWriter) Name() string {
	return "file"
}

func (w *FileWriter) Init() error {
	return nil
}

func (w *FileWriter) Write(r *record) error {
	os.Stdout.Write(str2byte(r.String()))
	return nil
}

func (w *FileWriter) Flush() error {
	return nil
}
