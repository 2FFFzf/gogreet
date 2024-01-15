package gogreet

import (
	greet "github.com/2FFFzf/gogreet/greet"
)

func Greet(who string) {
	greet.Hello(who)
}

func GreetMom() {
	greet.GreetMom()
}

func GreetMomWithName(who string) {
	greet.GreetMomWithName(who)
}
