package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"softwaremarket/models"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	tpl *template.Template
)

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=7850576Cc dbname=POMarket sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/product", productHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	products, err := models.AllProducts(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "index.html", products)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	product, err := models.GetProduct(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "product.html", product)
}
