package grid

import (
	"fmt"

	"github.com/VivaLaPanda/antipath/entity"
	"github.com/VivaLaPanda/antipath/grid/tile"
	uuid "github.com/satori/go.uuid"
)

type State struct {
	grid     ChunkLoader
	entities map[string]Coordinates
}

type Coordinates struct {
	X, Y int
}

type Direction int

const (
	Up    Direction = iota
	Right Direction = iota
	Left  Direction = iota
	Down  Direction = iota
	None  Direction = iota
)

var chunkSize = 16

func NewState() *State {
	return &State{
		grid:     *NewChunkLoader(chunkSize),
		entities: make(map[string]Coordinates),
	}
}

func (s *State) GetTile(pos Coordinates) (*tile.Tile, error) {
	return s.grid.GetTile(pos)
}

func (s *State) NewEntity(data entity.Entity, pos Coordinates) (id string, err error) {
	targetTile, err := s.grid.GetTile(pos)
	if err != nil {
		return "", err
	}

	if err := targetTile.SetEntity(data); err != nil {
		return "", fmt.Errorf("provided pos can't contain an entity, already full. Tile %v", targetTile)
	}

	id = uuid.Must(uuid.NewV4()).String()

	s.entities[id] = pos

	return id, nil
}

func (s *State) GetEntityPos(entityID string) (pos Coordinates, exists bool) {
	pos, exists = s.entities[entityID]
	return
}

func (s *State) Move(entityID string, dir Direction, speed int, altitude int) (err error) {
	// Get the location of the entity
	sourcePos, exists := s.entities[entityID]
	if !exists {
		return fmt.Errorf("provided entity ID not valid. ID: %s", entityID)
	}
	// Get the tile data at that location
	sourceTile, err := s.grid.GetTile(sourcePos)
	if err != nil {
		return fmt.Errorf("couldn't get tile at provided pos, pos: %v, err: %s", sourcePos, err)
	}

	// Calculate the total movement
	targetPos := sourcePos
	var targetTile *tile.Tile
	switch dir {
	case Up:
		targetPos.Y -= speed
	case Down:
		targetPos.Y += speed
	case Left:
		targetPos.X -= speed
	case Right:
		targetPos.X += speed
	}

	// Simulate entity movement with collision rules
	targetTile, _ = s.grid.GetTile(targetPos)

	// Move the entity
	entityData := sourceTile.PopEntity()
	targetTile.SetEntity(entityData)
	s.entities[entityID] = targetPos

	return nil
}
