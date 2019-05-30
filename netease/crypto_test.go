package netease

import "testing"

func TestAESCBCEncryptAndDecrypt(t *testing.T) {
	secKey := createSecretKey(16, Base62)
	origData := `{"id":10086}`

	encData, err := aesCBCEncrypt([]byte(origData), secKey, []byte(IV))
	if err != nil {
		t.Fatal("encrypt failed")
	}

	decData, err := aesCBCDecrypt(encData, secKey, []byte(IV))
	if err != nil {
		t.Fatal("decrypt failed")
	}

	if plainText := string(decData); plainText != origData {
		t.Errorf("aesCBCDecrypt got: %s, want: %s", plainText, origData)
	}
}
