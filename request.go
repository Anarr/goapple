package goapple

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//fetchLoginData handle user code and goaple config struct
//make http post request to handle data from apple
//return data and relevant error finally
func fetchLoginData(code string, c *Config) (string, error) {
	v := url.Values{}
	v.Set("client_id", c.ClientID)
	v.Set("client_secret", c.PrivateKey.(string))
	v.Set("code", code)
	v.Set("grant_type", c.GrantType)
	vs := v.Encode()

	req, err := http.NewRequest("POST", AppleLoginURL, bytes.NewBuffer([]byte(vs)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "curl")
	req.Header.Add("Accept", "application/json")

	if err != nil {
		return "", err
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return "", err
	}

	s, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(string(s))
	return string(s), nil
}
