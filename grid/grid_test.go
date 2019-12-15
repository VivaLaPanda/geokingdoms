package grid

import (
	"testing"

	"github.com/VivaLaPanda/antipath/entity/player"
)

func TestNewState(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Panic while checking if new state is valid: Err: %v", r)
		}
	}()

	NewState()
}

func TestGetTile(t *testing.T) {
	testState := NewState()

	testPos := Coordinates{0, 0}
	_, err := testState.GetTile(testPos)
	if err != nil {
		t.Errorf("Getting valid tile produced err: %v", err)
		return
	}

	testBadPos := Coordinates{200, 200}
	_, err = testState.GetTile(testBadPos)
	if err == nil {
		t.Errorf("Indexing off of the state should result in an error")
	}
}

func TestNewEntity(t *testing.T) {
	testState := NewState()
	pos := Coordinates{0, 0}
	player := player.NewPlayer()
	_, err := testState.NewEntity(player, pos)
	if err != nil {
		t.Errorf("Placing a valid entity into the state at a valid pos produced an error: %v", err)
	}
	_, err = testState.NewEntity(player, pos)
	if err == nil {
		t.Errorf("Placing an entity into a full space didn't produce an error")
	}
}

func TestGetEntityPos(t *testing.T) {
	testState := NewState()
	pos := Coordinates{0, 0}
	player := player.NewPlayer()
	playerID, _ := testState.NewEntity(player, pos)
	actualPos, exists := testState.GetEntityPos(playerID)
	if exists == false || pos != actualPos {
		t.Errorf("Was unable to properly fetch newly created entity's position")
	}
}

func TestMove(t *testing.T) {
	testState := NewState()
	pos := Coordinates{50, 50}
	testPlayer := player.NewPlayer()
	playerID, err := testState.NewEntity(testPlayer, pos)

	// Check a small move
	err = testState.Move(playerID, Up, testPlayer.Speed(), testPlayer.Altitude)
	if err != nil {
		t.Errorf("Moving the player to an empty space resulted in an error")
	}
	newPos, _ := testState.GetEntityPos(playerID)
	expectedPos := pos
	expectedPos.Y -= 1
	if newPos != expectedPos {
		t.Errorf("Move didn't result in the expected location. A: %v, E: %v", newPos, expectedPos)
	}
}
