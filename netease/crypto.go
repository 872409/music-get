package netease

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/winterssy/music-get/utils"
	"math/big"
	"math/rand"
	"time"
)

const (
	Base62                      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	PresetKey                   = "0CoJUm6Qyw8W8jud"
	IV                          = "0102030405060708"
	DefaultRSAPublicKeyModulus  = "e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	DefaultRSAPublicKeyExponent = 0x10001
)

func Encrypt(origData []byte) (params, encSecKey string, err error) {
	enc1, err := aesCBCEncrypt(origData, []byte(PresetKey), []byte(IV))
	if err != nil {
		return
	}

	secKey := createSecretKey(16, Base62)
	enc2, err := aesCBCEncrypt([]byte(enc1), secKey, []byte(IV))
	if err != nil {
		return
	}

	params, encSecKey = enc2, rsaEncrypt(utils.BytesReverse(secKey), DefaultRSAPublicKeyModulus, DefaultRSAPublicKeyExponent)
	return
}

func createSecretKey(size int, charset string) []byte {
	secKey, n := make([]byte, size), len(charset)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range secKey {
		secKey[i] = charset[r.Intn(n)]
	}
	return secKey
}

func aesCBCEncrypt(plainText, secKey, iv []byte) (string, error) {
	block, err := aes.NewCipher(secKey)
	if err != nil {
		return "", err
	}

	plainText = pkcs5Padding(plainText, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func aesCBCDecrypt(cipherText string, secKey, iv []byte) ([]byte, error) {
	cipherTextDec, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(secKey)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherTextDec))

	blockMode.CryptBlocks(plainText, cipherTextDec)
	plainText = pkcs5UnPadding(plainText)

	return plainText, nil
}

func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, paddingText...)
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:length-unPadding]
}

func rsaEncrypt(origData []byte, modulus string, exponent int64) string {
	bigOrigData := big.NewInt(0).SetBytes(origData)
	bigModulus, _ := big.NewInt(0).SetString(modulus, 16)
	bigRs := bigOrigData.Exp(bigOrigData, big.NewInt(exponent), bigModulus)
	return fmt.Sprintf("%0256x", bigRs)
}
