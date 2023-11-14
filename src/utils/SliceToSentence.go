package utils

import (
	"strings"
	"unicode"

	"github.com/neurosnap/sentences/english"
)

// SliceSentences は文字列を文の区切りごとにスライスにします.
func SliceSentences(inputText string) []string {
	if IsEnglish(inputText) {
		return SliceToSentenceEng(inputText)
	}

	return SliceToSentenceChore(inputText)
}

// SliceToSentenceChore は英語以外の文字列として文の区切りごとにスライスするのに使われます.
func SliceToSentenceChore(inputText string) []string {
	var (
		result     = []string{}
		buff       = []rune("")
		isEOS      = false
		wasDotPrev = false // 1 つ前の文字が "." の場合に true
	)

	// 英語以外の場合は、句点もしくは ". " or ".\n" を検知した場合に 1 文の区切りとしてスライス化
	for _, char := range inputText {
		switch {
		case (string(char) == "。"):
			isEOS = true
		case (unicode.IsSpace(char) && wasDotPrev):
			isEOS = true
		case (string(char) == "."):
			wasDotPrev = true
		default:
			wasDotPrev = false
		}

		if isEOS {
			buff = append(buff, char)
			result = append(result, strings.TrimSpace(string(buff)))

			// Reset
			buff = buff[:0]
			wasDotPrev = false
			isEOS = false

			continue
		}

		buff = append(buff, char)
	}

	// バッファの残りを追加
	if !isEOS {
		result = append(result, strings.TrimSpace(string(buff)))
	}

	return result
}

// SliceToSentenceEng は英語の文字列として文の区切りごとにスライスするのに使われます.
func SliceToSentenceEng(inputTextEn string) []string {
	result := []string{}

	tokenizer, err := english.NewSentenceTokenizer(nil)
	PanicOnErr(err)

	sentences := tokenizer.Tokenize(inputTextEn)
	for _, s := range sentences {
		// タブやスペースインデントの連続を削除（キャッシュ時の精度向上のため）
		result = append(result, strings.Join(strings.Fields(s.Text), " "))
	}

	return result
}
