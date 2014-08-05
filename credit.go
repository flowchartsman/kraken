package main

import (
	"github.com/joeljunstrom/go-luhn"
	//"log"
	"math/rand"
	//"regexp"
	"strconv"
	//"strings"
	"time"
)

type card struct {
	CType  string
	Number string
	CVV2   string
	exp    time.Time
}

func (c card) expMo() string {
	_, mo, _ := c.exp.Date()
	return strconv.Itoa(int(mo))
}

func (c card) expY() string {
	y, _, _ := c.exp.Date()
	return strconv.Itoa(y)
}

type pRange [2]int

type cardDef struct {
	cType    string
	prefixes []pRange
	length   int
}

func (cd *cardDef) getCard() card {
	prefix := cd.prefixes[rand.Intn(len(cd.prefixes))]
	var prefixInt int
	if prefix[1] == 0 {
		prefixInt = prefix[0]
	} else {
		prefixInt = rand.Intn(prefix[1]-prefix[0]) + prefix[0]
	}

	//Expiration should be sometime at least a month out and within the next three years
	exp := time.Now().Add(time.Duration(rand.Intn(3*335*24-31*24)+31*24) * time.Hour)

	return card{
		cd.cType,
		luhn.GenerateWithPrefix(cd.length, strconv.Itoa(prefixInt)),
		strconv.Itoa(rand.Intn(999)),
		exp,
	}
}

var cTable = []cardDef{
	{"American Express", []pRange{{34}, {37}}, 15},
	{"Discover", []pRange{{6011}, {622126, 622925}, {644, 649}, {65}}, 16},
	{"Mastercard", []pRange{{50, 55}}, 16},
	{"Visa", []pRange{{4}}, 16},
}

func getCC() card {
	return cTable[rand.Intn(len(cTable))].getCard()
}
