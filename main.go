package main

import (
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	screen.Init()
	screen.Clear()
	gameMap := NewMap(screen)
	player := NewPlayer(screen)
	enemy := NewEnemy(screen)
	screen.Show()

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyESC:
				screen.Fini()
				os.Exit(0)
			}
		}

		screen.Clear()
		gameMap.DrawMap()
		player.draw(screen)
		player.update(screen)
		enemy.Update()
		screen.Show()
	}
}
