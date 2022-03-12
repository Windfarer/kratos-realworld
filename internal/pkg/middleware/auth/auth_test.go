package auth

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGenerateToken(t *testing.T) {
	tk := GenerateToken("secret", "eric")
	spew.Dump(tk)
	panic("tk")
}
