package loggergo_test

import (
	"testing"

	loggergo "github.com/Alonza0314/logger-go"
)

var testDirectLoggerCases = []struct {
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

func TestDirectLogger(t *testing.T) {
	for _, testCase := range testDirectLoggerCases {
		switch testCase.tag {
		case "TEST_INFO":
			loggergo.Info(testCase.tag, testCase.msg)
		case "TEST_ERROR":
			loggergo.Error(testCase.tag, testCase.msg)
		case "TEST_WARN":
			loggergo.Warn(testCase.tag, testCase.msg)
		case "TEST_TEST":
			loggergo.Test(testCase.tag, testCase.msg)
		case "TEST_DEBUG":
			loggergo.Debug(testCase.tag, testCase.msg)
		}
	}
}
