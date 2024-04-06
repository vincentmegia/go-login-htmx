package api

import "net/http"

type UsersHandler struct{}

func (uh *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a sample response"))
}
