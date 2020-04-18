package main

import "github.com/gdamore/tcell"

type Player struct {
	X         int
	Y         int
	Character rune
	Screen    tcell.Screen
}

func NewPlayer(screen tcell.Screen) *Player {
	return &Player{
		X:         4,
		Y:         4,
		Character: '@',
		Screen:    screen,
	}
}

func (p *Player) Update() {
	p.Screen.SetContent(p.X, p.Y, p.Character, nil, tcell.StyleDefault)
}
