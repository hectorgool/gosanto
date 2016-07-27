package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/hectorgool/gosanto/common"
	"github.com/hectorgool/gosanto/routers"
)

//Entry point of the program
func main() {
	
	fmt.Printf("main:0\n")	
	// Calls startup logic
	common.StartUp()
	fmt.Printf("main:1\n")
	
	// Get the mux router object
	router := routers.InitRoutes()
	fmt.Printf("main:2\n")

	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)
	fmt.Printf("main:3\n")
	
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	fmt.Printf("main:4\n")
	log.Println("Listening...")
	server.ListenAndServe()
	
	fmt.Printf("main:5\n")

}
