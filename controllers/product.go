package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamclaytonray/go-commerce/models"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*

	- Add Pagination
	- Add Filtering

*/

// ProductController is a struct that points to mgo.Session
type ProductController struct {
	session *mgo.Session
}

// NewProductController will have a future comment
func NewProductController(s *mgo.Session) *ProductController {
	return &ProductController{s}
}

// GetProducts finds all Products
func (pc ProductController) GetProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := pc.session.DB("go-commerce").C("products")
	product := models.Products{}

	if err := c.Find(nil).All(&product); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

// GetProduct finds a Product with specified slug params
func (pc ProductController) GetProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := pc.session.DB("go-commerce").C("products")
	slug := p.ByName("slug")
	product := models.Product{}

	if err := c.Find(bson.M{"slug": slug}).One(&product); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateProduct inserts a new Product Collection into MongoDB
func (pc ProductController) CreateProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := pc.session.DB("go-commerce").C("products")
	product := models.Product{}

	err := json.NewDecoder(r.Body).Decode(&product)

	product.ID = bson.NewObjectId()

	if err = c.Insert(product); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

// UpdateProduct will PUT the Product based on the slug params
func (pc ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := pc.session.DB("go-commerce").C("products")
	slug := p.ByName("slug")
	product := models.Product{}

	err := json.NewDecoder(r.Body).Decode(&product)

	if err = c.Update(bson.M{"slug": slug}, &product); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

// DeleteProduct will delete a Product document from the MongoDB collection based on the slug params
func (pc ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := pc.session.DB("go-commerce").C("products")
	slug := p.ByName("slug")

	if err := c.Remove(bson.M{"slug": slug}); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
