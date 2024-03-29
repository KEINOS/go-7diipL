<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# helperfunc

```go
import "github.com/Qithub-BOT/QiiTrans/src/helperfunc"
```

Package helperfunc はテストをしやすくするためのヘルパー関数集です. テストでのみ利用することを前提としています.

例えば、標準入力やコマンドの引数をモックしてテスト中に仮想的なユーザ入力が簡単に行えます.

注意点

いずれのヘルパー関数も、処理後に関数を返します. 必ず呼び出し直後に defer で、その戻り関数を設定してください. この、戻りの関数は、内部でモック前のバックアップを行い、処理後にリカバリする（元に戻す）ための関数です.

基本的な使い方

```
// 以下は foo bar が標準入力から渡されたのと同等な状態になります
userInput := "foo bar"
funcDefer := helperfunc.MockSTDIN(t, userInput)
defer funcDefer()
```

より具体的な例

```
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
			value, err := io.ReadAll(os.Stdin)
			if err != nil {
				t.Fatalf("failed to read stdin during test")
			}

			expect := userInput
			actual := string(value)

			assert.Equal(t, expect, actual)
		}
```

## Index

- [Variables](<#variables>)
- [func FatalOnErr\(t \*testing.T, err error, comment ...string\)](<#FatalOnErr>)
- [func MockArgs\(t \*testing.T, argsDummy \[\]string\) func\(\)](<#MockArgs>)
- [func MockSTDIN\(t \*testing.T, inputDummy string\) func\(\)](<#MockSTDIN>)


## Variables

<a name="IsErrorDummy"></a>

```go
var (
    // IsErrorDummy が true の場合、FatalOnErr() は t.FailNow せずに標準エラー出力します.
    IsErrorDummy bool

    failNow func(format string, args ...interface{})
)
```

<a name="FatalOnErr"></a>
## func [FatalOnErr](<https://github.com/Qithub-BOT/QiiTrans/blob/main/src/helperfunc/FatalOnErr.go#L21>)

```go
func FatalOnErr(t *testing.T, err error, comment ...string)
```

FatalOnErr は err が nil ではない場合に t.FailNow でテストを終了します.

基本的に MockArgs と MockSTDIN などのヘルパー関数自身のカバレッジ目的で使用します. 一般的なテストでは t.Fatalf を利用してください.

<a name="MockArgs"></a>
## func [MockArgs](<https://github.com/Qithub-BOT/QiiTrans/blob/main/src/helperfunc/MockArgs.go#L15>)

```go
func MockArgs(t *testing.T, argsDummy []string) func()
```

MockArgs は、テストでユーザのコマンド引数（オプションやフラグ含む入力）をモックするヘルパー関数です. この関数は、モックの変更をリカバリーする defer 用の関数を返します.

```
dummyArgs := []string{"--debug", "ja", "en"}
funcDefer := helperfunc.MockArgd(t, dummyArgs)
defer funcDefer()
/* DO SOMETHING WITH ARGS HERE */
```

<a name="MockSTDIN"></a>
## func [MockSTDIN](<https://github.com/Qithub-BOT/QiiTrans/blob/main/src/helperfunc/MockSTDIN.go#L16>)

```go
func MockSTDIN(t *testing.T, inputDummy string) func()
```

MockSTDIN は、テストでユーザからの標準入力をモックするヘルパー関数です. この関数は、モックの変更をリカバリーする defer 用の関数を返します.

```
// 以下は echo "foo bar" | qiitrans と同等の状態
dummySTDIN := "foo bar"
funcDefer := helperfunc.MockSTDIN(t, dummySTDIN)
defer funcDefer()
/* DO SOMETHING WITH STDIN HERE */
```

------

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
