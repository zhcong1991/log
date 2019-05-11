package log

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

var logger Logger
var logRecordPool sync.Pool

type Logger struct {
	async           bool
	enableCallStack bool

	signalChan chan string
	recordChan chan *logRecord

	channels []LoggerChannel
}

var reactor = make(map[string]newLoggerChannelFunc)

func Register(channelType string, channel newLoggerChannelFunc) {
	if channel == nil {
		panic("cannot register nil logger channel")
	}
	if _, ok := reactor[channelType]; ok {
		panic("cannot repeat register same logger channel")
	}
	reactor[channelType] = channel
}

func (l *Logger) setChannel(channelType string, config string) error {
	newFunc, ok := reactor[channelType]
	if !ok {
		return errors.New("unsupported logger channel type: " + channelType)
	}
	channel := newFunc()
	err := channel.Init(config)
	if err != nil {
		return errors.New("init logger channel type: " + channelType + " error: " + err.Error())
	}
	l.channels = append(l.channels, channel)
	return nil
}

func (l *Logger) writeLog(level LogLevel, tag string, v ...interface{}) {
	timestamp := time.Now()
	log := logShortName[level] + "" + timestamp.Format("2006-01-02 15:04:05") + " [" + tag + "] "

	if l.enableCallStack {
		log += l.callStack(3) + " "
	}
	log += fmt.Sprintln(v...)

	if l.async {
		record := logRecordPool.Get().(*logRecord)
		record.level = level
		record.timestamp = timestamp
		record.content = log
		l.recordChan <- record
	} else {
		for _, channel := range l.channels {
			if err := channel.Write(level, l.str2byte(log)); err != nil {
				fmt.Fprintf(os.Stderr, "unable to write data to channel: %s, error:%v\n", channel.GetType(), err)
			}
		}
	}
}

func (l *Logger) callStack(depth int) string {
	var fileName string
	_, filePath, line, ok := runtime.Caller(depth)
	if !ok {
		fileName = "???"
		line = 0
	} else {
		fileName = path.Base(filePath)
	}
	return fmt.Sprintf("%s:%d", fileName, line)
}

func (l *Logger) str2byte(str string) []byte {
	array := (*[2]uintptr)(unsafe.Pointer(&str))
	tmp := [3]uintptr{array[0], array[1], array[1]}
	return *(*[]byte)(unsafe.Pointer(&tmp))
}

func init() {
	logger.enableCallStack = true
}
