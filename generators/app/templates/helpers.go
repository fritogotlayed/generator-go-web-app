package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/juju/loggo"
	"os"
)

var templates *template.Template

func LoadTemplates(templateRootDir string) error {
	if !strings.HasSuffix(templateRootDir, "/") {
		templateRootDir = templateRootDir + "/"
	}

	fileList := []string{}
	err := filepath.Walk(templateRootDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	logger := loggo.GetLogger("web")
	logger.Debugf("Found the following templates: %s", fileList)

	// For now template.ParseFiles uses the file name as the key so we are semi-forced to put our templates in a single
	// directory. Ex: /web/static/templates/index.html is keyed "index". If this becomes problematic for organization we
	// can explore indexing the templates directly and refactoring the below renderTemplate method to use the templates
	// directly. https://golang.org/pkg/html/template/
	templates = template.Must(template.ParseFiles(fileList...))
	return nil
}

func renderTemplate(w http.ResponseWriter, template string, data interface{}) {
	err := templates.ExecuteTemplate(w, template+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type empty struct{}
type semaphore chan empty

func (s semaphore) Acquire(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
