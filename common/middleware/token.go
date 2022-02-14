package middleware

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const (
	TokenSigned = "xiaofengjing"
	ExpiresTime = 12
)

// UserClaims ...
type UserClaims struct {
	UserID   int64  // 用户ID
	UserName string // 用户名
	Phone    string // 手机号
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(userID int64, userName string, phone string) (string, error) {
	claims := UserClaims{
		UserID:   userID,
		UserName: userName,
		Phone:    phone,
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(ExpiresTime)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(TokenSigned)) // 签名密钥
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(bearerToken string) (*UserClaims, error) {
	token := strings.ReplaceAll(bearerToken, "Bearer ", "")
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenSigned), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
