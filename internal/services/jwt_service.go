package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userID uint64) (string, error)
	ValidateToken(token string) (uint64, error)
}

type jwtServiceImpl struct {
	secretKey string
	issuer    string
}

// NewJWTService creates a new instance of JWTService
func NewJWTService(secretKey, issuer string) JWTService {
	return &jwtServiceImpl{
		secretKey: secretKey,
		issuer:    issuer,
	}
}

// Claims struct to hold the JWT claims
type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT for the given user ID
func (j *jwtServiceImpl) GenerateToken(userID uint64) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *jwtServiceImpl) ValidateToken(tokenString string) (uint64, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if time.Now().After(claims.ExpiresAt.Time) {
			return 0, errors.New("token has expired")
		}
		return claims.UserID, nil
	}

	return 0, errors.New("invalid token")
}
