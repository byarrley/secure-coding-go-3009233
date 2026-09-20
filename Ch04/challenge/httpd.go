package main

import (
	"fmt"
	"log"
	"net/http"
)

/* Task
- Write the middleware required to generate a Bearer token and restrict access to the admin endpoint
*/

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You're in!")
}

// requireAdmin is a middleware allowing only users with Admin role to access the handler
func requireAdmin(h http.Handler) http.Handler {
	// FIXME: Your code goes here
	return h
}

func main() {
	h := requireAdmin(http.HandlerFunc(adminHandler))
	http.Handle("/admin", h)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
