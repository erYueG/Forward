package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	UserName string
}

type IndexViewModel struct {
	Title string
	User *User
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := &User{
			UserName: "Owen",
		}
		ivm := &IndexViewModel{
			Title: "HomePage",
			User:user,
		}
		tem, err := template.ParseFiles("template/index.html")
		if err != nil {
			log.Fatalf("temlate paseFile is failed,err is %s", err.Error())
		}
		tem.Execute(w, &ivm)
	})
	http.ListenAndServe(":8888", nil)
}
