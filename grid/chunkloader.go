package grid

import (
	"fmt"

	"github.com/VivaLaPanda/antipath/grid/tile"
)

type Chunk [][]tile.Tile

// TODO: make chunks a DB with an in memory cache
type ChunkLoader struct {
	chunks    map[string]Chunk
	chunkSize int
}

func NewChunkLoader(chunkSize int) (chunkLoader *ChunkLoader) {
	return &ChunkLoader{
		chunks:    make(map[string]Chunk),
		chunkSize: chunkSize,
	}
}

func (cl *ChunkLoader) GetTile(pos Coordinates) (*tile.Tile, error) {
	xChunk, xPos := pos.X%cl.chunkSize, pos.X/cl.chunkSize
	yChunk, yPos := pos.Y%cl.chunkSize, pos.Y/cl.chunkSize
	hash := fmt.Sprintf("%d%d", xChunk, yChunk)
	chunk, found := cl.chunks[hash]
	if !found {
		chunk := make([][]tile.Tile, cl.chunkSize)
		for idx := range chunk {
			chunk[idx] = make([]tile.Tile, cl.chunkSize)
		}

		cl.chunks[hash] = chunk
	}

	return &chunk[xPos][yPos], nil
}
