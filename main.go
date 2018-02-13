package main

import (
	"fmt"
	"absensi-api/config"
	"log"
	"github.com/gorilla/mux"
	"strconv"
	"net/http"
	"time"
)

var formAddress string
var reportAddress string

func main() {
	err := config.Parse("config/config.toml")
	if err != nil {
		log.Fatalf("Can't read config.toml file, %s", err)
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

	router.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/form.html")
	})
	router.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/report.html")
	})

	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("public/css/"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("public/img/"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("public/js/"))))
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("public/fonts/"))))

	log.Printf("Server: %s", server.Addr)
	log.Fatal(server.ListenAndServe())

	fmt.Printf("Hai juga Irene\n")
}
