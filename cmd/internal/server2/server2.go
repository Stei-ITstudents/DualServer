package server2

import (
	"fmt"
	"net/http"
)

func Start() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server 2 is running")
	})
	return http.ListenAndServe(":8081", nil)
}
