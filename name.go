package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var (
	fNames stringSlice
	mNames stringSlice
	lNames stringSlice
)

//TODO: Godoc comments
func getNameSlice(file string) stringSlice {

	fh, err := os.Open(file)
	defer fh.Close()
	if err != nil {
		log.Fatalf("can't open file %s: %s", file, err)
	}

	scan := csv.NewReader(fh)

	/* Find the lowest number and determine the generic multiplier such that it
	   occurs once. We know that the namefiles are sorted, but in case we ever need
	   to use another source... */

	lowest_prob := 100.0

	rawnames, err := scan.ReadAll()
	if err != nil {
		log.Fatalln("Error reading names file %s: %s", file, err)
	}

	for _, name := range rawnames {
		prob, err := strconv.ParseFloat(name[1], 32)
		if err != nil {
			log.Fatalln("Malformed input in names file %s: %s", file, err)
		}
		if prob > 0 && prob < lowest_prob {
			lowest_prob = prob
		}
	}

	mult := 1 / lowest_prob

	//Determine the size of the slice by adding all of the multiplied weights together
	size := 0
	for _, name := range rawnames {
		prob, _ := strconv.ParseFloat(name[1], 32)
		prob *= mult
		if prob == 0 {
			continue
		}
		size += int(math.Floor(prob + 0.5))
	}

	out := make([]string, size)

	/* Finally, expand each entry into the slice.
	   Note: if there is another algorithm that handles multiples of the same weight,
	   it should probably be used here unless it impacts performance. Thusfar, all approaches
	   decorate and then bisect a list, and this is inadequate in this case */

	idx := 0
	for _, name := range rawnames {
		prob, _ := strconv.ParseFloat(name[1], 32)
		prob *= mult
		if prob == 0 {
			continue
		}
		entries := int(math.Floor(prob + 0.5))
		for i := 0; i < entries; i++ {
			out[idx] = name[0]
			idx++
		}
	}

	return stringSlice(out)
}

func getName(capitalize bool) string {
	if maybe() {
		return fmt.Sprintf("%s %s", fNames.getOne(capitalize), lNames.getOne(capitalize))
	} else {
		return fmt.Sprintf("%s %s", mNames.getOne(capitalize), lNames.getOne(capitalize))
	}
}
