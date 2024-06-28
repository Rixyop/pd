package utils

import (
	"seen/internal/types"

	"github.com/matthewhartstonge/argon2"
)

var argon = argon2.DefaultConfig()

func HashPassword(password []byte) (string, *types.Error) {
	encoded, err := argon.HashEncoded(password)
	if err != nil {
		return "", types.NewInternalError("خطای داخلی رخ داده است. کد خطا 1")
	}
	return string(encoded), nil
}

func VerifyPassword(password []byte, encodedPassword []byte) (bool, *types.Error) {
	ok, err := argon2.VerifyEncoded(password, encodedPassword)
	if err != nil {
		return false, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 2")
	}
	return ok, nil
}
