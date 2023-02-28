package common

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

const (
	privKeyPath = "keys/private.rsa"
	pubKeyPath  = "keys/public.rsa"
)

var (
	verifyKey, signKey []byte
	SignKey            *rsa.PrivateKey
)

// read the files
func initKeys() {
	var err error

	signKey, err = os.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = os.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signKey)
	if err != nil {
		log.Fatalf("[SignKey]: %v\n", err)
	}
}

func GenerateJWT(name, role string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["UserInfo"] = struct {
		Name string
		role string
	}{name, role}
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	tokenString, err := t.SignedString(SignKey)
	if err != nil {
		log.Fatalf("[signedString]: %v", err)
		return "", err
	}
	return tokenString, nil
}
