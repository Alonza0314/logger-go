package util

import "fmt"

func LevelStringToLevel(level LogLevelString) LogLevel {
	switch level {
	case LEVEL_STRING_ERROR:
		return LEVEL_ERROR
	case LEVEL_STRING_WARN:
		return LEVEL_WARN
	case LEVEL_STRING_INFO:
		return LEVEL_INFO
	case LEVEL_STRING_DEBUG:
		return LEVEL_DEBUG
	case LEVEL_STRING_TRACE:
		return LEVEL_TRACE
	case LEVEL_STRING_TEST:
		return LEVEL_TEST
	default:
		panic(fmt.Errorf("invalid level string: %s, must be one of %s, %s, %s, %s, %s, %s", level, LEVEL_STRING_ERROR, LEVEL_STRING_WARN, LEVEL_STRING_INFO, LEVEL_STRING_DEBUG, LEVEL_STRING_TRACE, LEVEL_STRING_TEST))
	}
}
