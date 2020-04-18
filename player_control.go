package main

import "github.com/gdamore/tcell"

type KeyboardMover struct {
	container *Element
	cr        *CharRenderer
}

func newKeyboardMover(e *Element) *KeyboardMover {
	km := KeyboardMover{
		container: e,
		cr:        e.getComponent(&CharRenderer{}).(*CharRenderer),
	}

	return &km
}

func (mover *KeyboardMover) onUpdate(screen tcell.Screen) error {
	// somehow get the keyboard events into here and then update the location
	cont := mover.container
	switch ev := screen.PollEvent().(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyLeft:
			if r, _, _, _ := screen.GetContent(cont.position.x-1, cont.position.y); r != ' ' {
				return nil
			}
			cont.position.x += -1

		case tcell.KeyRight:
			if r, _, _, _ := screen.GetContent(cont.position.x+1, cont.position.y); r != ' ' {
				return nil
			}
			cont.position.x += 1

		case tcell.KeyUp:
			if r, _, _, _ := screen.GetContent(cont.position.x, cont.position.y-1); r != ' ' {
				return nil
			}
			cont.position.y += -1

		case tcell.KeyDown:
			if r, _, _, _ := screen.GetContent(cont.position.x, cont.position.y+1); r != ' ' {
				return nil
			}
			cont.position.y += 1
		}
	}

	return nil
}

func (mover *KeyboardMover) onDraw(s tcell.Screen) error {
	return nil
}
