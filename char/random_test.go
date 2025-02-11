package char

import "testing"

func TestGenNonce(t *testing.T) {
	for i := 0; i < 10; i++ {
		n := GenNonce(15, 88)
		if n < 15 || n > 88 {
			t.Fatalf("the nonce [%d] is over the given range.", n)
		}
		t.Logf("get a new nonce: %d\n", n)
	}
}

func TestRandomBytes(t *testing.T) {
	for i := 0; i < 10; i++ {
		b := RandomBytes(64)
		if len(b) != 64 {
			t.Fatalf("the length of the string [%s] is over the given range.", string(b))
		}
		t.Logf("get a new random string: %s\n", string(b))
	}
}
