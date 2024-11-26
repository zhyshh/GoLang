package main

// Component interface
type Coffee interface {
    Cost() float64
    Description() string
}

// ConcreteComponent
type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 {
    return 5.0
}

func (c *SimpleCoffee) Description() string {
    return "Simple coffee"
}