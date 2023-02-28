package common

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	privKeyPath = "keys/private.rsa"
	pubKeyPath  = "keys/public.rsa"
)

var (
	verifyKey, signKey []byte
	SignKey            *rsa.PrivateKey
	VerifyKey          *rsa.PublicKey
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

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKey)
	if err != nil {
		log.Fatalf("[VerifyKey]: %v", err)
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

type TokenExtractor struct {
}

func (t TokenExtractor) ExtractToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		token = r.FormValue("access_token")
	}
	if token == "" {
		return "", request.ErrNoTokenInRequest
	}
	// for the sake of simplicity, extra spaces and type name Bearer are stripped
	return strings.TrimSpace(strings.TrimPrefix(token, "Bearer")), nil
}

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// validate token
	token, err := request.ParseFromRequest(r, TokenExtractor{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("[signingToken]: %v", ok)
		}
		return VerifyKey, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, fmt.Sprintf("{\"error\": \"%s %s\"}", "JWT not valid,", err), http.StatusUnauthorized)
	} else {
		next(w, r)
	}
}
