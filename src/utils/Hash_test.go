package utils_test

import (
	"encoding/hex"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const helloQiita = "Hello Qiita!"

func TestHash_unknown_algo(t *testing.T) {
	_, _, err := utils.Hash("unknown", "dummy input")

	require.Error(t, err, "unknown algorithm should return an error")
}

func TestHash_Blake3_256(t *testing.T) {
	{
		input := ""
		expect := "af1349b9f5f9a1a6a0404dea36dcc9499bcb25c9adc112b7cc9a93cae41f3262"
		actual, actualByte, err := utils.Hash("blake3_256", input)

		require.NoError(t, err)
		assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
		assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
	}
	{
		input := helloQiita
		expect := "79557e896c63f83cf4953eb852c465c326ebfcbb88173a13d864a42827902caf"
		actual, actualByte, err := utils.Hash("blake3_256", input)

		require.NoError(t, err)
		assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
		assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
	}
}

func TestHash_Blake3_512(t *testing.T) {
	{
		input := ""
		expect := "af1349b9f5f9a1a6a0404dea36dcc9499bcb25c9adc112b7cc9a93cae41f3262" +
			"e00f03e7b69af26b7faaf09fcd333050338ddfe085b8cc869ca98b206c08243a"
		actual, actualByte, err := utils.Hash("blake3_512", input)

		require.NoError(t, err)
		assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
		assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
	}
	{
		input := helloQiita
		expect := "79557e896c63f83cf4953eb852c465c326ebfcbb88173a13d864a42827902caf" +
			"ba337cd60b1618a79a1969daf0aeb620f6ea92163b53fe5eb4f1db9ed98e2fd8"
		actual, actualByte, err := utils.Hash("blake3_512", input)

		require.NoError(t, err)
		assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
		assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
	}
}

func TestHash_Fnv1_32(t *testing.T) {
	input := helloQiita
	expect := "b04649f6"
	actual, actualByte, err := utils.Hash("fnv1_32", input)

	require.NoError(t, err)
	assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
	assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
}

func TestHash_Fnv1_64(t *testing.T) {
	input := helloQiita
	expect := "97bfaffd885daad6"
	actual, actualByte, err := utils.Hash("fnv1_64", input)

	require.NoError(t, err)
	assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
	assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
}

func TestHash_MD5(t *testing.T) {
	input := helloQiita
	expect := "7c414ef7535afff21e05a36b1cfc9000"
	actual, actualByte, err := utils.Hash("md5", input)

	require.NoError(t, err)
	assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
	assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
}

func TestHash_SHA2_256(t *testing.T) {
	input := helloQiita
	expect := "e863d36c24ada694fa77454b33e8f9b9545d372aae251e8779fc25df16943fed"
	actual, actualByte, err := utils.Hash("sha2_256", input)

	require.NoError(t, err)
	assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
	assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
}

func TestHash_SHA2_512(t *testing.T) {
	input := helloQiita
	expect := "cac9036c1dd3652fc550e99a4ec2b066d69d6a40a369bc85e3078960e6f26012" +
		"4138f5d0f4e9a6e047dfb833c9dd9b3376d02d49be37de26dd6234d4e79cc09e"
	actual, actualByte, err := utils.Hash("sha2_512", input)

	require.NoError(t, err)
	assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
	assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
}

func TestHash_SHA3_256(t *testing.T) {
	input := helloQiita
	expect := "cedaea2333478d77bc9ed3e33303f45585af19174a451ce93029fedcdd4d1ecb"
	actual, actualByte, err := utils.Hash("sha3_256", input)

	require.NoError(t, err)
	assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
	assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
}

func TestHash_SHA3_512(t *testing.T) {
	input := helloQiita
	expect := "9d1aaed079bac1946a5fbafb4285ec5bdcf0053a9e00559046884a148f28fcb9" +
		"02441b49cb3a82775834a2444dd183eda36cee900f5662f82353fae7e7111740"
	actual, actualByte, err := utils.Hash("sha3_512", input)

	require.NoError(t, err)
	assert.Equal(t, expect, actual, "The input '%s' did not return the expected value", input)
	assert.Equal(t, expect, hex.EncodeToString(actualByte), input)
}
