package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func loginClient(url, email, password string) LoginResponse {
	login := Login{email, password}

	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error en marshal de login: %v", err)
	}
	resp := httpsClient(http.MethodPost, url, "", data)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("espected 200, but got: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading body at login: %v", err)
	}

	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("error unmarshal on body at login: %v", err)
	}

	return dataResponse
}
