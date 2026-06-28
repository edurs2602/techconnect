package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		secret:     []byte(secret),
		accessTTL:  15 * time.Minute,
		refreshTTL: 7 * 24 * time.Hour,
	}
}

func (j *JWTService) IssueAccess(userID string) (string, error) {
	return j.issue(userID, "access", j.accessTTL)
}

func (j *JWTService) IssueRefresh(userID string) (string, error) {
	return j.issue(userID, "refresh", j.refreshTTL)
}

func (j *JWTService) issue(userID, kind string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"kind": kind,
		"exp":  time.Now().Add(ttl).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.secret)
}

func (j *JWTService) Validate(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método inválido")
		}
		return j.secret, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("token inválido")
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["sub"].(string), nil
}
