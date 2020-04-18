package main

import "github.com/gdamore/tcell"

type Player struct {
	X         int
	Y         int
	Character rune
	Screen    *tcell.Screen
}

func NewPlayer(screen *tcell.Screen) *Player {
	return &Player{
		X:         0,
		Y:         0,
		Character: '@',
		Screen:    screen,
	}
}

func (p *Player) Update() {
	screen := *p.Screen
	screen.SetContent(p.X, p.Y, p.Character, nil, tcell.StyleDefault)
}
