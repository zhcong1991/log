package log

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"time"
	"unsafe"
)

type Level = int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
	MAX
)

const tunnelDefaultSize = 5 * 1024

var (
	logName = [MAX]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

type WriterConfig struct {
	Type      string `json:"type" toml:"type"`
	LogLevel  Level  `json:"log_level" toml:"log_level"`
	FilePath  string `json:"file_path" toml:"file_path"`
	SplitUnit string `json:"split_unit" toml:"split_unit"`
}

type Config struct {
	LogLevel Level          `json:"log_level" toml:"log_level"`
	Writes   []WriterConfig `json:"writes" toml:"writes"`
}

type record struct {
	level      Level
	ctx        context.Context
	timestamp  time.Time
	sourceCode string
	tag        string
	content    string
}

func (r *record) String() string {
	if r.ctx != nil && r.ctx.Value("trace") != nil {
		trace := r.ctx.Value("trace").(string)
		return fmt.Sprintf("[%s][%s][%s][%s]||trace=%s||%s\n", logName[r.level], r.timestamp.Format("2006-01-02T15:04:05.000-0700"), r.sourceCode, r.tag, trace, r.content)
	}
	return fmt.Sprintf("[%s][%s][%s][%s]||%s\n", logName[r.level], r.timestamp.Format("2006-01-02T15:04:05.000-0700"), r.sourceCode, r.tag, r.content)
}

type Writer interface {
	Name() string
	Write(r *record) error
	Flush() error
	Split() error
}

func callStack(depth int) string {
	var fileName string
	_, filePath, line, ok := runtime.Caller(depth)
	if !ok {
		fileName = "???.go"
		line = 0
	} else {
		fileName = path.Base(filePath)
	}
	return fmt.Sprintf("%s:%d", fileName, line)
}

func str2byte(str string) []byte {
	array := (*[2]uintptr)(unsafe.Pointer(&str))
	tmp := [3]uintptr{array[0], array[1], array[1]}
	return *(*[]byte)(unsafe.Pointer(&tmp))
}
