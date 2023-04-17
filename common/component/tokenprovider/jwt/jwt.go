package jwt

import (
	"fmt"
	"food_delivery/common/component/tokenprovider"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{
		secret: secret,
	}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Duration(expiry) * time.Second).Unix(),
			IssuedAt:  now.Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return &claims.Payload, nil
}

func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}
