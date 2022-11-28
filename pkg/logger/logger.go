package logger

import (
	"fmt"
	"io/ioutil"

	"log"
	"os"
	"runtime"
	dbg "runtime/debug"
	"strings"
	"sync"
)

const (
	errorLevel string = "ERROR"
	warnLevel  string = "WARN "
	infoLevel  string = "INFO "
	debugLevel string = "DEBUG"
)

const (
	debug bool = true
)

type Logger struct {
	loggers map[string]*log.Logger
}

var (
	instance *Logger

	initLog sync.Once
)

func GetLogger() *Logger {
	initLog.Do(func() {
		loggers := make(map[string]*log.Logger)

		loggers[errorLevel] = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
		loggers[warnLevel] = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
		loggers[infoLevel] = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
		if debug {
			loggers[debugLevel] = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
		} else {
			loggers[debugLevel] = log.New(ioutil.Discard, "", log.LstdFlags|log.Lmicroseconds)
		}

		instance = &Logger{loggers: loggers}
	})

	return instance
}

func (rlog *Logger) logInt(level, msg string, err error) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		fileParts := strings.Split(file, "/")
		msg = fmt.Sprintf("%s %s:%d %s", level, fileParts[len(fileParts)-1], line, msg)
	}

	errMsg := ""
	if err != nil {
		errMsg = fmt.Sprintf("[%s]", err)
	}
	if level == errorLevel {
		errMsg += fmt.Sprintf("\nThe stacktrace is: %s", string(dbg.Stack()))
	}

	logger := rlog.loggers[level]
	logger.Println(msg, errMsg)
}

func Err(msg string, err error) {
	GetLogger().logInt(errorLevel, msg, err)
}

func Warn(msg string, err error) {
	GetLogger().logInt(warnLevel, msg, err)
}

func Info(msg string) {
	GetLogger().logInt(infoLevel, msg, nil)
}

func Debug(msg string) {
	GetLogger().logInt(debugLevel, msg, nil)
}
