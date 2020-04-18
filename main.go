package main

import (
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	screen, _ := tcell.NewScreen()
	screen.Init()
	screen.Clear()
	gameMap := NewMap(&screen)
	player := NewPlayer(&screen)
	screen.Show()

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyLeft:
				player.X += -1
			case tcell.KeyRight:
				player.X += 1
			case tcell.KeyUp:
				player.Y += -1
			case tcell.KeyDown:
				player.Y += 1
			case tcell.KeyESC:
				screen.Fini()
				os.Exit(0)
			}
		}

		screen.Clear()
		player.Update()
		gameMap.DrawMap()
		screen.Show()
	}
}
