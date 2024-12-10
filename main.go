package main

import (
	"go-web-native/config"
	"go-web-native/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	// 	Database connection
	config.ConnectDB()

	// Routes
	// Homepage
	http.HandleFunc("/", homecontroller.Home)

	// Handler untuk file statis
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Run server
	log.Println("Server running on port: 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
