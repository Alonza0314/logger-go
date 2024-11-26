package loggergo_test

import (
	"testing"

	loggergo "github.com/Alonza0314/logger-go"
)

var testCases = []struct {
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
}

func TestInfo(t *testing.T) {
	for _, testCase := range testCases {
		switch testCase.tag {
		case "TEST_INFO":
			loggergo.Info(testCase.tag, testCase.msg)
		case "TEST_ERROR":
			loggergo.Error(testCase.tag, testCase.msg)
		case "TEST_WARN":
			loggergo.Warn(testCase.tag, testCase.msg)
		case "TEST_TEST":
			loggergo.Test(testCase.tag, testCase.msg)
		}
	}
}
