package log

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Logger struct {
	level  Level
	close  chan bool
	tunnel chan *record

	writers []Writer
}

func (l *Logger) Init(cfg *Config) {
	l.level = cfg.level
	l.writers = append(l.writers, &Console{})

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

	for {
		select {
		case r, ok = <-l.tunnel:
			if !ok {
				l.close <- true
				return
			}
			for _, w := range l.writers {
				if err := w.Write(r); err != nil {
					log.Println(err)
				}
			}
			recordPool.Put(r)
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