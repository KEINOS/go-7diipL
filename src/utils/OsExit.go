package utils

import "os"

// OsExit is an alias of os.Exit to ease mock the exit status.
var OsExit = os.Exit
