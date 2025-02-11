package crypto

import (
	"testing"
)

func TestAESCBC(t *testing.T) {
	key := "abcdefg123!@#njlabcdefg123!@#njl"
	inputContent := []string{
		"hahahha",
		"asmjdkoajs888@##",
		"google123!@##$$/\\",
		"",
		"0",
	}
	for _, i := range inputContent {
		o, err := AESCBCEncryptor([]byte(i), []byte(key))
		if err != nil {
			t.Fatal(err)
		}

		ori, err := AESCBCDecrypter(o, []byte(key))
		if err != nil {
			t.Fatal(err)
		}
		if string(ori) != i {
			t.Fatalf("AESCBCDecrypter got an unexpected result")
		}
	}
}

func TestAESGCM(t *testing.T) {
	key := "abcdefg123!@#njlabcdefg123!@#njl"
	inputContent := []string{
		"hahahha",
		"asmjdkoajs888@##",
		"google123!@##$$/\\",
		"",
		"0",
	}
	for _, i := range inputContent {
		o, err := AESGCMEncryptor([]byte(i), []byte(key))
		if err != nil {
			t.Fatal(err)
		}

		ori, err := AESGCMDecrypter(o, []byte(key))
		if err != nil {
			t.Fatal(err)
		}
		if string(ori) != i {
			t.Fatalf("AESGCMDecrypter got an unexpected result")
		}
	}
}
