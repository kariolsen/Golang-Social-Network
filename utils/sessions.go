package utils

import (
	"crypto/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var secretToken = generateRandomToken(10)
var store = sessions.NewCookieStore([]byte(secretToken))

func GetSession(c *gin.Context) *sessions.Session { // return session info
	session, err := store.Get(c.Request, "session") // Get the session info of this request
	Err(err)
	return session
}

func Get_Id_and_Username(c *gin.Context) (interface{}, interface{}) {
	session := GetSession(c)
	id := session.Values["id"]
	username := session.Values["username"]
	return id, username
}

func Get_Avatar(c *gin.Context) interface{} {
	session := GetSession(c)
	avatar := session.Values["avatar"]
	return avatar
}
