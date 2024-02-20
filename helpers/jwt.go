package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	GrouId     uint   `json:"group_id"`
	IsAdmin    bool   `json:"is_admin"`
	IsInvestor bool   `json:"is_investor"`
	IsStudent  bool   `json:"is_student"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, groupId uint, isAdmin bool, isInvestor bool, isStudent bool) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Id:         id,
		Username:   username,
		GrouId:     groupId,
		IsAdmin:    isAdmin,
		IsInvestor: isInvestor,
		IsStudent:  isStudent,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
