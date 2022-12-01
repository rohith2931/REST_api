package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"exercise/handlers"
	"exercise/schema"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	// db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	//command Line Flags
	var user, pass, dbname, sslmode string
	flag.StringVar(&user, "user", "", "a string var")
	flag.StringVar(&pass, "password", "", "a string var")
	flag.StringVar(&dbname, "dbname", "", "a string var")
	flag.StringVar(&sslmode, "sslmode", "disable", "a string var")
	flag.Parse()
	dbstring := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, pass, dbname, sslmode)
	db, err := gorm.Open("postgres", dbstring)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var s handlers.Server = handlers.Server{db}
	schema.SetUpDB(db)

	router := mux.NewRouter()

	router.HandleFunc("/api/products", s.GetProducts).Methods("GET")

	router.HandleFunc("/api/products/{id}", s.GetProductByID).Methods("GET")

	router.HandleFunc("/api/products/create", s.CreateProduct).Methods("POST")

	router.HandleFunc("/api/products/{id}/reviews", s.GetReviewByID).Methods("GET")

	router.HandleFunc("/api/products/{id}/reviews/create", s.CreateReview).Methods("POST")

	router.HandleFunc("/api/products/{id}/reviews/{rid}", s.UpdateReview).Methods("PATCH")

	router.HandleFunc("/api/reviews/{id}", s.DeleteReview).Methods("DELETE")
	fmt.Println("Server at 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
