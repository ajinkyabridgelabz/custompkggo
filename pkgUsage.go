package main

import (
	"fmt"
	"net/http"

	"example.com/custompackage/utils" // User created custom package
	"github.com/gorilla/mux"          // Custom packge from internet
)

func main() {
	fmt.Println("The main package")

	// Using custom package function
	var result = utils.SomeCustomFuncOfAddition(20, 30)
	fmt.Println(result)

	r := mux.NewRouter()
	http.ListenAndServe("127.0.0.1:8900", r)

}
