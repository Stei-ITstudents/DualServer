package server1

import (
	"fmt"
	"net/http"
)

func Start() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server 1 is running")
	})
	return http.ListenAndServe(":8080", nil)
}											
