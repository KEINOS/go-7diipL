<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# deepleng

```go
import "github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
```

Package deepleng は DeepL の翻訳 API を使った翻訳エンジンです.

このエンジンを利用するには、DeepL PRO の無料／有料アカウントから発行したアクセス・トークン（認証キー/API キー）が必要です。

\- アカウント登録後のアクセス・トークン確認先: https://www.deepl.com/pro-account/plan （認証キー）

## Index

- [Constants](<#constants>)
- [func GetInfoAPI\(prop \*engine.Properties\) \(engine.AccountInfo, error\)](<#GetInfoAPI>)
- [func GetURLBaseAPI\(p \*engine.Properties\) string](<#GetURLBaseAPI>)
- [func IsAPIKeyFree\(apikey string\) bool](<#IsAPIKeyFree>)
- [func New\(cacheID ...string\) \*engine.Properties](<#New>)
- [func Translate\(prop \*engine.Properties, inText string, langFrom string, langTo string\) \(result string, err error\)](<#Translate>)


## Constants

<a name="NameVarEnvAPIKey"></a>NameVarEnvAPIKey は環境変数の変数名で、DeepL の認証キー（アクセストークン）用の変数名です.

```go
const NameVarEnvAPIKey = "DEEPL_API_KEY"
```

<a name="GetInfoAPI"></a>
## func [GetInfoAPI](<https://github.com/Qithub-BOT/QiiTrans/blob/main/src/engines/deepleng/GetInfoAPI.go#L13>)

```go
func GetInfoAPI(prop *engine.Properties) (engine.AccountInfo, error)
```

GetInfoAPI はアクセス・トークンのアカウント情報のうち、engine.AccountInfo に必要なものだけセットして返します.

<a name="GetURLBaseAPI"></a>
## func [GetURLBaseAPI](<https://github.com/Qithub-BOT/QiiTrans/blob/main/src/engines/deepleng/GetURLBaseAPI.go#L7>)

```go
func GetURLBaseAPI(p *engine.Properties) string
```

GetURLBaseAPI は DeepL の翻訳 API のエンドポイントのドメインを返します. 無料・有料アカウントでドメインは異なるため、アクセス・トークン（認証キー）から自動検知して返します.

<a name="IsAPIKeyFree"></a>
## func [IsAPIKeyFree](<https://github.com/Qithub-BOT/QiiTrans/blob/main/src/engines/deepleng/IsAPIKeyFree.go#L6>)

```go
func IsAPIKeyFree(apikey string) bool
```

IsAPIKeyFree は apikey が無料枠のアクセストークンか返します.

<a name="New"></a>
## func [New](<https://github.com/Qithub-BOT/QiiTrans/blob/main/src/engines/deepleng/New.go#L11>)

```go
func New(cacheID ...string) *engine.Properties
```

New は翻訳エンジン（DeepL）の新規インスタンスのポインタを返します.

<a name="Translate"></a>
## func [Translate](<https://github.com/Qithub-BOT/QiiTrans/blob/main/src/engines/deepleng/Translate.go#L18>)

```go
func Translate(prop *engine.Properties, inText string, langFrom string, langTo string) (result string, err error)
```

Translate は inText を langFrom から langTo に DeepL の翻訳 API を使って翻訳した結果を返します.

この関数は結果をキャッシュしません。別途、呼び出し元でキャッシュを行ってください.

------

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
