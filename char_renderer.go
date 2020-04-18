package main

import "github.com/gdamore/tcell"

type CharRenderer struct {
	container *Element
	char      rune
}

func newCharRenderer(e *Element, r rune) *CharRenderer {
	return &CharRenderer{
		container: e,
		char:      r,
	}
}

func (cr *CharRenderer) onDraw(s tcell.Screen) error {
	x := cr.container.position.x
	y := cr.container.position.y

	s.SetContent(x, y, cr.char, nil, tcell.StyleDefault)

	return nil
}

func (cr *CharRenderer) onUpdate(s tcell.Screen) error {
	return nil
}
