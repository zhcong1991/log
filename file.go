package log

import (
	"bufio"
	"os"
	"path"
	"time"

	"github.com/pkg/errors"
)

type FileWriter struct {
	level          Level
	filePath       string
	splitUnit      string
	backupFileName string
	file           *os.File
	writer         *bufio.Writer
}

func (w *FileWriter) Name() string {
	return "file"
}

func (w *FileWriter) Write(r *record) error {
	if r.level < w.level {
		return nil
	}
	if w.writer == nil {
		return errors.New("no opened file")
	}
	if _, err := w.writer.WriteString(r.String()); err != nil {
		return err
	}
	return nil
}

func (w *FileWriter) Flush() error {
	if w.writer == nil {
		return nil
	}
	if err := w.writer.Flush(); err != nil {
		return err
	}
	if w.getFilePath() != w.backupFileName {
		return w.openExistingOrNew()
	}

	return nil
}

func (w *FileWriter) getFilePath() string {
	now := time.Now()

	suffix := ""
	switch w.splitUnit {
	case "year":
		suffix = "." + now.Format("2006")
	case "month":
		suffix = "." + now.Format("200601")
	case "day":
		suffix = "." + now.Format("20060102")
	case "hour":
		suffix = "." + now.Format("2006010215")
	}
	return w.filePath + suffix
}

func (w *FileWriter) openExistingOrNew() error {
	if err := os.MkdirAll(path.Dir(w.filePath), 0755); err != nil {
		if !os.IsExist(err) {
			return errors.WithMessage(err, "file exist")
		}
	}

	if w.file != nil {
		if err := w.file.Close(); err != nil {
			return errors.WithMessage(err, "close failed")
		}

		if err := os.Rename(w.filePath, w.backupFileName); err != nil {
			return errors.WithMessage(err, "rename log file failed")
		}
		w.file = nil
		w.backupFileName = w.getFilePath()
	}

	if file, err := os.OpenFile(w.filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err != nil {
		return errors.WithMessage(err, "open file failed")
	} else {
		w.file = file
	}

	if w.writer = bufio.NewWriterSize(w.file, 10240); w.writer == nil {
		return errors.New("create file writer failed")
	}

	return nil
}

func NewFileWriter(cfg *WriterConfig) *FileWriter {
	w := &FileWriter{
		level:     logNameToLevel[cfg.LogLevel],
		filePath:  cfg.FilePath,
		splitUnit: cfg.SplitUnit,
	}
	w.backupFileName = w.getFilePath()
	err := w.openExistingOrNew()
	if err != nil {
		panic(err)
	}
	return w
}
