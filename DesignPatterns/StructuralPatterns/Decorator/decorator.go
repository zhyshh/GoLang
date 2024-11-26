// Decorator
package main

type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) Cost() float64 {

	return d.coffee.Cost()

}

func (d *CoffeeDecorator) Description() string {

	return d.coffee.Description()

}

// ConcreteDecorators

type MilkDecorator struct {
	*CoffeeDecorator
}

func NewMilkDecorator(c Coffee) *MilkDecorator {

	return &MilkDecorator{&CoffeeDecorator{c}}

}

func (d *MilkDecorator) Cost() float64 {

	return d.CoffeeDecorator.Cost() + 1.0

}

func (d *MilkDecorator) Description() string {

	return d.CoffeeDecorator.Description() + ", milk"

}

type SugarDecorator struct {
	*CoffeeDecorator
}

func NewSugarDecorator(c Coffee) *SugarDecorator {

	return &SugarDecorator{&CoffeeDecorator{c}}

}

func (d *SugarDecorator) Cost() float64 {

	return d.CoffeeDecorator.Cost() + 0.5

}

func (d *SugarDecorator) Description() string {

	return d.CoffeeDecorator.Description() + ", sugar"

}
