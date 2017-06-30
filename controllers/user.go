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

// UserController is a struct that points to mgo.Session
type UserController struct {
	session *mgo.Session
}

// NewUserController will have a future comment
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUsers finds all Users
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := uc.session.DB("go-commerce").C("users")
	um := models.Users{}

	if err := c.Find(nil).All(&um); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(um)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// GetUser finds a User with a specified ID
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := uc.session.DB("go-commerce").C("users")
	username := p.ByName("username")
	um := models.User{}

	if err := c.Find(bson.M{"username": username}).One(&um); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(um)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser inserts a new User Collection into MongoDB
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := uc.session.DB("go-commerce").C("users")
	um := models.User{}

	json.NewDecoder(r.Body).Decode(&um)

	um.ID = bson.NewObjectId()

	if err := c.Insert(um); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(um)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// UpdateUser will update a single user with a matching slug as the parameter
func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := uc.session.DB("go-commerce").C("users")
	username := p.ByName("username")
	um := models.User{}

	json.NewDecoder(r.Body)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&um)

	if err != nil {
		panic(err)
	}

	if err = c.Update(bson.M{"username": username}, &um); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(um)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// DeleteUser will delete a single user with a matching ID as the parameter
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := uc.session.DB("go-commerce").C("users")
	username := p.ByName("username")

	if err := c.Remove(bson.M{"username": username}); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(204)
}
