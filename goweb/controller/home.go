package controller

import (
	"Forward/goweb/viewModel"
	"net/http"
)

type home struct {}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := viewModel.IndexViewModelOp{}
	v := vop.GetView()
	templates[indexHTMLName].Execute(w, &v)
}
