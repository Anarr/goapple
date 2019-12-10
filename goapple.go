package goapple

import (
	"fmt"
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

//User store Apple user data after successfully login
type User struct {
}

//Login login with apple
func Login(code string, c *Config) (*User, error) {
	var u User
	secret, err := createJWTClientSecret(c)

	if err != nil {
		return nil, ClientSecretErr
	}

	c.PrivateKey = secret

	s, err := fetchLoginData(code, c)

	if err != nil {
		return nil, FetchLoginDataErr
	}

	fmt.Println("The secret", secret)
	fmt.Println("Apple data", s)

	return &u, err
}
