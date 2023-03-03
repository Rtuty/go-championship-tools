package cyphers

import (
	"errors"
	modular "github.com/TheAlgorithms/Go/math/modular"
)

func Encrypt(message []rune, publicExponent, modulus int64) ([]rune, error) {
	var encrypted []rune

	for _, letter := range message {
		encryptedLetter, err := modular.Exponentiation(int64(letter), publicExponent, modulus)
		if err != nil {
			return nil, errors.New("failed to Encrypt")
		}
		encrypted = append(encrypted, rune(encryptedLetter))
	}

	return encrypted, nil
}
