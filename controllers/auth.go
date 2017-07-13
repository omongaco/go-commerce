package controllers

import mgo "gopkg.in/mgo.v2"

// AuthController is a struct
type AuthController struct {
	session *mgo.Session
}

// Login will allow a User to login to the application and will use JWTs
func Login() {

}

// Logout will remove a User's JWT from the Header
func Logout() {

}

// ForgotPassword will send an email to the User with a link to reset their password
func ForgotPassword() {

}

// ResetPassword will have a comment later
func ResetPassword() {

}
