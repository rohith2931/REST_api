package handlers

import (
	"encoding/json"
	"exercise/schema"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	Db *gorm.DB
}

func (s Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	products_list := []schema.Product{}
	s.Db.Model(&schema.Product{}).Preload("Variants").Find(&products_list)
	products_json, _ := json.Marshal(products_list)
	w.Header().Set("Content-Type", "application/json")
	w.Write(products_json)
}
func (s Server) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	products_list := []schema.Product{}
	s.Db.Model(&schema.Product{}).Find(&products_list)
	res := schema.Product{}
	for _, i := range products_list {
		if i.ID == uint(id) {
			res = i
			break
		}
	}

	product_json, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(product_json)
}

func (s Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var t schema.Product
	err := json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	s.Db.Create(&t)
	product_create_json, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.Write(product_create_json)
}
