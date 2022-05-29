package log

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	level     Level
	fileSplit string
	close     chan bool
	tunnel    chan *record

	writers []Writer
}

func (l *Logger) Init(cfg *Config) {
	l.level = logNameToLevel[cfg.LogLevel]
	for index, item := range cfg.Writers {
		var writer Writer
		switch item.Type {
		case "console":
			writer = NewConsole(&cfg.Writers[index])
		case "file":
			writer = NewFileWriter(&cfg.Writers[index])
		default:
			panic("unsupported writer type: " + item.Type)
		}
		l.writers = append(l.writers, writer)
	}

	go l.asyncWrite()
}

func (l *Logger) Close() {
	close(l.tunnel)
	<-l.close

	for _, w := range l.writers {
		if err := w.Flush(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "flush writer: %s failed, err: %v\n", w.Name(), err)
		}
	}
}

func (l *Logger) writeLog(level Level, tag string, f string, v ...interface{}) {
	if level < l.level {
		return
	}

	r := recordPool.Get().(*record)
	r.level = level
	r.ctx = nil
	r.timestamp = time.Now()
	r.sourceCode = callStack(3)
	r.tag = tag
	r.content = fmt.Sprintf(f, v...)
	l.tunnel <- r
}

func (l *Logger) writeLogCtx(level Level, ctx context.Context, tag string, f string, v ...interface{}) {
	if level < l.level {
		return
	}

	r := recordPool.Get().(*record)
	r.level = level
	r.ctx = ctx
	r.timestamp = time.Now()
	r.sourceCode = callStack(3)
	r.tag = tag
	r.content = fmt.Sprintf(f, v...)
	l.tunnel <- r
}

func (l *Logger) asyncWrite() {
	r, ok := <-l.tunnel
	if !ok {
		l.close <- true
		return
	}

	for _, w := range l.writers {
		if err := w.Write(r); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "write to writer: %s failed, err: %v\n", w.Name(), err)
		}
	}

	flushTimer := time.NewTimer(time.Millisecond * 1000)
	splitTimer := time.NewTimer(time.Second * 5)
	for {
		select {
		case r, ok = <-l.tunnel:
			if !ok {
				l.close <- true
				return
			}
			for _, w := range l.writers {
				if err := w.Write(r); err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "write to writer: %s failed, err: %v\n", w.Name(), err)
				}
			}
			recordPool.Put(r)
		case <-splitTimer.C:
			for _, w := range l.writers {
				if err := w.Split(); err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "split writer: %s failed, err: %v\n", w.Name(), err)
				}
			}
			splitTimer.Reset(time.Second * 5)
		case <-flushTimer.C:
			for _, w := range l.writers {
				if err := w.Flush(); err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "flush writer: %s failed, err: %v\n", w.Name(), err)
				}
			}
			flushTimer.Reset(time.Millisecond * 1000)
		}
	}
}

func NewLogger() *Logger {
	l := &Logger{
		level:   DEBUG,
		close:   make(chan bool, 1),
		tunnel:  make(chan *record, tunnelDefaultSize),
		writers: make([]Writer, 0),
	}

	return l
}

var _logger *Logger
var recordPool sync.Pool

func init() {
	_logger = NewLogger()
	recordPool = sync.Pool{New: func() interface{} {
		return &record{}
	}}
}
