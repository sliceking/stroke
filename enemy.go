package main

import "github.com/gdamore/tcell"

type Enemy struct {
	X         int
	Y         int
	Character rune
	Screen    tcell.Screen
}

func NewEnemy(screen tcell.Screen) *Enemy {
	return &Enemy{
		X:         7,
		Y:         6,
		Character: 'e',
		Screen:    screen,
	}
}

func (e *Enemy) Update() {
	e.Screen.SetContent(e.X, e.Y, e.Character, nil, tcell.StyleDefault)
}
