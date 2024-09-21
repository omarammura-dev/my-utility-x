package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	RoleUser      Role = "USER"
	RoleAdmin     Role = "ADMIN"
	RoleModerator Role = "MODERATOR"
)

func GenerateToken(email, username string, userId primitive.ObjectID, addedTime time.Duration, role Role) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
		"userId":   userId,
		"role":     role,
		"exp":      time.Now().Add(addedTime).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func VerifyToken(token string) (primitive.ObjectID, Role, error) {

	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return primitive.ObjectID{}, "", errors.New("could not parse token")
	}

	if !parseToken.Valid {
		return primitive.ObjectID{}, "", errors.New("invalid token")

	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)

	if !ok {
		return primitive.ObjectID{}, "", errors.New("invalid token claims")
	}

	userIdStr, ok := claims["userId"].(string)

	if !ok {
		return primitive.ObjectID{}, "", errors.New("userId claim could not be parsed")
	}
	userId, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		return primitive.ObjectID{}, "", errors.New("userId claim is not a valid ObjectID")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return primitive.ObjectID{}, "", errors.New("role claim could not be parsed")
	}

	return userId, Role(role), nil
}
