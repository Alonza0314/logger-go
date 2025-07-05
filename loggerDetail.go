package loggergo

import (
	"fmt"
	"log"

	"github.com/Alonza0314/logger-go/util"
)

type LoggerImplementation struct {
	logger    *log.Logger
	level     util.LogLevel
	tags      []string
	debugMode bool
}

func (l *LoggerImplementation) checkLevel(level util.LogLevel) bool {
	if l.level >= level {
		return true
	}
	return false
}

func (l *LoggerImplementation) getTags() string {
	tags := ""
	for _, tag := range l.tags {
		tags += fmt.Sprintf("[%s]", tag)
	}
	return tags
}

func (l *LoggerImplementation) Errorf(format string, args ...interface{}) {
	if !l.checkLevel(util.LEVEL_ERROR) {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log.Printf("%s%s%s %s%s %s\n", COLOR_RED, TOP_TAG_ERROR, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s\n", COLOR_RED, TOP_TAG_ERROR, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Warnf(format string, args ...interface{}) {
	if !l.checkLevel(util.LEVEL_WARN) {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log.Printf("%s%s%s %s%s %s\n", COLOR_YELLOW, TOP_TAG_WARN, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s\n", COLOR_YELLOW, TOP_TAG_WARN, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Infof(format string, args ...interface{}) {
	if !l.checkLevel(util.LEVEL_INFO) {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log.Printf("%s%s%s %s%s %s\n", COLOR_BLUE, TOP_TAG_INFO, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s\n", COLOR_BLUE, TOP_TAG_INFO, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Debugf(format string, args ...interface{}) {
	if !l.checkLevel(util.LEVEL_DEBUG) {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log.Printf("%s%s%s %s%s %s\n", COLOR_PURPLE, TOP_TAG_DEBUG, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s\n", COLOR_PURPLE, TOP_TAG_DEBUG, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Tracef(format string, args ...interface{}) {
	if !l.checkLevel(util.LEVEL_TRACE) {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log.Printf("%s%s%s %s%s %s\n", COLOR_RESET, TOP_TAG_TRACE, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s\n", COLOR_RESET, TOP_TAG_TRACE, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Testf(format string, args ...interface{}) {
	if !l.checkLevel(util.LEVEL_TEST) {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log.Printf("%s%s%s %s%s %s\n", COLOR_GREEN, TOP_TAG_TEST, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s\n", COLOR_GREEN, TOP_TAG_TEST, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Errorln(args ...interface{}) {
	if !l.checkLevel(util.LEVEL_ERROR) {
		return
	}

	msg := fmt.Sprintln(args...)
	log.Printf("%s%s%s %s%s %s", COLOR_RED, TOP_TAG_ERROR, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s", COLOR_RED, TOP_TAG_ERROR, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Warnln(args ...interface{}) {
	if !l.checkLevel(util.LEVEL_WARN) {
		return
	}

	msg := fmt.Sprintln(args...)
	log.Printf("%s%s%s %s%s %s", COLOR_YELLOW, TOP_TAG_WARN, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s", COLOR_YELLOW, TOP_TAG_WARN, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Infoln(args ...interface{}) {
	if !l.checkLevel(util.LEVEL_INFO) {
		return
	}

	msg := fmt.Sprintln(args...)
	log.Printf("%s%s%s %s%s %s", COLOR_BLUE, TOP_TAG_INFO, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s", COLOR_BLUE, TOP_TAG_INFO, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Debugln(args ...interface{}) {
	if !l.checkLevel(util.LEVEL_DEBUG) {
		return
	}

	msg := fmt.Sprintln(args...)
	log.Printf("%s%s%s %s%s %s", COLOR_PURPLE, TOP_TAG_DEBUG, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s", COLOR_PURPLE, TOP_TAG_DEBUG, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Traceln(args ...interface{}) {
	if !l.checkLevel(util.LEVEL_TRACE) {
		return
	}

	msg := fmt.Sprintln(args...)
	log.Printf("%s%s%s %s%s %s", COLOR_RESET, TOP_TAG_TRACE, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s", COLOR_RESET, TOP_TAG_TRACE, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}

func (l *LoggerImplementation) Testln(args ...interface{}) {
	if !l.checkLevel(util.LEVEL_TEST) {
		return
	}

	msg := fmt.Sprintln(args...)
	log.Printf("%s%s%s %s%s %s", COLOR_GREEN, TOP_TAG_TEST, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)

	if !l.debugMode {
		l.logger.Printf("%s%s%s %s%s %s", COLOR_GREEN, TOP_TAG_TEST, COLOR_BLUE, l.getTags(), COLOR_RESET, msg)
	}
}
