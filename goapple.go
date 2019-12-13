package goapple

import (
	"encoding/json"
)

//Config store credentials for login process
type Config struct {
	ClientID string
	ClientSecret
	GrantType  string
	PrivateKey interface{}
}

//ClienSecret store depenedencies especially uses for create jwt client secrets
type ClientSecret struct {
	ISS, AUD, SUB, KID string
	IAT, EXP           int
}

type AppleLoginResponse struct {
	IDToken string `json:"id_token"`
}

//UserPaylod store AppleAppleLoginResponse parsed data
type UserPayload struct {
	Email string `json:"email"`
	EmailVerified bool `json:"email_verified"`
	IsPrivateEmail bool `json:"is_private_email"`
	SUB string `json:"sub"`
}

//Login login with apple
func Login(code string, c *Config) (*UserPayload, error) {
	var u UserPayload
	var appleLoginRes AppleLoginResponse

	secret, err := createJWTClientSecret(c)

	if err != nil {
		return nil, ClientSecretErr
	}

	c.PrivateKey = secret

	s, err := fetchLoginData(code, c)

	if err != nil {
		return nil, FetchLoginDataErr
	}
	//decode first step data and get id_token
	err = json.Unmarshal([]byte(s), &appleLoginRes)
	if err != nil || appleLoginRes.IDToken == "" {
		return nil, err
	}

	//parse id_token and get user apple data
	dataStr, err := parseDataFormBase64(appleLoginRes.IDToken)

	if err != nil {
		return nil, err
	}

	//decode user apple data
	err = json.Unmarshal([]byte(dataStr), &u)

	if u.Email == "" || err != nil {
		return nil, err
	}

	return &u, err
}