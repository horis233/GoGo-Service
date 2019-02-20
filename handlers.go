package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	
	"github.com/cloudnativego/gogo-engine"
	"github.com/unrolled/render"
)

func createMatchHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		payload, _ := ioutil.ReadAll(req.Body)
		var newMatchRequest newMatchRequest
		json.Unmarshal(payload, &newMatchRequest)

		newMatch := gogo.NewMatch(newMatchRequest.GridSize)		
		repo.addMatch(newMatch)
		guid := uuid.New()
		w.Header().Add("Location", "/matches/"+guid.String())
		formatter.JSON(w, http.StatusCreated, &newMatchResponse{ID: guid.String(), GridSize: newMatch.GridSize})
	}
}
