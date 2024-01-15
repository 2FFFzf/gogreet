package greet

import (
	"fmt"
)

func GreetMom() {
	fmt.Println("Hello mom I'm learing golang module")
}

func GreetMomWithName(who string) {
	fmt.Printf("Hello my lovely mom %s, I'm learning golang module!\n", who)
}
