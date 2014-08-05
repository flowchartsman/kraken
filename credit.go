package main

import (
	"github.com/joeljunstrom/go-luhn"
	//"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

var ccPrefixes = []string{
	//American Express
	"34",
	"37",
	//Discover
	"6011",
	"622126-622925",
	"644-649",
	"65",
	//MasterCard
	"50-55",
	//Visa
	"4",
}

var hasDash = regexp.MustCompile("-")

func getCCNum() string {
	prefix := ccPrefixes[rand.Intn(len(ccPrefixes))]
	if hasDash.MatchString(prefix) {
		prefixRange := strings.Split(prefix, "-")
		prefixMin, _ := strconv.Atoi(prefixRange[0])
		prefixMax, _ := strconv.Atoi(prefixRange[1])
		prefixInt := rand.Intn(prefixMax-prefixMin) + prefixMin
		return luhn.GenerateWithPrefix(12, strconv.Itoa(prefixInt))
	} else {
		return luhn.GenerateWithPrefix(12, prefix)
	}
}
