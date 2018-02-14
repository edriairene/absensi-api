package main

//go:generate go-bindata-assetfs templates/... pages/...

import (
	"fmt"
	"absensi-api/config"
	"log"
	"github.com/gorilla/mux"
	"strconv"
	"net/http"
	"time"
	"html/template"
)

var formAddress string
var reportAddress string

var templates = template.New("")
type Model struct {
	Title string
}

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	err := config.Parse("config/config.toml")
	if err != nil {
		log.Fatalf("Can't read config.toml file, %s", err)
	}

	for _,path := range AssetNames() {
		bytes, err := Asset(path)
		if err != nil {
			log.Panicf("Unable to parse: path=%s, err=%s", path, err)
		}
		templates.New(path).Parse(string(bytes))
	}

	formAddress = fmt.Sprintf("http://%s:%d", config.APIConfig.FormURL, config.APIConfig.FormPort)
	reportAddress = fmt.Sprintf("http://%s:%d", config.APIConfig.ReportURL, config.APIConfig.ReportPort)

	router := mux.NewRouter()
	port := strconv.Itoa(config.APIConfig.ListenPort)
	address := config.APIConfig.ListenURL + ":" + port
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/login.html")
	})

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/register.html")
	})


	router.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		model := Model{Title: "Form"}
		renderTemplate(w, "pages/form.html", &model)
	})

	router.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/report.html")
	})

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(assetFS())))


	log.Printf("Server: %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
