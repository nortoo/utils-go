package file

import (
	"testing"
)

func TestExists(t *testing.T) {
}

func TestIsValidPath(t *testing.T) {
	samples := map[string]bool{
		"/a/a/c.txt":          true,
		"/b/c/a/g/./a/..":     false,
		"/../.../a/c/":        false,
		"./a/c//ccc/cd/b.txt": true,
	}
	for s, isValid := range samples {
		if IsValidPath(s) != isValid {
			t.Fatalf("IsValidPath does not cover: %s", s)
		}
	}
}

func TestIsDir(t *testing.T) {
}

func TestSize(t *testing.T) {

}
