package bunny

import "github.com/hajimehoshi/ebiten/v2"

// draw draws the bunny's sprite on the screen.
func (b *Bunny) draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}

	// Flip the sprite, if left is pressed
	if b.direction == Left {
		b.hFlip(opts)
	}

	opts.GeoM.Translate(b.x, b.y)

	// Cache the current frame in the sprite field
	b.sprite = b.currentFrame()
	screen.DrawImage(b.sprite, opts)
}

// hFlip modifies the DrawImageOptions to flip the sprite horizontally.
func (b *Bunny) hFlip(opts *ebiten.DrawImageOptions) {
	opts.GeoM.Scale(-1, 1)
	opts.GeoM.Translate(float64(b.sprite.Bounds().Dx()), 0)
}

// vFlip modifies the DrawImageOptions to flip the sprite vertically.
func (b *Bunny) vFlip(opts *ebiten.DrawImageOptions) {
	opts.GeoM.Scale(1, -1)
	opts.GeoM.Translate(0, float64(b.sprite.Bounds().Dy()))
}

// hvFlip modifies the DrawImageOptions to flip the sprite horizontally and vertically.
func (b *Bunny) hvFlip(opts *ebiten.DrawImageOptions) {
	opts.GeoM.Scale(-1, -1)
	opts.GeoM.Translate(float64(b.sprite.Bounds().Dx()), float64(b.sprite.Bounds().Dy()))
}

// Draw draws the bunny on the screen.
func (b *Bunny) Draw(screen *ebiten.Image) {
	b.draw(screen)
}

