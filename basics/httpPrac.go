package main

import (
	"fmt"
	"net/http"
)

func homeLander(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "This is the home page")
}

func aboutLander(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "This ios the about page")
}

func main() {
	http.HandleFunc("/", homeLander)
	http.HandleFunc("/about", aboutLander)


	fmt.Println("Server started on port 3000")
	http.ListenAndServe(":3000", nil)
}
