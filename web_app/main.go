package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["names"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		message := map[string]string{"name": name}
		enc := json.NewEncoder(rw)
		enc.Encode(message)
	})
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// go build . -> this created executable file based on os => Execute it and reach the port for accessing
