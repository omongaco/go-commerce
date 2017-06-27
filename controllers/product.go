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
	pm := models.Products{}

	if err := c.Find(nil).All(&pm); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(pm)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", uj)
}

// GetProduct finds a Product with a specified ID
func (pc ProductController) GetProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := pc.session.DB("go-commerce").C("products")
	id := p.ByName("id")
	pm := models.Product{}

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	c.FindId(oid).One(&pm)

	uj, _ := json.Marshal(pm)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateProduct inserts a new Product Collection into MongoDB
func (pc ProductController) CreateProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := pc.session.DB("go-commerce").C("products")
	pm := models.Product{}

	json.NewDecoder(r.Body).Decode(&pm)

	pm.ID = bson.NewObjectId()

	c.Insert(pm)

	uj, _ := json.Marshal(pm)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}
