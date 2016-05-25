package cclog

import (
	"io"
	"os"
	"fmt"
	"log"
	"time"
	"path"
	"runtime"
	"runtime/debug"
)

func Debug(tag string, args ...interface {}) {
	LogInstance.record(LogLevel_Debug, tag, args...)
}

func Warning(tag string, args ...interface {}) {
	LogInstance.record(LogLevel_Warning, tag, args...)
}

func Error(tag string, args ...interface {}) {
	LogInstance.record(LogLevel_Error, tag, args...)
}

func Info(tag string, args ...interface {}) {
	LogInstance.record(LogLevel_Info, tag, args...)
}

func Panic(tag string, args ...interface {}) {
	args = append(args, "\n", string(debug.Stack()))
	LogInstance.record(LogLevel_Panic, tag, args...)
}

func Fatal(tag string, args ...interface {}) {
	args = append(args, "\n", string(debug.Stack()))
	LogInstance.record(LogLevel_Fatal, tag, args...)
//	os.Exit(-1)
}

type LogLevel int

const (
	LogLevel_Debug LogLevel = iota
	LogLevel_Warning
	LogLevel_Error
	LogLevel_Info
	LogLevel_Panic
	LogLevel_Fatal
)

var (
	logLevelShortString = []string{"D", "W", "E", "I", "P", "F"}
	logLevelString = []string{"Debug", "Warning", "Error", "Info", "Panic", "Fatal"}
)

func (level LogLevel) ShortString() string {
	if level < 0 || int(level) > len(logLevelShortString) {
		return "U"
	}
	return logLevelShortString[level]
}

func (level LogLevel) String() string {
	if level < 0 || int(level) > len(logLevelString) {
		return "Unknown"
	}
	return logLevelString[level]
}

var LogInstance *CCLog = new(CCLog)

type logRecord struct {
	timestamp	string
	content		string
}

type CCLog struct {
	m_logLevel	LogLevel
	m_printFlag	bool
	m_saveFlag 	bool

	m_recordChan chan logRecord
}

func (ccLog *CCLog) Init(level LogLevel, print bool, save bool) {
	ccLog.m_logLevel = level
	ccLog.m_printFlag = print
	ccLog.m_saveFlag = save

	if ccLog.m_saveFlag {
		ccLog.m_recordChan = make(chan logRecord, 1024)
		go ccLog.runSave()
	}
}

func (ccLog *CCLog) record(level LogLevel, tag string, args ...interface {}) {
	if level < ccLog.m_logLevel {
		return
	}
	var writer io.Writer
	if LogLevel_Debug == level || LogLevel_Info == level {
		writer = os.Stdout
	}else{
		writer = os.Stderr
	}
	currentTime := time.Now()
	logStr := fmt.Sprintf("%s %s [%s] %s %s", level.ShortString(), currentTime.Format("2006-01-02 15:04:05"), tag, ccLog.callStack(3), fmt.Sprintln(args...))
	writer.Write([]byte(logStr))

	if ccLog.m_saveFlag {
		var record logRecord
		record.timestamp = currentTime.Format("2006-01-02")
		record.content = logStr
		select {
		case ccLog.m_recordChan <- record:
		case <- time.After(time.Millisecond * 100):
		}
	}
}

func (ccLog *CCLog) runSave() {
	defer func() {
		if err := recover(); err != nil {
			Panic("LogSystem", "runSaveLog", err)
		}
	}()
	for record := range ccLog.m_recordChan {
		ccLog.saveToFile(record)
	}
}

func (ccLog *CCLog) saveToFile(record logRecord) {
	filePath := fmt.Sprintf("./log/%s.log", record.timestamp)
	file, fileErr := os.OpenFile(filePath, os.O_APPEND | os.O_RDWR, 0666)
	defer file.Close()
	if nil != fileErr {
		err := os.MkdirAll(path.Dir(filePath), 0777)
		if nil != err {
			log.Println("E", "LogSystem", "SaveToFile MkdirAll", err)
			return
		}
		file, fileErr = os.Create(filePath)
		if nil != fileErr {
			log.Println("E", "LogSystem", "SaveToFile Create", fileErr)
			return
		}
	}
	file.WriteString(record.content)
}

func (ccLog *CCLog) callStack(depth int) string {
	var fileName string
	_, filePath, line, ok := runtime.Caller(depth)
	if !ok {
		fileName = "???"
		line = 0
	}else{
		fileName = path.Base(filePath)
	}
	return fmt.Sprintf("%s:%d", fileName, line)
}
