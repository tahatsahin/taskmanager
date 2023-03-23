package common

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	absPriv, err := filepath.Abs(privKeyPath)
	absPriv = strings.Split(absPriv, "taskmanager")[0] + "taskmanager\\keys\\private.rsa"
	if err != nil {
		log.Fatalf("cannot get absolute path: %v", err)
	}

	absPub, err := filepath.Abs(pubKeyPath)
	absPub = strings.Split(absPub, "taskmanager")[0] + "taskmanager\\keys\\public.rsa"
	if err != nil {
		log.Fatalf("cannot get absolute path: %v", err)
	}

	signKey, err = os.ReadFile(absPriv)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = os.ReadFile(absPub)
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

// GenerateJWT generates a token which is signed with proper claims.
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

// TokenExtractor struct to extract token from request
type TokenExtractor struct {
}

// ExtractToken takes token from the request header and trims spaces/prefixes
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

// Authorize is the gateway for certain pages. Users cannot reach resources (/tasks) without authorization
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
