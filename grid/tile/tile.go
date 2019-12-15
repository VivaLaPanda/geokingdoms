package tile

import (
	"fmt"

	"github.com/VivaLaPanda/antipath/entity/player"
	"github.com/VivaLaPanda/geokingdoms/entity/claim"
)

type Tile struct {
	player *player.Player
	claim  *claim.Claim
}

func (tile *Tile) SetPlayer(player *player.Player) error {
	if tile.player != nil {
		return fmt.Errorf("can only SetPlayer if player is already nil, remove before setting")
	}
	tile.player = player

	return nil
}

func (tile *Tile) PopPlayer() *player.Player {
	var ref *player.Player // declare so the pointer logic is a little clearer
	ref = tile.player
	tile.player = nil
	return ref
}
