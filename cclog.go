package cclog

import (
	"log"
)

const (
	LogLevel_Debug = iota
	LogLevel_Warning
	LogLevel_Error
	LogLevel_Info
	LogLevel_Panic
	LogLevel_Fatal
)

func Debug(args ...interface {}) {

}

func Warning(args ...interface {}) {

}

func Error(args ...interface {}) {

}

func Info(args ...interface {}) {

}

func Panic(args ...interface {}) {

}

func Fatal(args ...interface {}) {

}

type CCLog struct {

}

func (this *CCLog) test() {
	log.Println("CCLog")
}
