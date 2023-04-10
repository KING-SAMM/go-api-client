package user

import (
	"fmt"
	"github.com/KING-SAMM/go-api-client/get_json"
	"html/template"
	"net/http"
)

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	// Get url path after "/user/", i.e the ID
	path := r.URL.Path[len("/user/"):]
	if path == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	id := path

	fmt.Printf("ID is %v\n", id)

	url := `http://localhost:3000/user/`+id

	var user *User

	err := get_json.GetJson(url, &user)

	if err != nil {
		fmt.Printf("Oops. Something went wrong getting user %v\n", id)
		fmt.Printf("Error is: %v\n", err)
		// fmt.Printf("Error is: %v\n", errmsg)
		templ, _ := template.ParseFiles(
			"template/user.html",
			"template/err.html",
		)

		templ.Execute(w, nil)
	} else {
		templ, _ := template.ParseFiles(
			"template/user.html",
			"template/err.html",
		)
	
		fmt.Println(&user)
	
		templ.Execute(w, &user)
	}
}