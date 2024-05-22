package util

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMain(m *testing.M) {
	m.Run()
}

func Test_HashAndCheck(t *testing.T) {
	password := "mySuperSecretPassword"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		panic(err)
	}

	match := CheckPasswordHash(password, hashedPassword)

	assert.Equal(t, true, match)
}
