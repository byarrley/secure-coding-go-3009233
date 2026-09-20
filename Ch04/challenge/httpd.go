package main

import (
	"context"
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
	/*
		1. Request the user
		2. If the user is invalid, return status 'forbidden' (http403)
		3. if user does not have the admin role, return status 'unauthorized' (http401)
		4. Otherwise, return the handler
	*/

	//Adapted from Ch04_02's "requireAuth" function
	fn := func(w http.ResponseWriter, r *http.Request) {
		u, ok := RequestUser(r)

		if !ok {
			http.Error(w, "unauthenticated user", http.StatusForbidden)
			return
		}

		ok = u.HasRole(Admin)

		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", u)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func main() {
	h := requireAdmin(http.HandlerFunc(adminHandler))
	http.Handle("/admin", h)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
