package pkg

import (
	"fmt"
	"seen/internal/models"
	"seen/internal/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(sid models.Token, expireTime int32) (string, *types.Error)
	VerifyToken(token string) (*jwtCustomClaim, *types.Error)
}

type jwtCustomClaim struct {
	SID models.Token `json:"sid"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "seen",
		secretKey: "f69f0580a78ff1aa841cdde6780f76e22164694c6f64ef11dccc396ff158baa8",
	}
}

func (s *jwtService) GenerateToken(sid models.Token, expireTime int32) (string, *types.Error) {
	claims := &jwtCustomClaim{
		SID: sid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(expireTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", types.NewInternalError("خطای داخلی رخ داده است. کد خطا 4")
	}
	return t, nil
}

func (s *jwtService) VerifyToken(token string) (*jwtCustomClaim, *types.Error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &jwtCustomClaim{}, keyFunc)
	if err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 5")
	}

	payload, ok := jwtToken.Claims.(*jwtCustomClaim)
	if !ok {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 6")
	}

	return payload, nil
}
