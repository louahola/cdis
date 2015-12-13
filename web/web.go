package web

import (
	"github.com/louahola/cdis/repository"
	"net/http"
	"html/template"
	"path/filepath"
	"os"
	"log"
	"github.com/louahola/auth"
)

var templates *template.Template
var repo *repository.MongoRepository

func Initialize(repo repository.Repository) {

	wd, err := os.Getwd()
	if err != nil {
		log.Print(wd)
	}
	basePath := "resources/templates/"
	err = filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// don't process folders themselves
		if info.IsDir() {
			return nil
		}
		templateName := path[len(basePath):]
		if templates == nil {
			templates = template.New(templateName)
			templates.Delims("{{%", "%}}")
			_, err = templates.ParseFiles(path)
		} else {
			_, err = templates.New(templateName).ParseFiles(path)
		}
		log.Printf("Processed template %s\n", templateName)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/favicon", IcoHandler{})
	http.Handle("/", WebHandler(IndexHandler{}))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("resources/js"))))
	http.HandleFunc("/auth", auth.Home)
	http.HandleFunc("/FBLogin", auth.FBLogin)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func WebHandler(h http.Handler) http.Handler {
	return	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		token, err := r.FormValue("token")
		if err != nil {
			//login
		} else {
			//GetUser(SessionToken)

			//IsAuthorized(User, resource)
		}
		repo

		h.ServeHTTP(w, r)
	});
}

type IcoHandler struct {
}
func (this IcoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

type greetings struct {
	Intro string
	Messages []string
}

type IndexHandler struct {
}
func (this IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	passedObj := greetings{
		Intro:    "Hello from Go!",
		Messages: []string{"Hello!", "Hi!", "¡Hola!", "Bonjour!", "Ciao!", "<script>evilScript()</script>"},
	}
	templates.ExecuteTemplate(w, "homePage", passedObj)
}

