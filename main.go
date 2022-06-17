package main

import (
	"log"
	"net/http"

	"github.com/krish8learn/BasicURLShortener/urlHandler"
)

func main() {
	log.Println("Starting URLShortener, running server port: 8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/home", urlHandler.Home())

	// http.ListenAndServe(":8080", urlHandler.MapHandler(urlHandler.PathURLs, mux))
	yamlWay, err := urlHandler.JsonHandler(mux)
	// yamlWay, err := urlHandler.YamlHandler(mux)
	if err != nil {
		log.Fatalln("Error in Yaml/json", err)
	}

	http.ListenAndServe(":8080", yamlWay)
}
