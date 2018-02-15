package main

//go:generate go-bindata-assetfs templates/... pages/...

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"absensi-api/config"
	"github.com/gorilla/mux"
	"fmt"
	"absensi-api/database"
	"database/sql"
)

var templates = template.New("")

type Model struct {
	Bolos int
	Hadir int
	Izin int
}

// Render template
func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	err := config.Parse("config/config.toml")
	if err != nil {
		log.Fatalf("Can't read config.toml file, %s", err)
	}

	for _, path := range AssetNames() {
		bytes, err := Asset(path)
		if err != nil {
			log.Panicf("Unable to parse: path=%s, err=%s", path, err)
		}
		templates.New(path).Parse(string(bytes))
	}

	router := mux.NewRouter().StrictSlash(true)
	port := strconv.Itoa(config.APIConfig.ListenPort)
	address := config.APIConfig.ListenURL + ":" + port
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Handle static files
	router.PathPrefix("/static").HandlerFunc(noDirListing(http.StripPrefix("/static", http.FileServer(http.Dir("static/")))))

	// Handle form.html
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Model
		var pres []Model

		// Check method
		fmt.Println("method:", r.Method)

		if r.Method == "GET" {
			renderTemplate(w, "pages/form.html", &pres)
		} else {
			r.ParseForm()

			// Insert to  database logic
			fmt.Println("nim", r.Form["nim"])
			fmt.Println("value", r.Form["present-button"])
			nimInt, _ := strconv.Atoi(r.Form["nim"][0])
			valInt, _ := strconv.Atoi(r.Form["present-button"][0])
			err := database.InsertPresence(nimInt, valInt)
			if err != nil {
				fmt.Println("Can't insert to database", err)
			}
			//renderTemplate(w, "pages/form.html", &pres)

			// Refresh page
			http.Redirect(w, r, "localhost:8080/", http.StatusMovedPermanently)
		}
	})

	// Handle report.html
	router.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		// Model
		//presence := models.Presence{}
		model := Model{}

		// Check method
		fmt.Println("method:", r.Method)

		if r.Method == "GET" {
			renderTemplate(w, "pages/report.html", &model)
		} else {
			r.ParseForm()

			// Get from form
			fmt.Println("nim", r.Form["nim"])
			if err != nil {
				fmt.Println(err)
			}
			nim := r.Form["nim"][0]
			db, err := sql.Open("mysql", "root:@/presence")
			if err != nil {
				return
			}
			rows, err := db.Query("SELECT COUNT(*) FROM presence WHERE nim = ? AND presence = ?", nim, 0)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&model.Bolos)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(model.Bolos)
				//pres = append(pres, Model{Bolos: model.Bolos})
			}
			rows, err = db.Query("SELECT COUNT(*) FROM presence WHERE nim = ? AND presence = ?", nim, 1)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&model.Hadir)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(model.Hadir)
				//pres = append(pres, Model{Hadir: model.Hadir})
			}
			rows, err = db.Query("SELECT COUNT(*) FROM presence WHERE nim = ? AND presence = ?", nim, 2)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&model.Izin)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(model.Izin)
				//pres = append(pres, Model{Izin: model.Izin})
			}
			renderTemplate(w, "pages/report.html", &model)

			// Refresh page
			//http.Redirect(w, r, "localhost:8080/report", http.StatusMovedPermanently)
		}

	})

	log.Printf("Server: %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
