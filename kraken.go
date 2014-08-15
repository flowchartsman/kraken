package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	"log"
)

func main() {
	log.SetFlags(0)
	dataInit()

	log.Println(ansi.ColorCode("blue") + "\n  __\n (oO)\n /||\\\nkraken\n" + ansi.ColorCode("reset"))

	log.Println("reading names into memory")

	for i := 0; i < 45; i++ {
		cc := getCC()

		log.Println(getName(true), cc.CType, cc.Number, cc.CVV2, fmt.Sprintf("%s/%s", cc.expMo(), cc.expY()))
	}

}

// 8===D EG
