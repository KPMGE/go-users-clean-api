package providers

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct{}

func (hasher *BcryptHasher) Hash(plainText string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), 10)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}
