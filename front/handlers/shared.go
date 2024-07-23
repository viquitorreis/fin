package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type HTTPHandler func(http.ResponseWriter, *http.Request) error

func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("error handling request", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
