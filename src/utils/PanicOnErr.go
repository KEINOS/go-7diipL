package utils

import (
	"fmt"
	"os"
)

// PanicOnErr は err がエラーの場合のみパニックを発生します.
//
// OsExit では都合が悪い時（defer を実行させる必要があるなど）に利用します.
func PanicOnErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(err)
	}
}
