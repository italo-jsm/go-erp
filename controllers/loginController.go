package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"github.com/italosm/go-erp/crypto"
)

type LoginController struct{

}

func (loginController *LoginController) AttemptAuthentication(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)

	var credential crypto.Credential

	if err != nil {
		panic(err.Error)
	}

	json.Unmarshal(reqBody, &credential)

	w.Write([]byte(generateToken(credential)))
}

func generateToken(credential crypto.Credential) string{

	var header crypto.Header
	header.Alg = "HS512"
	header.Type = "JWT"

	var token crypto.Token
	token.IssuedAt = time.Now()
	token.Role = "admin"
	token.Username = credential.Username

	headerJson, _ := json.Marshal(header)

	tokenJson, _ := json.Marshal(token)

	headerPlusTokenString := strings.Join([]string{base64.StdEncoding.EncodeToString(headerJson), base64.StdEncoding.EncodeToString(tokenJson)}, ".")

	signature := base64.StdEncoding.EncodeToString(sha512.New().Sum([]byte(strings.Join([]string{headerPlusTokenString, "mykey"}, ""))))
	return strings.Join([]string{headerPlusTokenString, signature}, ".")
}