// Running under go1.20.7
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Rest API")
	fmt.Println("========")

	respData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(respData))
	// That's all for the time being
}
