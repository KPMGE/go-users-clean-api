package providers

import "github.com/golang-jwt/jwt"

type JwtTokenProvider struct {
}

func (j *JwtTokenProvider) Generate(data any) (string, error) {
	secretKey := []byte("super secret key")
	token := jwt.New(jwt.SigningMethodEdDSA)
}
