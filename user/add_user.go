package user

import (
	"fmt"
	"log"
	"bytes"
	"net/http"
	"html/template"
	"encoding/json"
)

type Msg struct {
	Success string
	Error string
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var msg Msg

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		email := r.FormValue("email")

		// If any field is empty
		if firstname != "" && lastname != "" && email != "" {
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
	
			if resp.Status == "200 OK" {
				templ, _ := template.ParseFiles(
					"template/add.html",
					"template/err.html",
				)

				msg.Success = "User created successfully"
				msg.Error = ""
		
				templ.Execute(w, &msg)
			}
		} else {
			templ, _ := template.ParseFiles(
				"template/add.html",
				"template/err.html",
			)

			msg.Error = "No empty fields allowed"
			msg.Success = ""
	
			templ.Execute(w, &msg)
		} // End empty fields

		
	}
	templ, _ := template.ParseFiles(
		"template/add.html",
		"template/err.html",
	)

	templ.Execute(w, nil)
}