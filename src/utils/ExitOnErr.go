package utils

import (
	"fmt"
	"os"
)

func ExitOnErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		OsExit(FAILURE)
	}
}
