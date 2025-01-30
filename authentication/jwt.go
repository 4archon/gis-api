package authentication

import (
	"errors"
	"strconv"
	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	secretKey []byte
}

func (j *JwtToken) Init(key []byte) {
	j.secretKey = key
}

func (j *JwtToken) GetToken(id int, role string) (string, error) {
	sub := strconv.Itoa(id)
	payload := jwt.MapClaims{
		"sub": sub,
		"role": role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JwtToken) GetPayload(token string) (int, string, error) {
	claims := jwt.MapClaims{}
	tok, err := jwt.ParseWithClaims(token, claims, func(tok *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return 0, "", err
	}
	if !tok.Valid {
		return 0, "", errors.New("invalid token")
	}
	var id int
	var role string
	for key, value := range claims {
		if key == "sub" {
			sub := value.(string)
			id, err = strconv.Atoi(sub)
			if err != nil {
				return 0, "", err
			}
		} else if key == "role" {
			role = value.(string)
		}
	}
	return id, role, nil
}