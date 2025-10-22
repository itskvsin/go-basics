package main

import (
	"fmt"
	"net/http"
)

func userHandler(res http.ResponseWriter, req *http.Request){
	// path := req.URL.Path

	// parts := len("/user/")
	// if len(path) <= parts {
	// 	res.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(res,"user id missing")
	// 	return
	// }
	// id := path[parts:]

	id := req.PathValue("userId")
	fmt.Fprintf(res,"User ID: %s", id)
}

func main(){
	http.HandleFunc("/user/{userId}", userHandler)
	http.ListenAndServe(":3000", nil)
}