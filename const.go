package loggergo

import "os"

const (
	COLOR_RED    = "\033[31m"
	COLOR_YELLOW = "\033[33m"
	COLOR_BLUE   = "\033[36m"
	COLOR_GREEN  = "\033[32m"
	COLOR_PURPLE = "\033[35m"
	COLOR_RESET  = "\033[0m"

	DEFAULT_FLAG  = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	DEFAULT_PERM  = os.FileMode(0644)
	DEFAULT_COLOR = false
)
