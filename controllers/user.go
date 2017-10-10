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
	user := models.Users{}

	if err := c.Find(nil).All(&user); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

// GetUser finds a User with a specified ID
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := uc.session.DB("go-commerce").C("users")
	username := p.ByName("username")
	user := models.User{}

	if err := c.Find(bson.M{"username": username}).One(&user); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateUser inserts a new User Collection into MongoDB
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := uc.session.DB("go-commerce").C("users")
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)

	user.ID = bson.NewObjectId()

	if err = c.Insert(user); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

// UpdateUser will update a single User with a matching slug as the parameter
func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := uc.session.DB("go-commerce").C("users")
	username := p.ByName("username")
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err = c.Update(bson.M{"username": username}, &user); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// DeleteUser will delete a single user with a matching ID as the parameter
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := uc.session.DB("go-commerce").C("users")
	username := p.ByName("username")

	if err := c.Remove(bson.M{"username": username}); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
