package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func createPerson(url, token string, person *Person) GeneralResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("error at marshal of person: %v", err)
	}

	resp := httpsClient(http.MethodPost, url, token, data)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading body at createPerson: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("spected code 201, but got: %d", resp.StatusCode)
	}

	dataResponse := GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("error at unmarshal of body at create person")
	}

	return dataResponse
}

func getPersonByID(url, token string) PersonResponse {
	resp := httpsClient(http.MethodGet, url, token, nil)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading body at createPerson: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("spected code 200, but got: %d, body: %v", resp.StatusCode, string(body))
	}

	dataResponse := PersonResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("error at unmarshal of body at create person")
	}

	return dataResponse
}

func updatePerson(url, token string, person *Person) GeneralResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("error at marshal of person: %v", err)
	}

	resp := httpsClient(http.MethodPut, url, token, data)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading body at createPerson: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("spected code 201, but got: %d", resp.StatusCode)
	}

	dataResponse := GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("error at unmarshal of body at create person")
	}

	return dataResponse
}

func deletePerson(url, token string) GeneralResponse {
	resp := httpsClient(http.MethodDelete, url, token, nil)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading body at createPerson: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("spected code 200, but got: %d, body: %v", resp.StatusCode, string(body))
	}

	dataResponse := GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("error at unmarshal of body at create person")
	}

	return dataResponse
}
