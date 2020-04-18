package main

import "github.com/gdamore/tcell"

type Map struct {
	Height    int
	Width     int
	Character rune
	Screen    *tcell.Screen
}

func NewMap(screen *tcell.Screen) *Map {
	return &Map{
		Height:    10,
		Width:     20,
		Character: '#',
		Screen:    screen,
	}
}

func (m *Map) DrawMap() {
	screen := *m.Screen
	for i := 0; i < m.Width; i++ {
		for j := 0; j < m.Height; j++ {
			if i == 0 || j == 0 {
				screen.SetContent(i, j, '#', nil, tcell.StyleDefault)
			} else if i == 19 || j == 9 {
				screen.SetContent(i, j, '#', nil, tcell.StyleDefault)
			}
		}
	}
}
