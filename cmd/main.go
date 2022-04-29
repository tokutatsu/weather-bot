package main

import (
	"net/http"

	function "github.com/tokutatsu/weather-bot"
)

func main() {
	http.HandleFunc("/", function.EntryPoint)
	http.ListenAndServe(":8080", nil)
}
