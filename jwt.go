package goapple

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)


//createJWTClientSecret create apple client secret with given private key
func createJWTClientSecret(c *Config) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": c.ISS,
		"aud": c.AUD,
		"sub": c.SUB,
		"iat": c.IAT,
		"exp": c.EXP,
	})

	t.Header["kid"] =c.KID
	t.Header["alg"] = ALG

	s, err := t.SignedString(c.PrivateKey)

	if err != nil {
		return "", err
	}

	return s, nil
}


//parseDataFormBase64 decode base64 encoded string and get data
func parseDataFormBase64(s string) (string, error) {

	if !strings.Contains(s, ".") {
		return "", DecodeErr
	}

	xs := strings.Split(s, ".")

	if len(xs) >= 1 {
		dstr := xs[1]

		fmt.Println(dstr)

		b, err :=jwt.DecodeSegment(dstr)

		if err != nil {
			return "", err
		}

		return string(b), nil
	}

	return "", DecodeErr
}