package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	"log"
	"math/rand"
	"time"
)

func main() {

	log.Println(ansi.ColorCode("blue") + "\n  __\n (oO)\n /||\\\nkraken\n" + ansi.ColorCode("reset"))
	log.Println("kraken")

	log.Println("reading names into memory")

	fNames := getNameSlice(mNameFile, fNameFile)
	lNames := getNameSlice(lNameFile)

	for i := 0; i < 50; i++ {
		cc := getCC()
		log.Println(fNames.getOne(true), lNames.getOne(true), cc.CType, cc.Number, cc.CVV2, fmt.Sprintf("%s/%s", cc.expMo(), cc.expY()))
	}

}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// 8===D EG
