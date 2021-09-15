package utils

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/xerrors"
)

// InteractSTDIN は対話モードで標準入力を処理します.
//
// 標準入力から stopWord が入力されるまで、funcUser() に標準入力の値を入れた結果を表示し続けます.
// 単発の標準入力が欲しい場合は GetSTDIN() を利用してください.
func InteractSTDIN(funcUser func(string) error, stopWord string) (err error) {
	if !IsTerminal() {
		return xerrors.New("対話モードエラー:ターミナル/コマンドラインからの実行ではありません.")
	}

	fmt.Printf("- ストップワード: %s\n", stopWord)

	scanner := bufio.NewScanner(os.Stdin)
	input := ""

	for scanner.Scan() {
		input = scanner.Text()

		if input == stopWord {
			LogDebug("%s を検知しました。対話モードを終了します", stopWord)

			break
		}

		err = funcUser(input)

		if err != nil {
			return err
		}
	}

	return nil
}
