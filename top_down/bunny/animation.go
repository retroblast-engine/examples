package bunny

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Animation holds the animation details for a state.
type Animation struct {
	index       int
	totalFrames int
	lastChange  time.Time
	duration    []time.Duration
}

// initAnimation initializes the animation for the given state.
func (b *Bunny) initAnimation(state State) {
	b.animation = &Animation{
		index:       0,
		totalFrames: len(b.file.State[state].Frames),
		lastChange:  time.Now(),
		duration:    b.file.State[state].FrameDuration[state],
	}
	b.sprite = b.file.State[state].Frames[0]
}

// spriteFrame returns the frame image for the given state and frame index.
func (b *Bunny) spriteFrame(state State, frameIdx int) *ebiten.Image {
	return b.file.State[state].Frames[frameIdx]
}

// nextFrame returns the next frame image for the current state.
func (b *Bunny) nextFrame() *ebiten.Image {
	if time.Since(b.animation.lastChange) >= b.animation.duration[b.animation.index] {
		b.animation.index++
		if b.animation.index >= b.animation.totalFrames {
			b.animation.index = 0
		}
		b.animation.lastChange = time.Now()
	}

	return b.file.State[b.state].Frames[b.animation.index]
}

// currentFrame returns the current frame image for the bunny's state.
func (b *Bunny) currentFrame() *ebiten.Image {
	if b.file.State[b.state].HasAnimations {
		return b.nextFrame()
	}
	return b.spriteFrame(b.state, 0)
}

// SetState updates the bunny's state and reinitializes the animation if the state has changed.
func (b *Bunny) SetState(state State) {
	if b.state != state {
		b.state = state
		b.initAnimation(state)
	}
}
	