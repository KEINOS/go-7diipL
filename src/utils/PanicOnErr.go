package utils

import (
	"fmt"
	"os"
)

func PanicOnErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(err)
	}
}
