package Auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secretkeydataimpact")

type JWTClaim struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func generateJWT(id string, password string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	claims := &JWTClaim{
		ID:       id,
		Password: string(passwordHash),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("Unexpected signing method")
	}
	return []byte(jwtKey), nil
}

func validateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, keyFunc)
	if err != nil {
		return errors.New("Erreur token")
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func GenerateToken(c *fiber.Ctx, id string, password string) (string, error) {

	tokenString, err := generateJWT(id, password)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckToken(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()
	tokenString := headers["Authorization"]
	if tokenString == "" {
		return errors.New("Access Unauthorized")
	}

	err := validateToken(tokenString)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
