package main

import (
	"embed"
	handlers "front-fin/handlers"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := chi.NewMux()

	router.Handle("/*", Public())
	router.Get("/", handlers.Make(handlers.HandlerHome))
	router.Get("/login", handlers.Make(handlers.HandleLogin))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	if err := http.ListenAndServe(listenAddr, router); err != nil {
		panic(err)
	}
}

//go:embed public
var publicFS embed.FS

func Public() http.Handler {
	return http.FileServerFS(publicFS)
}
