package main

import "fmt"

func CatChecker(cats int, name string) (int, int) {
	if cats > 5 {
		fmt.Printf("%s has a lot of cats!\n", name)
		return -1, 4
	}
	fmt.Printf("%d is a tolerable number of cats for %s.\n", cats, name)
	return 5,0
}

func main() {
	a,b := CatChecker(17, "Jeromy")
	fmt.Printf("Function returned %d and %d.\n", a, b)
}
