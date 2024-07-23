package handlers

import (
	"front-fin/views/home"
	"net/http"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
