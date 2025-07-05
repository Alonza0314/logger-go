package loggergo_test

import (
	"os"
	"path/filepath"
	"testing"

	loggergo "github.com/Alonza0314/logger-go"
	"github.com/Alonza0314/logger-go/model"
	"github.com/Alonza0314/logger-go/util"
)

func TestLoggerCreation(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "test.log")

	logger := loggergo.NewLogger(logFile, true)
	if logger == nil {
		t.Fatal("Failed to create logger")
	}
	defer logger.Close()

	logger.SetLevel(util.LEVEL_STRING_DEBUG)
}

var testTagCases = []struct {
	name    string
	tags    []string
	message string
}{
	{
		name:    "Single Tag",
		tags:    []string{"TestTag"},
		message: "Test single tag",
	},
	{
		name:    "Multiple Tags",
		tags:    []string{"Tag1", "Tag2"},
		message: "Test multiple tags",
	},
}

func TestLoggerTags(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "tags.log")
	logger := loggergo.NewLogger(logFile, true)
	defer logger.Close()

	for _, tc := range testTagCases {
		t.Run(tc.name, func(t *testing.T) {
			var taggedLogger model.LoggerInterface
			if len(tc.tags) == 1 {
				taggedLogger = logger.WithTag(tc.tags[0])
			} else {
				taggedLogger = logger.WithTags(tc.tags...)
			}

			if taggedLogger == nil {
				t.Fatal("Failed to create tagged logger")
			}
			taggedLogger.Infof(tc.message)
		})
	}
}

var testFormatMethodCases = []struct {
	name     string
	message  string
	expected string
}{
	{"Debug", "Debug message: %s", "test"},
	{"Info", "Info message: %s", "test"},
	{"Warn", "Warn message: %s", "test"},
	{"Error", "Error message: %s", "test"},
	{"Trace", "Trace message: %s", "test"},
	{"Test", "Test message: %s", "test"},
}

func TestLoggerFormatMethods(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "format.log")
	logger := loggergo.NewLogger(logFile, true)
	defer logger.Close()
	taggedLogger := logger.WithTag("FormatTest")

	logFuncs := map[string]func(format string, args ...interface{}){
		"Debug": taggedLogger.Debugf,
		"Info":  taggedLogger.Infof,
		"Warn":  taggedLogger.Warnf,
		"Error": taggedLogger.Errorf,
		"Trace": taggedLogger.Tracef,
		"Test":  taggedLogger.Testf,
	}

	for _, tc := range testFormatMethodCases {
		t.Run(tc.name, func(t *testing.T) {
			logFunc := logFuncs[tc.name]
			if logFunc == nil {
				t.Fatalf("Log function not found for level: %s", tc.name)
			}
			logFunc(tc.message, tc.expected)
		})
	}
}

var testLineMethodCases = []struct {
	name    string
	message string
}{
	{"Debug", "Debug message"},
	{"Info", "Info message"},
	{"Warn", "Warn message"},
	{"Error", "Error message"},
	{"Trace", "Trace message"},
	{"Test", "Test message"},
}

func TestLoggerLineMethods(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "line.log")
	logger := loggergo.NewLogger(logFile, true)
	defer logger.Close()
	taggedLogger := logger.WithTag("LineTest")

	logFuncs := map[string]func(args ...interface{}){
		"Debug": taggedLogger.Debugln,
		"Info":  taggedLogger.Infoln,
		"Warn":  taggedLogger.Warnln,
		"Error": taggedLogger.Errorln,
		"Trace": taggedLogger.Traceln,
		"Test":  taggedLogger.Testln,
	}

	for _, tc := range testLineMethodCases {
		t.Run(tc.name, func(t *testing.T) {
			logFunc := logFuncs[tc.name]
			if logFunc == nil {
				t.Fatalf("Log function not found for level: %s", tc.name)
			}
			logFunc(tc.message)
		})
	}
}

var testLoggerOptionCases = []struct {
	name    string
	options []loggergo.Option
}{
	{
		name: "Custom Flags",
		options: []loggergo.Option{
			loggergo.WithFlag(os.O_APPEND | os.O_CREATE | os.O_WRONLY),
		},
	},
	{
		name: "Custom Permissions",
		options: []loggergo.Option{
			loggergo.WithPerm(0644),
		},
	},
	{
		name: "Combined Options",
		options: []loggergo.Option{
			loggergo.WithFlag(os.O_APPEND | os.O_CREATE | os.O_WRONLY),
			loggergo.WithPerm(0644),
		},
	},
}

func TestLoggerOptions(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "options.log")

	for _, tc := range testLoggerOptionCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := loggergo.NewLogger(logFile, true, tc.options...)
			defer logger.Close()
			if logger == nil {
				t.Fatal("Failed to create logger with custom options")
			}
			logger.WithTag(tc.name).Infoln("Test with custom options")
		})
	}
}
