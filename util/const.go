package util

type LogLevel int
type LogLevelString string

const (
	LEVEL_ERROR LogLevel = 0
	LEVEL_WARN  LogLevel = 1
	LEVEL_INFO  LogLevel = 2
	LEVEL_DEBUG LogLevel = 3
	LEVEL_TRACE LogLevel = 4
	LEVEL_TEST  LogLevel = 5

	LEVEL_STRING_ERROR LogLevelString = "error"
	LEVEL_STRING_WARN  LogLevelString = "warn"
	LEVEL_STRING_INFO  LogLevelString = "info"
	LEVEL_STRING_DEBUG LogLevelString = "debug"
	LEVEL_STRING_TRACE LogLevelString = "trace"
	LEVEL_STRING_TEST  LogLevelString = "test"
)
