package main

import (
	"fmt"

	go_say_hello "github.com/rahmat1996/go-say-hello"
	go_say_hello2 "github.com/rahmat1996/go-say-hello/v2"
)

func main() {
	fmt.Println(go_say_hello.SayHello())
	fmt.Println(go_say_hello2.SayHello("Rahmat"))
}
