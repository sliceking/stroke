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
			case tcell.KeyLeft:
				if r, _, _, _ := screen.GetContent(player.X-1, player.Y); r != ' ' {
					continue
				}
				player.X += -1
			case tcell.KeyRight:
				if r, _, _, _ := screen.GetContent(player.X+1, player.Y); r != ' ' {
					continue
				}
				player.X += 1
			case tcell.KeyUp:
				if r, _, _, _ := screen.GetContent(player.X, player.Y-1); r != ' ' {
					continue
				}
				player.Y += -1
			case tcell.KeyDown:
				if r, _, _, _ := screen.GetContent(player.X, player.Y+1); r != ' ' {
					continue
				}
				player.Y += 1
			case tcell.KeyESC:
				screen.Fini()
				os.Exit(0)
			}
		}

		screen.Clear()
		gameMap.DrawMap()
		player.Update()
		enemy.Update()
		screen.Show()
	}
}
