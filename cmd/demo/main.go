package main

import (
	"fmt"

	"github.com/leordev/sandbox"
)

func main() {
	fmt.Println(sandbox.Foo("Leo"))
	fmt.Println("2 + 3 =", sandbox.Bar(2, 3))
}
