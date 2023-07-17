package util

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

const SESSION_NAME = "team_todo_session"

type Session_info struct {
	Id    string
	Email string
}

func GetId_Session(req *http.Request) (string, error) {
	session, err := store.Get(req, SESSION_NAME)
	if err != nil {
		return "", err
	}
	if session.IsNew {
		return "", nil
	}
	Id := session.Values["id"].(string)
	return Id, nil
}

func SetSession(w http.ResponseWriter, req *http.Request, userId, email string) error {
	session, _ := store.Get(req, SESSION_NAME)
	session.Values["id"] = userId
	session.Values["email"] = email
	return session.Save(req, w)
}
