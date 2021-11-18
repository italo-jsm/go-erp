package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	var token crypto.Token
	token.IssuedAt = time.Now()
	token.Role = "admin"
	token.Username = credential.Username
	result, _ := json.Marshal(token)
}