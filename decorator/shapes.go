package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

type Square struct {
	Side float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// Now If I want a colored circle, I could break OCD and add color attribute to circle struct
// or I could create a struct ColoredCircle aggregating Circle,
// that would work fine, but if I have too many shapes like square, rectangle,etc. then writing
// Colored struct for each would not be a good design. Thus I can create a generic struct ColoredShape
// which implements shape, has Shape field and has a color field. This uses the decorator pattern,
// Shape is decorated
type ColoredShape struct {
	Shape Shape
	Color string
}

func (cs *ColoredShape) Render() string {
	return fmt.Sprintf("%s has color %s", cs.Shape.Render(), cs.Color)
}

// Extending behavior furthur, adding transparency, this is a decorated object
type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (ts *TransparentShape) Render() string {
	return fmt.Sprintf("%s has transparency %f%%", ts.Shape.Render(), ts.Transparency)
}
func main() {
	c := Circle{}
	c.Radius = 10
	// Not using Square here as first arg because *Square implements Shape interface, not Square (value)
	cs := ColoredShape{&Square{Side: 10}, "Red"}
	fmt.Println(cs.Render())
	fmt.Println(c.Render())
	// Decorated Objects can be composed because they all implement Shape interface
	tc := TransparentShape{&cs, 10}
	fmt.Println(tc.Render())
}
