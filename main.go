package main

import (
	"github.com/gdamore/tcell"
)

func main() {
	screen, _ := tcell.NewScreen()
	screen.Init()
	screen.Clear()
	x := 1
	y := 1
	player := rune('@')
	screen.SetContent(x, y, player, []rune{}, tcell.StyleDefault)
	screen.Show()

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyLeft:
				x += -1
			case tcell.KeyRight:
				x += 1
			case tcell.KeyUp:
				y += -1
			case tcell.KeyDown:
				y += 1
			}
		}

		screen.Clear()
		screen.SetContent(x, y, player, []rune{}, tcell.StyleDefault)
		screen.Show()
	}
}
