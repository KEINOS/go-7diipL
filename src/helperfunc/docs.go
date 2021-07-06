/*
Package helperfunc はテストをしやすくするためのヘルパー関数集です. テストでのみ利用することを前提としています.

例えば、標準入力やコマンドの引数をモックしてテスト中に仮想的なユーザ入力が簡単に行えます.

注意点

いずれのヘルパー関数も、処理後に関数を返します. 必ず呼び出し直後に defer で、その戻り関数を設定してください.
この、戻りの関数は、内部でモック前のバックアップを行い、処理後にリカバリする（元に戻す）ための関数です.

基本的な使い方

	// 以下は foo bar が標準入力から渡されたのと同等な状態になります
	userInput := "foo bar"
	funcDefer := helperfunc.MockSTDIN(t, userInput)
	defer funcDefer()

より具体的な例

	// 以下のテストは `echo "Hello, world." | qiitrans --debug ja en` と同等の状態.
	func TestMockSTDIN(t *testing.T) {
		dummySTDIN := "Hello, world."
        dummyArgs := []string{"--debug", "ja", "en"}

		// 標準入力のモックとリカバリ準備
		funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
		defer funcDeferSTDIN() // モックのリカバリ

		// フラグ・オプションのモックとリカバリ準備
		funcDeferArgs := helperfunc.MockArgs(t, dummyArgs)
		defer funcDeferArgs() // モックのリカバリ

		// 標準入力のテスト
		{
			value, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				t.Fatalf("failed to read stdin during test")
			}

			expect := userInput
			actual := string(value)

			assert.Equal(t, expect, actual)
		}

		// 引数のテスト
		{
			expect := dummyArgs
			actual := os.Args[1:] // Args[0] は実行ファイルのパスなので、それ以降を取得

			assert.Equal(t, expect, actual)
		}
	}
*/
package helperfunc
