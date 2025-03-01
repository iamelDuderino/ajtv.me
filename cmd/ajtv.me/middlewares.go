package main

import (
	"net/http"
)

func (x *userInterface) sessionManager(fn func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := x.globalSession.Get(r, globalSessionCookieName)
		if err != nil {
			x.logger.Error.Println(err)
		}
		if s.IsNew {
			err = s.Save(r, w)
			if err != nil {
				x.logger.Error.Println(err)
			}
		}
		fn(w, r)
	}
}
