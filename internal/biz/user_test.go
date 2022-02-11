package biz

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	s := hashPassword("abc")
	spew.Dump(s)
}

func TestVerifyPassword(t *testing.T) {
	a := assert.New(t)

	a.True(verifyPassword("$2a$10$9glSiir1PnzNgZzM1hXXaemPWECsyoYfWKbcr62APlXDvjl4yeTaK", "abc"))
	a.False(verifyPassword("$2a$10$9glSiir1PnzNgZzM1hXXaemPWECsyoYfWKbcr62APlXDvjl4yeTaK", "abc1"))
	a.False(verifyPassword("$2a$10$ssssd", "abc"))
}
