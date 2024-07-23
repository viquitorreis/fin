package handlers

import (
	"front-fin/views/auth"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}
