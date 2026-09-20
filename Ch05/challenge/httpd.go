package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
)

/* Task: protect everything under "/debug" with basic auth
- can use 'isValidAuth'
- '/debug/vars' will be available
- Hint: look at the second parameter of pkg.go.dev/net/http#ListenAndServe
*/

var (
	numCalls = expvar.NewInt("messages.calls")
)

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	numCalls.Add(1)

	// TODO:
	fmt.Fprintf(w, "TBD\n")
}

func main() {
	http.HandleFunc("/messages", messagesHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
