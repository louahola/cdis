package web

import (
	"github.com/louahola/cdis/repository"
	"net/http"
	"html/template"
	"path/filepath"
	"os"
	"log"
	"github.com/louahola/auth"
	"fmt"
)

type WebManager struct {
	templates *template.Template
	repo repository.Repository
	sessionManager *auth.SessionManager
}

func (this *WebManager) Initialize(repo repository.Repository) {
	this.repo = repo
	this.sessionManager = &auth.SessionManager{SessionRepo: repo}

	basePath := "resources/templates/"
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// don't process folders themselves
		if info.IsDir() {
			return nil
		}
		templateName := path[len(basePath):]
		if this.templates == nil {
			this.templates = template.New(templateName)
			this.templates.Delims("{{%", "%}}")
			_, err = this.templates.ParseFiles(path)
		} else {
			_, err = this.templates.New(templateName).ParseFiles(path)
		}
		log.Printf("Processed template %s\n", templateName)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/favicon", IcoHandler{})
	http.Handle("/", this.WebHandler(IndexHandler{}))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("resources/js"))))
	http.HandleFunc("/auth", auth.Home)
	http.HandleFunc("/FBLogin", auth.FBLogin)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)
func (this *WebManager) WebHandler(h http.Handler) http.Handler {
	return	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		token := r.FormValue("token")
		if token == "" {
			//login
			http.Redirect(w, r, "/auth", 300)
			return
		} else {
			//GetUser(SessionToken)
			user, err := this.sessionManager.GetUser(token)
			if err != nil {
				http.Redirect(w, r, "/auth", 300)
				return
			}
			fmt.Println("User %s is accessing <resource>", user.Email)
			h.ServeHTTP(w, r)
			//IsAuthorized(User, resource)
		}
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
	templates *template.Template
}
func (this IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	passedObj := greetings{
		Intro:    "Hello from Go!",
		Messages: []string{"Hello!", "Hi!", "¡Hola!", "Bonjour!", "Ciao!", "<script>evilScript()</script>"},
	}
	this.templates.ExecuteTemplate(w, "homePage", passedObj)
}

