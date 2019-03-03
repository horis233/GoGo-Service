package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	
	"github.com/cloudnativego/gogo-engine"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func createMatchHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		payload, _ := ioutil.ReadAll(req.Body)
		var newMatchRequest newMatchRequest
		err := json.Unmarshal(payload, &newMatchRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse create match request")
			return
		}
		if !newMatchRequest.isValid() {
			formatter.Text(w, http.StatusBadRequest, "Invalid new match request")
			return
		}
		newMatch := gogo.NewMatch(newMatchRequest.GridSize, newMatchRequest.PlayerBlack, newMatchRequest.PlayerWhite)
		repo.addMatch(newMatch)
		guid := uuid.New()
		w.Header().Add("Location", "/matches/"+guid.String())
		formatter.JSON(w, http.StatusCreated, &newMatchResponse{ID: newMatch.ID, GridSize: newMatch.GridSize, PlayerBlack: newMatchRequest.PlayerBlack, PlayerWhite: newMatchRequest.PlayerWhite})	}
}

func getMatchListHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		repoMatches := repo.getMatches()
		matches := make([]newMatchResponse, len(repoMatches))
		for idx, match := range repoMatches {
			matches[idx] = newMatchResponse{ID: match.ID, GridSize: match.GridSize, PlayerBlack: match.PlayerBlack, PlayerWhite: match.PlayerWhite}
		}
		formatter.JSON(w, http.StatusOK, matches)
	}
}
  
func getMatchDetailsHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		matchID := vars["id"]
		fmt.Printf("Fetching match details for match (%s)\n", matchID)
	}
}