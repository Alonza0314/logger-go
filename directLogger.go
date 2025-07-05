package loggergo

import "log"

func Info(tag, msg string) {
	log.Printf("%s[INFO]%s [%s]%s %s\n", COLOR_BLUE, COLOR_BLUE, tag, COLOR_RESET, msg)
}

func Error(tag, msg string) {
	log.Printf("%s[EROR]%s [%s]%s %s\n", COLOR_RED, COLOR_BLUE, tag, COLOR_RESET, msg)
}

func Warn(tag, msg string) {
	log.Printf("%s[WARN]%s [%s]%s %s\n", COLOR_YELLOW, COLOR_BLUE, tag, COLOR_RESET, msg)
}

func Test(tag, msg string) {
	log.Printf("%s[TEST]%s [%s]%s %s\n", COLOR_GREEN, COLOR_BLUE, tag, COLOR_RESET, msg)
}

func Debug(tag, msg string) {
	log.Printf("%s[DBUG]%s [%s]%s %s\n", COLOR_PURPLE, COLOR_BLUE, tag, COLOR_RESET, msg)
}
