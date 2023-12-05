package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

type JWT struct {
	key []byte
}

func NewJWT(key []byte) *JWT {
	return &JWT{key: key}
}

func (j JWT) CreateToken(ttl time.Duration, payload any) (string, error) {
	now := time.Now().UTC()

	claims := jwt.MapClaims{
		"exp": now.Add(ttl).Unix(),
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"iss": "favaa-mitra",
		"sub": payload,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.key)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create token: ")
	}

	return token, nil
}

func (j JWT) CreateRefreshToken(ttl time.Duration) (string, error) {
	now := time.Now().UTC()

	claims := jwt.MapClaims{
		"exp": now.Add(ttl).Unix(),
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"iss": "favaa-mitra",
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.key)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create token: ")
	}

	return token, nil
}

func (j JWT) ValidateToken(token string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}

		return j.key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, errors.New("invalid token")
	}

	return claims["sub"], nil
}
