package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken(account *CredentialParam) (AuthHeader, error) {
	var claims = jwt.MapClaims{
		"sub":      "BackendDev",
		"username": account.Username,
		"iss":      "MiniProject",
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return AuthHeader{}, err
	}

	var newAuthHeader = AuthHeader{Bearer: signedToken}

	return newAuthHeader, nil
}

func VerifyToken(auth *AuthHeader) (AuthHeader, error) {
	verifiedToken, err := jwt.Parse(auth.Bearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret-key"), nil
	})
	if err != nil {
		return AuthHeader{}, err
	}

	if !verifiedToken.Valid {
		return AuthHeader{}, errors.New("token not valid")
	}

	var claims, ok = verifiedToken.Claims.(jwt.MapClaims)
	if !ok {
		return AuthHeader{}, errors.New("failed parsing claim")
	}

	for key, value := range claims {
		switch key {
		case "sub":
			if value != "BackendDev" {
				return AuthHeader{}, errors.New("wrong claim sub")
			}
		case "iss":
			if value != "MiniProject" {
				return AuthHeader{}, errors.New("wrong claim iss")
			}
		case "exp":
			var tm time.Time
			switch exp := claims["exp"].(type) {
			case float64:
				tm = time.Unix(int64(exp), 0)
			case json.Number:
				v, _ := exp.Int64()
				tm = time.Unix(v, 0)
			}
			if tm.Before(time.Now()) {
				return AuthHeader{}, errors.New("expires")
			}

		}
	}
	return *auth, nil
}
