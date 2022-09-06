package main 

import "fmt"

// Following are the dev variables 

var username = "dev_user"
var api_address = "http://some-random-dev-website-that-does-not-exist.dev"
var api_key = "letmeintothedevenvironment"


func main() {
	fmt.Print("User ", username, " is making an API call to the address: ", api_address, " and specifies ", api_key, " as API key.")
}
