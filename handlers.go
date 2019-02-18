package main

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/unrolled/render"
)

func createMatchHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		guid := uuid.New()
		w.Header().Add("Location", "/matches/"+guid.String())
		formatter.JSON(w, http.StatusCreated, struct{ Test string }{"This is a test"})
	}
}
