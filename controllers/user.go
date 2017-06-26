package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/iamclaytonray/go-commerce/models"
	"github.com/julienschmidt/httprouter"
)

// UserController guys
type UserController struct {
	session *mgo.Session
}

// NewUserController guys
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// CreateUser people
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	c := uc.session.DB("go-commerce").C("users")

	json.NewDecoder(r.Body).Decode(&u)

	u.ID = bson.NewObjectId()

	c.Insert(u)

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}
