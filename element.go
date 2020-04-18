package main

import (
	"fmt"
	"reflect"

	"github.com/gdamore/tcell"
)

type Vector struct {
	x, y int
}

type Component interface {
	onUpdate(renderer tcell.Screen) error
	onDraw(renderer tcell.Screen) error
}

// Element should be used to build out a more modular design
type Element struct {
	position   Vector
	active     bool
	components []Component
}

func (e *Element) draw(s tcell.Screen) error {
	for _, comp := range e.components {
		err := comp.onDraw(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Element) update(s tcell.Screen) error {
	for _, comp := range e.components {
		err := comp.onUpdate(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Element) addComponent(new Component) {
	for _, existing := range e.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"Tried to add new component with existing type %+v",
				reflect.TypeOf(new)))
		}
	}

	e.components = append(e.components, new)
}

func (e *Element) getComponent(withType Component) Component {
	typ := reflect.TypeOf(withType)
	for _, comp := range e.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}
	// If we could not find the component, fix your code.
	panic(fmt.Sprintf(
		"Could not find component: %+v in element: %+v",
		e, typ))
}
