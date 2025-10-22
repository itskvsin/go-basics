package main

import (
	"fmt"
	"net/http"
)

func routeHandler(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(res, "This is the Home Page")
	case "/about":
		fmt.Fprintf(res, "This is the About Page")
	case "/contact":
		fmt.Fprintf(res, "This is the Contact Page")
	default:
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "404 Page is not found")
	}
}

func main() {
	http.HandleFunc("/", routeHandler)

	http.ListenAndServe(":3000", nil)
}
