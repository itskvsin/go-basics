package main

import (
	"fmt"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request){
	name := req.URL.Query().Get("name")
	if name == ""{
		name = "Guest"
	}

	fmt.Fprintf(res,"Welcome %s!",name)
}

func main(){
	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":3000",nil)
}