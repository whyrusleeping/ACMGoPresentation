package main
import "fmt"

type Animal interface { Speak() }

type Dog struct {}
func (d *Dog) Speak() { fmt.Println("Woof Woof!") }

type Cat struct {}
func (c *Cat) Speak() {fmt.Println("Meaaoow!") }

func main() {
	var a Animal
	a = new(Dog)
	a.Speak()
	a = new(Cat)
	a.Speak()
}
