package integrations_test

import (
	"errors"
	"fmt"
	"os"

	"github.com/cloudnativego/cfmgo"
)

func getRepo(collectionName string) (col cfmgo.Collection, err error) {
	host := os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	port := os.Getenv("MONGO_PORT_27017_TCP_PORT")

	if len(host) == 0 {
		err = errors.New("Could not retrieve mongo host information.")
		return
	}
	if len(port) == 0 {
		port = "27017"
	}

	uri := fmt.Sprintf("mongodb://%s:%s/fake-guid", host, port)
	col = cfmgo.Connect(cfmgo.NewCollectionDialer, uri, collectionName)
	return
}

/*func TestAddMatchToRepo(t *testing.T) {
	matchesCollection, err := getRepo(MatchesCollectionName)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	repo := NewMongoMatchRepository(matchesCollection)
	match := gogo.NewMatch(19, "buckshank", "d'squarius")
	err = repo.AddMatch(match)
	if err != nil {
		t.Errorf("Error adding match to mongo: %v", err)
	}

	matches, err := repo.GetMatches()
	if err != nil {
		t.Errorf("Error retrieving matches: %v", err)
	}
	if len(matches) == 0 {
		t.Errorf("Expected matches length to be greater than 0; received %d", len(matches))
	}

	foundMatch, err := repo.GetMatch(match.ID)
	if err != nil {
		t.Errorf("Unable to find match with ID: %v... %s", match.ID, err)
	}
	if foundMatch.GridSize != match.GridSize || foundMatch.PlayerBlack != match.PlayerBlack {
		t.Errorf("Unexpected match results: %v", foundMatch)
	}
}

func TestUpdateMatch(t *testing.T) {
	matchesCollection, err := getRepo(MatchesCollectionName)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	repo := NewMongoMatchRepository(matchesCollection)
	match := gogo.NewMatch(13, "d'pez", "poopsie")
	err = repo.AddMatch(match)
	if err != nil {
		t.Errorf("Error adding match to mongo: %v", err)
	}

	match.TurnCount = 3
	err = repo.UpdateMatch(match.ID, match)
	if err != nil {
		t.Errorf("Error updating match: %v", err)
	}

	foundMatch, err := repo.GetMatch(match.ID)
	if err != nil {
		t.Errorf("Unable to find match with ID: %v... %s", match.ID, err)
	}
	if foundMatch.TurnCount != match.TurnCount {
		t.Errorf("Match update failed: expected %d; received %d", match.TurnCount, foundMatch.TurnCount)
	}
}
*/
