package user

import (
	"fmt"
	"log"
	"github.com/KING-SAMM/go-api-client/get_json"

	//Added
	"html/template"
	"net/http"
	"strconv"
)

type User struct {
	ID int				`json:"id"`
	Firstname string	`json:"firstname"`
	Lastname string		`json:"lastname"`
	Email string		`json:"email"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:3000/users"
	var users []User
		
	// var user *User
	err := get_json.GetJson(url, &users)

	if err != nil {
		fmt.Println("Oops. Something went wrong getting all users...")
		fmt.Printf("Error is: %v\n", err)
		// fmt.Printf("Error is: %v\n", err)
		templ, _ := template.ParseFiles(
			"template/users.html",
			"template/err.html",
		)

		templ.Execute(w, nil)
	} else if r.Method == "POST" {
		// Get the user id from the POST form data
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Fatal("Could not convert id to int")
		}

		// Find the user with the matching id
		var user *User
		for _, u := range users {
			if u.ID == id {
				user = &u
				break
			}
		}

		// If the user wasn't found, return a 404 error
		if user == nil {
			http.NotFound(w, r)
			return
		}
		templ, _ := template.ParseFiles(
			"template/user.html",
			"template/err.html",
		)

		fmt.Println(user)
	
		templ.Execute(w, user)
	} else {
		
		// Added
		templ, err := template.ParseFiles(
			"template/users.html",
			"template/err.html",
		)

		if err != nil {
			fmt.Println("Could not find template NOWW!!!")
		}

		templ.Execute(w, &users)
	}
	
}
