package loggergo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Alonza0314/logger-go/v2/model"
	"github.com/Alonza0314/logger-go/v2/util"
)

type Logger struct {
	file      *os.File
	logger    *log.Logger
	level     util.LogLevel
	debugMode bool
}

func NewLogger(loggerFilePath string, debugMode bool, opts ...Option) *Logger {
	options := &options{
		flag: DEFAULT_FLAG,
		perm: DEFAULT_PERM,
	}

	for _, opt := range opts {
		opt(options)
	}

	if debugMode {
		return &Logger{
			file:      nil,
			logger:    log.New(os.Stdout, "", log.Ldate|log.Ltime),
			level:     DEFAULT_LEVEL,
			debugMode: debugMode,
		}
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
	return &Logger{
		file:      file,
		logger:    logger,
		level:     DEFAULT_LEVEL,
		debugMode: debugMode,
	}
}

func (l *Logger) SetLevel(level util.LogLevelString) {
	l.level = util.LevelStringToLevel(level)
}

func (l *Logger) WithTag(tag string) model.LoggerInterface {
	return &LoggerImplementation{
		logger:    l.logger,
		level:     l.level,
		tags:      []string{tag},
		debugMode: l.debugMode,
	}
}

func (l *Logger) WithTags(tags ...string) model.LoggerInterface {
	return &LoggerImplementation{
		logger:    l.logger,
		level:     l.level,
		tags:      tags,
		debugMode: l.debugMode,
	}
}

func (l *Logger) Close() {
	if l.file == nil {
		return
	}
	if err := l.file.Close(); err != nil {
		panic(err)
	}
}
