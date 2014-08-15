package main

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	//No need for crypto/rand (yet)
	rand.Seed(time.Now().UTC().UnixNano())
}

//generate a random integer between min and max
func randRange(min int, max int) int {
	max += 1
	if min >= max {
		log.Panicln("min cannot be greater than max")
	}
	return (rand.Intn(max-min) + min)
}

//maybe is good for quick randomness with or without probability.
//without any arguments, it is a simple 50/50 true/false. With a single integer
//argument, it will have that probability of returning true
func maybe(prob ...int) bool {
	if len(prob) > 0 {
		if rand.Intn(100) <= prob[0]-1 {
			return true
		}
		return false
	} else {
		if rand.Intn(2) == 1 {
			return true
		}
		return false
	}
}
