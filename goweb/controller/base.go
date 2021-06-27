package controller

import "html/template"

const (
	basePath      = "template"
	baseHTMLName  = "base.html"
	secondPath    = "content"
	indexHTMLName = "index.html"
)

var (
	homeController home
	templates      map[string]*template.Template
)

func init() {
	templates = PopulateTemplates()
}

func StartUp() {
	homeController.registerRoutes()
}
