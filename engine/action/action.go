package action

import "github.com/VivaLaPanda/antipath/grid"

type Set struct {
	Movement grid.Direction
}

func DefaultSet() Set {
	return Set{
		Movement: grid.None,
	}
}
