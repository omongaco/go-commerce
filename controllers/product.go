package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamclaytonray/go-commerce/models"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

// ProductController is a struct that points to mgo.Session
type ProductController struct {
	session *mgo.Session
}

// NewProductController will have a future comment
func NewProductController(s *mgo.Session) *ProductController {
	return &ProductController{s}
}

// CreateProduct inserts a new Product Collection into MongoDB
func (pc ProductController) CreateProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := pc.session.DB("go-commerce").C("products")
	pm := models.Product{}

	json.NewDecoder(r.Body).Decode(&pm)

	c.Insert(pm)

	uj, _ := json.Marshal(pm)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}
