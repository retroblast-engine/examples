package bunny

import (
    "fmt"

    "github.com/retroblast-engine/asevre"
    "github.com/hajimehoshi/ebiten/v2"
)

// Bunny represents the bunny entity.
type Bunny struct {
	x, y      float64
	file      asevre.ASEFile
	sprite    *ebiten.Image
	state     State
	animation *Animation
	direction Direction
}

// NewBunny creates a new bunny instance.
func NewBunny(x, y float64, file string) (*Bunny, error) {
	aseprite, err := asevre.ParseAseprite(file)
	if err != nil {
		return nil, fmt.Errorf("failed to parse aseprite file: %w", err)
	}
		
	bunny := &Bunny{
		x:     x,
		y:     y,
		file:  aseprite,
		state: idle_front,
		direction: Right,
	}
		
	bunny.initAnimation(bunny.state)
	
	return bunny, nil
}
