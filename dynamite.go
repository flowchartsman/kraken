package main

import (
	"bufio"
	"fmt"
	"github.com/mgutz/ansi"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

type nameSlice []string

func main() {

	binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln("dynamite can't figure out where it is, so it's gonna go ahead and die now")
	}

	dataPath := filepath.Join(binPath, "data")

	_, err = os.Stat(dataPath)
	if os.IsNotExist(err) {
		log.Println("data directory doesn't exist, creating")
		if err = os.Mkdir(dataPath, os.ModePerm); err != nil {
			log.Fatalf("can't create data directory %s\n", dataPath)
		}
	}

	mNameFile := filepath.Join(dataPath, "mnames.txt")
	fNameFile := filepath.Join(dataPath, "fnames.txt")
	lNameFile := filepath.Join(dataPath, "lnames.txt")

	downloadIfNot("http://www.census.gov/genealogy/www/data/1990surnames/dist.male.first", mNameFile)
	downloadIfNot("http://www.census.gov/genealogy/www/data/1990surnames/dist.female.first", fNameFile)
	downloadIfNot("http://www.census.gov/genealogy/www/data/1990surnames/dist.all.last", lNameFile)

	log.Println(ansi.ColorCode("yellow") + "*" + ansi.ColorCode("white") + "--" + ansi.ColorCode("red") + "=====" + ansi.ColorCode("reset"))
	log.Println("dynamite")

	log.Println("reading names into memory")

	fNames := getNameSlice(mNameFile, fNameFile)
	lNames := getNameSlice(mNameFile, lNameFile)

	for i := 0; i < 10; i++ {
		cc := getCC()
		log.Println(fNames.getOne(false), lNames.getOne(false), cc.CType, cc.Number, cc.CVV2, fmt.Sprintf("%s/%s", cc.expMo(), cc.expY()))
	}

}

func downloadIfNot(url string, filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Printf("%s doesn't exist, downloading from %s\n", filename, url)
		out, err := os.Create(filename)
		if err != nil {
			log.Fatalf("can't create %s: %s\n", filename, err)
		}
		defer out.Close()

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("can't download %s: %s\n", url, err)
		}
		defer resp.Body.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Fatalf("can't write %s: %s\n", filename, err)
		}
		log.Println("done")
	}

	return
}

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

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// 8===D EG
