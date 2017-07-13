package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamclaytonray/go-commerce/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// OrderController is a struct!!!!
type OrderController struct {
	session *mgo.Session
}

// NewOrderController needs a good comment
func NewOrderController(s *mgo.Session) *OrderController {
	return &OrderController{s}
}

// GetOrders will GET all Orders
func (oc OrderController) GetOrders(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := oc.session.DB("go-commerce").C("orders")
	om := models.Orders{}

	if err := c.Find(nil).All(&om); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(om)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// GetOrder will GET an Order by the params of "slug"
func (oc OrderController) GetOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := oc.session.DB("go-commerce").C("orders")
	s := p.ByName("slug")
	om := models.Order{}

	if err := c.Find(bson.M{"slug": s}).One(&om); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(om)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateOrder will POST an Order when a Product has a status == "paid"
func (oc OrderController) CreateOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := oc.session.DB("go-commerce").C("orders")
	om := models.Order{}

	json.NewDecoder(r.Body).Decode(&om)

	if err := c.Insert(om); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(om)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// UpdateOrder will PUT an Order by the params of "slug"
func (oc OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := oc.session.DB("go-commerce").C("orders")
	s := p.ByName("slug")
	om := models.Order{}

	err := json.NewDecoder(r.Body).Decode(&om)

	c.Update(bson.M{"slug": s}, &om)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
}

// DeleteOrder will DELETE an Order by the params of "slug"
// func (oc OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// }
