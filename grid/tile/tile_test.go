package tile

import (
	"testing"

	"github.com/VivaLaPanda/antipath/entity/player"
)

func TestSetPlayer(t *testing.T) {
	testTile := Tile{}
	testPlayer := player.NewPlayer()

	err := testTile.SetPlayer(testPlayer)
	if err != nil {
		t.Errorf("Error setting entity on empty tile: %s", err)
		return
	}

	err = testTile.SetPlayer(testPlayer)
	if err == nil {
		t.Errorf("No error setting entity on tile that already has one")
	}
}

func TestPopPlayer(t *testing.T) {
	testTile := Tile{}
	testPlayer := player.NewPlayer()

	err := testTile.SetPlayer(testPlayer)
	if err != nil {
		t.Errorf("Error setting entity on empty tile: %s", err)
		return
	}

	resultPlayer := testTile.PopPlayer()
	if resultPlayer != testPlayer {
		t.Errorf("Push and then pop don't match.")
		return
	}

	resultPlayer = testTile.PopPlayer()
	if resultPlayer != nil {
		t.Errorf("Pop on an empty tile isn't nil.")
		return
	}

	err = testTile.SetPlayer(testPlayer)
	if err != nil {
		t.Errorf("Error setting entity when it should be empty: %s", err)
		return
	}
}
