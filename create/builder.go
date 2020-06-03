package create

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH Speed = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports" //赛车轮
	SteelWheels         = "steel"
)

type CarBuilder interface {
	Color(Color) CarBuilder
	Wheels(Wheels) CarBuilder
	TopSpeed(Speed) CarBuilder
	Build() Driver
}

type Driver interface {
	Drive() error
	Stop() error
}

type Car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

func (c *Car) Drive() error {
	fmt.Printf("color=%s,wheels=%s,speed=%f\n", c.color, c.wheels, c.speed)
	return nil
}

func (c *Car) Stop() error {
	c.speed = 0.0
	fmt.Printf("color=%s,wheels=%s,speed=%f\n", c.color, c.wheels, c.speed)
	return nil
}

func (c *Car) Color(color Color) CarBuilder {
	c.color = color
	return c
}

func (c *Car) Wheels(wheels Wheels) CarBuilder {
	c.wheels = wheels
	return c
}

func (c *Car) TopSpeed(speed Speed) CarBuilder {
	c.speed = speed
	return c
}

func (c *Car) Build() Driver {
	return c
}

func NewCarBuilder() CarBuilder {
	return &Car{}
}
