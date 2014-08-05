package main

import (
	"fmt"
	"github.com/mgutz/ansi"
)

//D YN AMITE
func main() {
	fmt.Println(ansi.ColorCode("yellow") + "*" + ansi.ColorCode("white") + "--" + ansi.ColorCode("red") + "=====" + ansi.ColorCode("reset"))
	fmt.Println("dynamite")
}
