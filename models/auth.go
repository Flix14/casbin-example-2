package models

import (
	"fmt"
	"time"

	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
)

var JWT_KEY = []byte("Mi2rmk3237FVb408yuIR3237FVb408yi2rIH9y7CAaIgm")

type Claims struct {
	Id   int64
	User *User `json:"payload"`
	jwt.StandardClaims
}

func getExpirationJWT() time.Time {
	// Expiration, tomorrow at 5:00am
	now := time.Now()
	yyyy, mm, dd := now.Date()
	expiration := time.Date(yyyy, mm, dd+1, 5, 0, 0, 0, now.Location())
	return expiration
}

func (user *User) GenerateTokenJWT() (tokenStr string, exp time.Time, err error) {
	exp = getExpirationJWT()
	claims := Claims{
		Id:   exp.Unix(),
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
			Issuer:    "test.mx",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(JWT_KEY)
	if err != nil {
		return
	}
	return
}

func ValidateTokenJWT(tokenStr string, obj string, act string) (claims Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Token malformed - %v", err.Error())
		}
		return JWT_KEY, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = fmt.Errorf("Token no valid - %v", err.Error())
		return
	}

	ok, err := EnforceAuth(claims.User.Role, obj, act)
	if err != nil {
		return
	}

	if !ok {
		err = fmt.Errorf("Unauthorized | No autorizado")
		return
	}

	return
}

func EnforceAuth(sub string, obj string, act string) (bool, error) {
	enforcer, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, err := enforcer.Enforce(sub, obj+"/", act)
	return ok, nil
}
