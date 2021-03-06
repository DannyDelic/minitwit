package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"minitwit/internal/conf"
	"minitwit/internal/store"
	"strconv"
	"time"

	"github.com/cristalhq/jwt/v3"
)

var (
	jwtSigner   jwt.Signer
	jwtVerifier jwt.Verifier
)

func jwtSetup(conf conf.Config) {
	var err error
	key := []byte(conf.JwtSecret)

	jwtSigner, err = jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		log.Println("Error creating JWT signer")
	}

	jwtVerifier, err = jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		log.Println("Error creating JWT verifier")
	}
}

func generateJWT(account *store.Account) string {
	claims := &jwt.RegisteredClaims{
		ID:        fmt.Sprint(account.AccountID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	}
	builder := jwt.NewBuilder(jwtSigner)
	token, err := builder.Build(claims)
	if err != nil {
		log.Println("Error building JWT")
	}
	return token.String()
}

func verifyJWT(tokenStr string) (int, error) {
	token, err := jwt.Parse([]byte(tokenStr))
	if err != nil {
		log.Println("Error parsing JWT")
		return 0, err
	}

	if err := jwtVerifier.Verify(token.Payload(), token.Signature()); err != nil {
		log.Println("Error verifying token")
		return 0, err
	}

	var claims jwt.StandardClaims
	if err := json.Unmarshal(token.RawClaims(), &claims); err != nil {
		log.Println("Error unmarshalling JWT claims")
		return 0, err
	}

	if notExpired := claims.IsValidAt(time.Now()); !notExpired {
		return 0, errors.New("Token expired.")
	}

	id, err := strconv.Atoi(claims.ID)
	if err != nil {
		log.Println("Error converting claims ID to number")
		return 0, errors.New("ID in token is not valid")
	}
	return id, err
}
