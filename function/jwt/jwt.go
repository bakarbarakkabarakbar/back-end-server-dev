package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go/types"
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
	switch account.RoleId {
	case 1:
		signedToken, err := token.SignedString([]byte("secret-key-super-admin"))
		if err != nil {
			return AuthHeader{}, err
		}

		var newAuthHeader = AuthHeader{Bearer: signedToken}

		return newAuthHeader, nil
	case 2:
		signedToken, err := token.SignedString([]byte("secret-key-admin"))
		if err != nil {
			return AuthHeader{}, err
		}

		var newAuthHeader = AuthHeader{Bearer: signedToken}

		return newAuthHeader, nil
	case 3:
		signedToken, err := token.SignedString([]byte("secret-key-customer"))
		if err != nil {
			return AuthHeader{}, err
		}

		var newAuthHeader = AuthHeader{Bearer: signedToken}

		return newAuthHeader, nil
	}
	return AuthHeader{}, errors.New("role id not defined")
}

func VerifySuperAdminToken(auth *AuthHeader) (AuthHeader, error) {
	var verifiedToken *jwt.Token
	var err error
	verifiedToken, err = jwt.Parse(auth.Bearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret-key-super-admin"), nil
	})

	if err != nil {
		return AuthHeader{}, errors.New("wrong secret key")
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

func VerifyAdminToken(auth *AuthHeader) (AuthHeader, error) {
	var verifiedToken *jwt.Token
	var err error
	verifiedToken, err = jwt.Parse(auth.Bearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret-key-admin"), nil
	})

	if err != nil {
		return AuthHeader{}, errors.New("wrong secret key")
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
			case types.Nil:
				continue
			}
			if tm.Before(time.Now()) {
				return AuthHeader{}, errors.New("expires")
			}

		}
	}
	return *auth, nil
}
