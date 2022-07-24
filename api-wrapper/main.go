// Making calls to the Poke API (https://pokeapi.co/?ref=public-apis)
// and returning the result. However, this wrapper only works for 
// pokemons, i.e., the calls to http://pokeapi.co/api/v2/pokemon/ditto

package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
)

// returnJsonResponse: return the JSON-formatted response via HTTP
func returnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(resMessage)
}

func startingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! If you can read this, things worked out. \nWelcome to the starting page!")
}

func apicall(w http.ResponseWriter, r *http.Request) {
	// Get the pokemon's name from the HTTP request's query
	name := string(r.URL.Query()["name"][0])
	// Append that name to the base query
	// Please note that the name is passed as a query element
	// not as a header
	url := "http://pokeapi.co/api/v2/pokemon/"
	url += name
	// Make the HTTP request to the url
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// Parse the response data and return the outcome
	responseData, err := ioutil.ReadAll(response.Body)
	fmt.Println(responseData)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(string(responseData))
}

func main() {
	// Define the functions, i.e., the corresponding websites
	http.HandleFunc("/", startingPage)
	http.HandleFunc("/poke", apicall)
	// Start the server and let it run
	log.Fatal(http.ListenAndServe(":1000", nil))
}