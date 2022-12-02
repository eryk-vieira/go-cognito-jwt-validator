package validator

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

type Config struct {
	Region          string
	CognitoPoolId   string
	CognitoClientId string
}

func New(config *Config) *Config {
	return config
}

func (config *Config) Validate(jwtToken string) error {
	pKey, err := getPublicKeys(config.Region, config.CognitoPoolId)

	if err != nil {
		log.Fatal("Error trying to get Cognito public keys, check your config")
	}

	keySet, _ := jwk.Parse(pKey)

	parsedToken, err := jwt.Parse([]byte(jwtToken), jwt.WithKeySet(keySet))

	if err != nil {
		return errors.New("INVALID TOKEN")
	}

	clientId, _ := parsedToken.Get("aud")
	token_use, _ := parsedToken.Get("token_use")

	if clientId.([]string)[0] != config.CognitoClientId {
		return errors.New("TOKEN IS FROM A DIFFERENT client_id")
	}

	if parsedToken.Issuer() != fmt.Sprintf("https://cognito-idp.us-east-2.amazonaws.com/%s", config.CognitoPoolId) {
		return errors.New("TOKEN IS FROM A DIFFERENT pool_id")
	}

	if token_use != "id" && token_use != "access" {
		return errors.New("TOKEN IS FROM A DIFFERENT source")
	}

	if time.Now().After(parsedToken.Expiration()) {
		return errors.New("TOKEN EXPIRED")
	}

	return nil
}

func getPublicKeys(region string, cognitoPoolId string) ([]byte, error) {
	var url = fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, cognitoPoolId)

	resp, err :=
		http.Get(url)

	if err != nil {
		fmt.Println("Error fetching public keys")
		return nil, errors.New("Error")
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return body, nil
}
