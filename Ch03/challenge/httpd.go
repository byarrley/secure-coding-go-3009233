package main

import (
	"fmt"
	"html/template" //Ensure that html/template is imported, *not* "text/template"
	"log"
	"net/http"
)

/*
Task: fix security issues in the code

Issues found
  - Successful login: JavaScript injection
  - Unsuccessful: status html is returned, instead of exiting the program
  - Check for other improperly handled errors

Solution:
- Instructor included a template for the loginHTML (which I've added here)
*/

// use http.template package
var (
	loginHTML = `<!DOCTYPE html>
<html>
	<body>
		<form method="post">
			<h2>Please Login</h2>
			User: <input name="user"> <br/>
			Password: <input type="password" name="passwd"> <br/>
			<input type="submit"/>
		</form>
	<body>
</html>
`

	statusHTML = `<!DOCTYPE html>
<html>
	<body>
		<h2>Status</h2>
		{{.}}
	</body>
</html>
`
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	loginTemplate := template.Must(template.New("login").Parse(loginHTML))    //Create an HTML template for the login screen
	statusTemplate := template.Must(template.New("status").Parse(statusHTML)) //Create an HTML template for the status output

	if r.Method != http.MethodPost {
		loginTemplate.Execute(w, nil)
		return
	}

	user, passwd := r.FormValue("user"), r.FormValue("passwd")
	if !authUser(user, passwd) {
		//Remove the pw from the error
		http.Error(w, fmt.Sprintf("%s - bad login", user), http.StatusUnauthorized)
		return //return instead of continuing on failed login
	}

	statusTemplate.Execute(w, getStatus())
}

func main() {
	http.HandleFunc("/status", statusHandler)

	log.Println("Server is running on :8080") //Add log message indicating server is running
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
