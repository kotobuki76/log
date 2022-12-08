package log

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
)

var logLevelName = map[int64]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
	4: "CRITICAL",
}

var LogContext context.Context

func SetContext(ctx context.Context) {
	LogContext = ctx
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	Logf(0, format, args...)
}

// Infof is like Debugf, but at Info level.
func Infof(ctx context.Context, format string, args ...interface{}) {
	Logf(1, format, args...)
}

func InfofNoCtx(format string, args ...interface{}) {
	//fmt.Println("InfofNoCtx")
	Infof(nil, format, args...)
	//s := fmt.Sprintf(format, args...)
	//fmt.Println(os.Getenv("PROCESS_ID") + " " + logLevelName[1] + ": " + s)
}

// Warningf is like Debugf, but at Warning level.
func Warningf(ctx context.Context, format string, args ...interface{}) {
	Logf(2, format, args...)
}

// Errorf is like Debugf, but at Error level.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	Logf(3, format, args...)
}

// Criticalf is like Debugf, but at Critical level.
func Criticalf(ctx context.Context, format string, args ...interface{}) {
	Logf(4, format, args...)
}

func Logf(level int64, format string, args ...interface{}) {
	LogFile(level, format, args...)
}

func LogFile(level int64, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	s = strings.TrimRight(s, "\n")
	log.Print(os.Getenv("PROCESS_ID") + " " + logLevelName[level] + ": " + s)
}

func WriteFile(format string) {
	f, err := os.OpenFile("./log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(format)
}
