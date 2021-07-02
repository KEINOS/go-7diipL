package utils

import (
	"io/ioutil"
	"os"
)

// ValSTDINDummy は GetSTDIN() の挙動を mock するための値です.
//
// この値が空でない場合、GetSTDIN() はセットされた値を返します.
var ValSTDINDummy string = ""

// GetSTDIN はパイプ渡し or リダイレクトされた標準入力からのデータを返します.
//
// 対話式で標準入力を取得したい場合は InteractSTDIN() を利用してください.
func GetSTDIN() (stdin string, err error) {
	if ValSTDINDummy != "" {
		return ValSTDINDummy, nil
	}

	value, err := ioutil.ReadAll(os.Stdin)

	return string(value), err
}
