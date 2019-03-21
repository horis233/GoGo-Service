package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
		var mr newMatchResponse
		mr.parse(newMatch)
		w.Header().Add("Location", "/matches/"+newMatch.ID)
		formatter.JSON(w, http.StatusCreated, &mr)
	}
}

func getMatchListHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		repoMatches, err := repo.getMatches()
		if err == nil {
			matches := make([]newMatchResponse, len(repoMatches))
			for idx, match := range repoMatches {
				matches[idx].parse(match)
			}
			formatter.JSON(w, http.StatusOK, matches)
		} else {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		}
	}
}

func getMatchDetailsHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		matchID := vars["id"]
		match, err := repo.getMatch(matchID)
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		} else {
			var mdr matchDetailsResponse
			mdr.parse(match)
			formatter.JSON(w, http.StatusOK, &mdr)
		}
	}
}

func addMoveHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		matchID := vars["id"]
		match, err := repo.getMatch(matchID)
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		} else {
			payload, _ := ioutil.ReadAll(req.Body)
			var moveRequest newMoveRequest
			err := json.Unmarshal(payload, &moveRequest)
			newBoard, err := match.GameBoard.PerformMove(gogo.Move{Player: moveRequest.Player, Position: gogo.Coordinate{X: moveRequest.Position.X, Y: moveRequest.Position.Y}})
			// TODO - this only works because we're using an in-memory repository. A stateless external one
			// would need another call here to persist changes to the match.
			if err != nil {
				match.GameBoard = newBoard
			}
			formatter.JSON(w, http.StatusCreated, matchDetailsResponse{ID: matchID, GameBoard: match.GameBoard.Positions})
		}
	}
}
