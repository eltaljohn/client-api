package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const urlBase = "http://localhost:8080"

func main() {
	lc := loginClient(urlBase+"/v1/login", "contact@ed.team", "123456")
	// fmt.Println(lc)

	/*person := Person{
		Name:        "Maria",
		Age:         72,
		Communities: []Community{{"EDteam"}}}
	resp := createPerson(urlBase+"/v1/persons", lc.Data.Token, &person)
	fmt.Println(resp)*/

	p := getPersonByID(urlBase+"/v1/persons/1", lc.Data.Token)
	fmt.Println(p)

	/*person := Person{
		Name:        "Javier",
		Age:         22,
		Communities: []Community{{"EDteam"}}}
	resp := updatePerson(urlBase+"/v1/persons/1", lc.Data.Token, &person)
	fmt.Println(resp)*/

	res := deletePerson(urlBase+"/v1/persons/1", lc.Data.Token)
	fmt.Println(res)

}

func httpsClient(method, url, token string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("request: %v", err)
	}

	return res
}
