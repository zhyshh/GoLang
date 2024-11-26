package main
import "fmt"

func main() {

    var coffee Coffee = &SimpleCoffee{}

    fmt.Println(coffee.Description(), ":", coffee.Cost()) // Simple coffee : 5


    coffee = NewMilkDecorator(coffee)

    fmt.Println(coffee.Description(), ":", coffee.Cost()) // Simple coffee, milk : 6


    coffee = NewSugarDecorator(coffee)

    fmt.Println(coffee.Description(), ":", coffee.Cost()) // Simple coffee, milk, sugar : 6.5

}