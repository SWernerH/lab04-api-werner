package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(rw, r)

		log.Printf("%s %s %d %v",
			r.Method,
			r.URL.Path,
			rw.statusCode,
			time.Since(start),
		)
	})
}

type application struct {
	logger *slog.Logger
}

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "status: available\n")
	app.logger.Info("healthcheck handler called")
}

func (app *application) listBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "list of books (coming soon)\n")
	app.logger.Info("listBooks handler called")
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "get book with id: %s\n", id)
	app.logger.Info("getBook handler called", "id", id)
}

func (app *application) createBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "book created (coming soon)\n")
	app.logger.Info("createBook handler called")
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.WriteHeader(http.StatusNoContent)
	app.logger.Info("deleteBook handler called", "id", id)
}