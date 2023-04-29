package auth

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodetoken string) (*jwt.Token, error)
}

type jwtService struct {
}

var signKey = []byte("kdnjsndjnd*jdnj212md")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = json.Number(strconv.FormatInt(time.Now().Add(time.Minute*15).Unix(), 10))
	claim["aud"] = "Audience" // OPTIONAL AUDIENCE

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodetoken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodetoken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(signKey), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
