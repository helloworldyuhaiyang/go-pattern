package create

import (
	"errors"
	"fmt"
)

type Speed float64

const (
	MPH Speed = 1
	KPH Speed = 1.60934
)

type Color string

const (
	NoColor    Color = ""
	BlueColor        = "blue"
	GreenColor       = "green"
	RedColor         = "red"
	BlackColor       = "black"
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
	Build() (Driver, error)
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

func (c *Car) Build() (Driver, error) {
	// 判断必填参数，在参数过多时候，防止用户少传参数
	if c.color == NoColor {
		return nil, errors.New("color is empty")
	}
	return c, nil
}

func NewCarBuilder() CarBuilder {
	return &Car{color: NoColor}
}
