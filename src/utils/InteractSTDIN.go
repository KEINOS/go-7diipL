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
func InteractSTDIN(funcUser func(string) error, stopWord string, prompt string) (err error) {
	setoff := PrintMsgWait("Please wait ...")

	if !IsTerminal() {
		return xerrors.New("対話モードエラー:ターミナル/コマンドラインからの実行ではありません.")
	}

	setoff()

	fmt.Printf("- ストップワード: %s\n", stopWord)

	scanner := bufio.NewScanner(os.Stdin)
	input := ""

	fmt.Print(prompt)

	for scanner.Scan() {
		input = scanner.Text()

		if input == stopWord {
			fmt.Printf("%s を検知しました。対話モードを終了します。\n", stopWord)

			break
		}

		// Run user function
		err = funcUser(input)

		if err != nil {
			return err
		}

		fmt.Print(prompt)
	}

	return nil
}
