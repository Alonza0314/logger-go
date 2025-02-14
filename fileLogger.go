package loggergo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type fileLogger struct {
	file   *os.File
	logger *log.Logger
	color  bool
}

func NewFileLogger(loggerFilePath string, opts ...Option) *fileLogger {
	options := &options{
		flag:  DEFAULT_FLAG,
		perm:  DEFAULT_PERM,
		color: DEFAULT_COLOR,
	}

	for _, opt := range opts {
		opt(options)
	}

	if !filepath.IsAbs(loggerFilePath) {
		absPath, err := filepath.Abs(loggerFilePath)
		if err != nil {
			panic(fmt.Errorf("invalid file path: %v", err))
		}
		loggerFilePath = absPath
	}

	dir := filepath.Dir(loggerFilePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Errorf("failed to create directory: %v", err))
	}

	file, err := os.OpenFile(loggerFilePath, options.flag, options.perm)
	if err != nil {
		panic(err)
	}
	logger := log.New(file, "", log.Ldate|log.Ltime)
	return &fileLogger{
		file:   file,
		logger: logger,
	}
}

func (l *fileLogger) Close() {
	l.file.Close()
}

func (l *fileLogger) Info(tag string, msg string) {
	if l.color {
		l.logger.Printf("%s[INFO]%s [%s]%s %s\n", COLOR_BLUE, COLOR_BLUE, tag, COLOR_RESET, msg)
	} else {
		l.logger.Printf("[INFO] [%s] %s\n", tag, msg)
	}
}

func (l *fileLogger) Error(tag string, msg string) {
	if l.color {
		l.logger.Printf("%s[EROR]%s [%s]%s %s\n", COLOR_RED, COLOR_BLUE, tag, COLOR_RESET, msg)
	} else {
		l.logger.Printf("[EROR] [%s] %s\n", tag, msg)
	}
}

func (l *fileLogger) Warn(tag string, msg string) {
	if l.color {
		l.logger.Printf("%s[WARN]%s [%s]%s %s\n", COLOR_YELLOW, COLOR_BLUE, tag, COLOR_RESET, msg)
	} else {
		l.logger.Printf("[WARN] [%s] %s\n", tag, msg)
	}
}

func (l *fileLogger) Test(tag string, msg string) {
	if l.color {
		l.logger.Printf("%s[TEST]%s [%s]%s %s\n", COLOR_GREEN, COLOR_BLUE, tag, COLOR_RESET, msg)
	} else {
		l.logger.Printf("[TEST] [%s] %s\n", tag, msg)
	}
}

func (l *fileLogger) Debug(tag string, msg string) {
	if l.color {
		l.logger.Printf("%s[DBUG]%s [%s]%s %s\n", COLOR_PURPLE, COLOR_BLUE, tag, COLOR_RESET, msg)
	} else {
		l.logger.Printf("[DBUG] [%s] %s\n", tag, msg)
	}
}
