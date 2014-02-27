package main
import "fmt"

type Animal interface { Speak() }

type Cougar struct {}
func (c *Cougar) Speak() {fmt.Println("Go Cougs!") }

type Husky struct {}
func (d *Husky) Speak() { fmt.Println("Huskies Suck!") }

func main() {
	var a Animal
	a = new(Cougar)
	a.Speak()
	a = new(Husky)
	a.Speak()
}
