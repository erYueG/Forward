package main

import (
	"Forward/goweb/controller"
	"net/http"
)

func main() {
	controller.StartUp()
	http.ListenAndServe(":8888", nil)
}
