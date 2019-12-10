package goapple

import "errors"

var ClientSecretErr = errors.New("Can not create ClientSecret with given private key")
var FetchLoginDataErr = errors.New("Can not fetch data from apple with given config")