package main

import (
	"fmt"
	"net/http"
	"time"
	// "net/url"
	"html/template"
	// "github.com/KING-SAMM/go-api-client/template"
	"github.com/KING-SAMM/go-api-client/user"
	"github.com/KING-SAMM/go-api-client/get_json"
)

//09020377290

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("template/index.html")

	if err != nil {
		fmt.Println("Could not find template")
	}

	templ.Execute(w, nil)
}


func main() {
	// Serve static files from the "public" directory
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Timeout after 10s if no response
	get_json.Client = &http.Client{Timeout: 10 * time.Second}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/user/", user.GetSingleUser)
	http.HandleFunc("/users", user.GetUsers)
	http.HandleFunc("/user/add", user.AddUser)

	// urlOptions := &url.URL{
	// 	Scheme: "https",
	// 	Host: "domain.com",
	// 	Path: "/shop",
	// 	RawQuery: "user=kcsamm",
	// }

	// coolurl := urlOptions.String()
	// fmt.Println(coolurl)

	// user.GetUser()

	http.ListenAndServe(":8081", nil)
}