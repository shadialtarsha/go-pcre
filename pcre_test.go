package pcre_test

import (
	"testing"

	"github.com/shadialtarsha/go-pcre"
)

func TestCompileError(t *testing.T) {
	r, err := pcre.Compile("(")
	if err == nil {
		t.Error("expected error, got nil")
	}
	defer r.Close()
}
