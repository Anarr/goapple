package goapple

import "github.com/dgrijalva/jwt-go"


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
