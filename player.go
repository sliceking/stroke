package main

import "github.com/gdamore/tcell"

func NewPlayer(screen tcell.Screen) *Element {
	player := Element{}
	player.position = Vector{x: 5, y: 5}
	player.active = true

	renderer := newCharRenderer(&player, '@')
	player.addComponent(renderer)

	mover := newKeyboardMover(&player)
	player.addComponent(mover)

	return &player
}
