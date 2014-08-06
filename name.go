package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func getNameSlice(files ...string) nameSlice {
	var lines int64

	for _, file := range files {
		fh, err := os.Open(file)
		if err != nil {
			log.Fatalf("can't open file %s: %s", file, err)
		}

		scanner := bufio.NewScanner(fh)
		for scanner.Scan() {
			lines++
		}
		fh.Close()
	}

	out := make([]string, lines)

	var idx int64

	for _, file := range files {
		fh, err := os.Open(file)
		if err != nil {
			log.Fatalf("can't open file %s: %s", file, err)
		}

		scanner := bufio.NewScanner(fh)
		for scanner.Scan() {
			out[idx] = strings.ToLower(strings.Split(scanner.Text(), " ")[0])
			idx++
		}
		fh.Close()
	}
	return nameSlice(out)
}

type nameSlice []string

func (ns nameSlice) getOne(capitalize bool) string {
	idx := rand.Intn(len(ns))
	if capitalize {
		return upperFirst(ns[idx])
	} else {
		return ns[idx]
	}
}

func upperFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}
