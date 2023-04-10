package user

import (
	"fmt"
	"log"
	"bytes"
	"net/http"
	"html/template"
	"encoding/json"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		email := r.FormValue("email")
		user := User{
			Firstname: firstname,
			Lastname: lastname,
			Email: email,
		}

		userjson, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User is: %v\n", user)
		// fmt.Printf("Userjson byte slice is: %v\n", userjson)
		fmt.Printf("Userjson string is: %v\n", string(userjson))

		resp, err := http.Post("http://localhost:3000/user", "application/json", bytes.NewBuffer(userjson))
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		log.Printf("Response status: %s", resp.Status)

		// templ, _ := template.ParseFiles(
		// 	"template/add.html",
		// 	"template/err.html",
		// )

		// templ.Execute(w, nil)
		
	}
	templ, _ := template.ParseFiles(
		"template/add.html",
		"template/err.html",
	)

	templ.Execute(w, nil)
}