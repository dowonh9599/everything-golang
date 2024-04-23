package utils

import (
	"fmt"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/configs/app_config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(app_config.SECRET_KEY))

	if err != nil {
		return "", err
	}

	return webToken, nil
}

func VerifyToken(token string) (*jwt.Token, error) {
	tokenJwt, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)

		if !isValid {
			return nil, fmt.Errorf("unexpected sigining method: %v", t.Header["alg"])
		}
		return []byte(app_config.SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return tokenJwt, nil

}

func DecodeToken(token string) (jwt.MapClaims, error) {
	tokenJwt, errVerify := VerifyToken(token)

	if errVerify != nil {
		return nil, errVerify
	}

	claims, isOK := tokenJwt.Claims.(jwt.MapClaims)

	if !isOK || !tokenJwt.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
