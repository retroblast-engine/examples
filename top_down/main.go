package main

import (
    "image/color"
    "log"

    "github.com/retroblast-engine/examples/top_down/bunny"

    "github.com/hajimehoshi/ebiten/v2"
)

var mybunny, _ = bunny.NewBunny(64, 64, "assets/aseprite/bunny.aseprite")


type Game struct {
}

func (g *Game) Update() error {
    mybunny.Update()

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.NRGBA{0, 0, 0, 255})
    mybunny.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 128, 128
}

func main() {

    ebiten.SetWindowSize(128*4, 128*4)
    ebiten.SetWindowTitle("Tutorial 1")

    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }

}
