package handlers

import (
	"encoding/json"
	"exercise/schema"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s Server) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	products_list := []schema.Product{}
	s.Db.Model(&schema.Product{}).Preload("Rating").Find(&products_list)
	res := []schema.Rating{}
	for _, i := range products_list {

		if i.ID == uint(id) {
			res = i.Rating
			break
		}
	}
	product_json, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(product_json)
}

func (s Server) CreateReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	body, _ := ioutil.ReadAll(r.Body)
	var rat schema.Rating
	err := json.Unmarshal(body, &rat)
	if err != nil {
		panic(err)
	}
	rat.ProductID = uint(id)
	// db.Model(&Product{}).Where("id = ?", id).Association("Rating").Append(&rat)
	s.Db.Model(&schema.Rating{}).Create(&rat)
	product_review_json, _ := json.Marshal(rat)
	w.Header().Set("Content-Type", "application/json")
	w.Write(product_review_json)
}

func (s Server) UpdateReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rid, _ := strconv.Atoi(params["rid"])

	var rev schema.Rating
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &rev)
	if err != nil {
		panic(err)
	}
	s.Db.Model(&schema.Rating{}).Where("id=?", rid).Updates(rev)
	response, _ := json.Marshal(struct{ Message string }{Message: "Review Successfully Updated"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (s Server) DeleteReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	s.Db.Delete(&schema.Rating{}, id)
	response, _ := json.Marshal(struct{ Message string }{Message: "Review Successfully Deleted"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
