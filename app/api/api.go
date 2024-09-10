package api

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func Render(w http.ResponseWriter, files []string, name string, data interface{}) error {
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to generate template")
		return err
	}
	return ts.ExecuteTemplate(w, name, data)
}

func GoHtmlFilePath(dirPrefix string, name string) string {
	return fmt.Sprint(dirPrefix, "/views/", name, ".gohtml")
}

func CreateApi(dirPrefix string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Render(w, []string{GoHtmlFilePath(dirPrefix, "index")}, "index", nil)
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		name := r.FormValue("name")
		w.Header().Set("Content-Type", "text/plain")
		Render(w, []string{GoHtmlFilePath(dirPrefix, "name")}, "name", name)
	})
}
