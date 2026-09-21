package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
	"strings"
)

/* Task: protect everything under "/debug" with basic auth
- can use 'isValidAuth'
- '/debug/vars' will be available
- Hint: look at the second parameter of pkg.go.dev/net/http#ListenAndServe

Working notes:
- a 'Handle' is not the same as a 'Handler' (hence the function renames)...
- More like the 'Handle' functions register 'Handlers', so the original names were correct to refer to 'Handler"
- Had to do a fair bit of research on middleware and handlers; https://drstearns.github.io/tutorials/gomiddleware/ made what I had to do more clear
- Still not sure that this is the best way, but it seems like expvar's handler is automagically registered with the defaultServeMux
	(there's probably a way to change it but I didn't find it before coming up with this)
*/

var (
	numCalls = expvar.NewInt("messages.calls")
)

type EnsureAuth struct {
	handler http.Handler
}

func (ea *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Require auth for protected routes
	if strings.Contains(r.RequestURI, "/debug") {
		u, p, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "authentication failure", http.StatusForbidden)
			return
		}

		if !isValidAuth(u, p) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}

	//ServeHTTP called from wrapped handler
	ea.handler.ServeHTTP(w, r)
}

func NewEnsureAuth(handlerToWrap http.Handler) *EnsureAuth {
	return &EnsureAuth{handlerToWrap}
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	numCalls.Add(1)

	// TODO:
	fmt.Fprintf(w, "TBD\n")
}

func main() {
	http.HandleFunc("/messages", messagesHandler)

	ea := NewEnsureAuth(http.DefaultServeMux)
	if err := http.ListenAndServe(":8080", ea); err != nil {
		log.Fatal(err)
	}
}
