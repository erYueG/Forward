package controller

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func PopulateTemplates() map[string]*template.Template {

	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/" + baseHTMLName))
	dir, err := os.Open(basePath + "/" + secondPath)
	if err != nil {
		log.Fatalf("Failed to open template directory,err is %s", err.Error())
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatalf("Failed to read contents of template directory,err is %s", err.Error())
	}

	for _, file := range files {
		func() {
			f, err := os.Open(basePath + "/" + secondPath + "/" + file.Name())
			if err != nil {
				log.Fatalf("Failed to open template %s", f.Name())
			}
			defer f.Close()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				log.Fatalf("Failed to read content from template file,err is %s", err.Error())
			}

			tmpl := template.Must(layout.Clone())
			_, err = tmpl.Parse(string(content))
			if err != nil {
				log.Fatalf("Failed to parse content of %s", file.Name())
			}
			result[file.Name()] = tmpl
		}()
	}
	return result
}
