package main

import (
	"github.com/mgutz/ansi"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

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

	downloadIfNot("http://www.census.gov/genealogy/www/data/1990surnames/dist.male.first", filepath.Join(binPath, "data", "mnames.txt"))
	downloadIfNot("http://www.census.gov/genealogy/www/data/1990surnames/dist.female.first", filepath.Join(binPath, "data", "fnames.txt"))

	log.Println(ansi.ColorCode("yellow") + "*" + ansi.ColorCode("white") + "--" + ansi.ColorCode("red") + "=====" + ansi.ColorCode("reset"))
	log.Println("dynamite")
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

// 8===D EG
