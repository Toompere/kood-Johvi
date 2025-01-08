package main

import (
	"fmt"
	"net/http"
)

func main() {
	GetData()
	go Geocode()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", IndexHandler)

	// Print link to site
	fmt.Println("Open http://localhost:8080")
	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println(http.ListenAndServe(":8080", nil))

}
