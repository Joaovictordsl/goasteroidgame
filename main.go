package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"gogame/game"
)

func main() {

	g := game.NewGame()

	err := ebiten.RunGame(g)

	if err != nil {
		panic(err)
	}
}
