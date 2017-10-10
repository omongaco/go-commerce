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
	order := models.Orders{}

	if err := c.Find(nil).All(&order); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

// GetOrder will GET an Order by the params of "slug"
func (oc OrderController) GetOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := oc.session.DB("go-commerce").C("orders")
	s := p.ByName("slug")
	order := models.Order{}

	if err := c.Find(bson.M{"slug": s}).One(&order); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

// CreateOrder will POST an Order when a Product has a status == "paid"
func (oc OrderController) CreateOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := oc.session.DB("go-commerce").C("orders")
	order := models.Order{}

	json.NewDecoder(r.Body).Decode(&order)

	if err := c.Insert(order); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

// UpdateOrder will PUT an Order by the params of "slug"
func (oc OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := oc.session.DB("go-commerce").C("orders")
	s := p.ByName("slug")
	order := models.Order{}

	err := json.NewDecoder(r.Body).Decode(&order)

	if err = c.Update(bson.M{"slug": s}, &order); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

// DeleteOrder will DELETE an Order by the params of "slug"
func (oc OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := oc.session.DB("go-commerce").C("orders")
	s := p.ByName("slug")

	if err := c.Remove(bson.M{"slug": s}); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
