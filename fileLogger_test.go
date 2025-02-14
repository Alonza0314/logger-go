package loggergo_test

import (
	"os"
	"testing"

	loggergo "github.com/Alonza0314/logger-go"
)

var fileLoggerTestCases = []struct {
	tag string
	msg string
}{
	{
		"TEST_INFO",
		"this is a info test",
	},
	{
		"TEST_ERROR",
		"this is a error test",
	},
	{
		"TEST_WARN",
		"this is a warn test",
	},
	{
		"TEST_TEST",
		"this is a test test",
	},
	{
		"TEST_DEBUG",
		"this is a debug test",
	},
}

func TestFileLogger(t *testing.T) {
	logger := loggergo.NewFileLogger("test.log", loggergo.WithFlag(os.O_APPEND|os.O_CREATE|os.O_WRONLY), loggergo.WithPerm(os.FileMode(0644)))
	defer logger.Close()

	for _, testCase := range fileLoggerTestCases {
		switch testCase.tag {
		case "TEST_INFO":
			logger.Info(testCase.tag, testCase.msg)
		case "TEST_ERROR":
			logger.Error(testCase.tag, testCase.msg)
		case "TEST_WARN":
			logger.Warn(testCase.tag, testCase.msg)
		case "TEST_TEST":
			logger.Test(testCase.tag, testCase.msg)
		case "TEST_DEBUG":
			logger.Debug(testCase.tag, testCase.msg)
		}
	}
}
