package util

import (
	"net/http"
	"time"
	"errors"
)

type Session struct {
	Id int
	Uuid string
	Email string
	UserId int
	CreateAt time.Time
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")

	if err == nil {
		sess = data.Session{Uuid: cookie.Value}

		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}

	return
}