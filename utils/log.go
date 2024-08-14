package utils

import (
	"fmt"
	"os"
	"time"
)

type Log struct {
	file   os.File
	silent bool
}

var instance *Log

func NewLog(fileName string) (*Log, error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &Log{file: *file, silent: false}, nil
}

func GetLog() (*Log, error) {
	var err error
	logFile := os.Getenv("LOG_FILE")

	if instance == nil {
		instance, err = NewLog(logFile)
		if err != nil {
			return nil, err
		}
	}

	return instance, nil
}

func (l *Log) SetSilent(silent bool) {
	l.silent = silent
}

func (l *Log) IsSilent() bool {
	return l.silent
}

func (l *Log) WriteLog(msg string) error {
	date := time.Now().Format("2006-01-02 15:04:05")
	text := fmt.Sprintf("[%s] %s", date, msg)

	_, err := l.file.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}

func (l *Log) Printf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)

	if !l.silent {
		fmt.Print(text)
	}

	if err := l.WriteLog(text); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func (l *Log) PrintfErr(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	text = fmt.Sprintf("[ERROR] %s", text)

	if !l.silent {
		fmt.Fprint(os.Stderr, text)
	}

	if err := l.WriteLog(text); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func (l *Log) Close() error {
	return l.file.Close()
}
