package file

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// MD5 returns the MD5 hash of a specific file.
func MD5(f *os.File) (string, error) {
	m := md5.New()
	_, err := io.Copy(m, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}

// Sha1 returns the Sha1 hash of a specific file.
func Sha1(f *os.File) (string, error) {
	hash := sha1.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Sha256 returns the Sha256 hash of a specific file.
func Sha256(f *os.File) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// GetMD5ByPath returns the file's MD5 hash by providing a file path.
func GetMD5ByPath(p string) (string, error) {
	f, err := os.Open(p)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return MD5(f)
}

// GetSha1ByPath returns the file's Sha1 hash by providing a file path.
func GetSha1ByPath(p string) (string, error) {
	f, err := os.Open(p)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return Sha1(f)
}

// GetSha256ByPath returns the file's Sha256 hash by providing a file path.
func GetSha256ByPath(p string) (string, error) {
	f, err := os.Open(p)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return Sha256(f)
}
