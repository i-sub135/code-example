package addons

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func secret() string {
	return os.Getenv("SECRET_KEY")
}

// CreateJwt -- create jwt
func (r *RestAddons) CreateJwt(phone string) (string, error) {
	keys := []byte(secret())

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// Set some claims

	claims["phone"] = phone
	claims["exp"] = time.Now().Add(time.Hour * 24).Local()
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(keys)
	return tokenString, err
}

// ValidJwt -- validasi
func (r *RestAddons) ValidJwt(MyToken string) (jwt.MapClaims, error) {
	keys := secret()
	token, err := jwt.Parse(MyToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(keys), nil
	})
	if token == nil {
		return nil, errors.New("not format token")

	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err

}
