package goapple

import (
	"fmt"
)


type Config struct {
	ClientID string
	ClientSecret
	GrantType string
	PrivateKey string
}

type ClientSecret struct {
	ISS, AUD, SUB, KID string
	IAT, EXP int
}

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