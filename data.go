package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	mNameFile   string
	fNameFile   string
	lNameFile   string
	zipCodeFile string
	streetFile  string
)

func init() {

	binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln("kraken can't figure out where it is, so it's gonna go ahead and die now")
	}

	dataPath := filepath.Join(binPath, "data")

	_, err = os.Stat(dataPath)
	if os.IsNotExist(err) {
		log.Println("data directory doesn't exist, creating")
		if err = os.Mkdir(dataPath, os.ModePerm); err != nil {
			log.Fatalf("can't create data directory %s\n", dataPath)
		}
	}
	mNameFile = filepath.Join(dataPath, "mnames.txt")
	fNameFile = filepath.Join(dataPath, "fnames.txt")
	lNameFile = filepath.Join(dataPath, "lnames.txt")
	zipCodeFile = filepath.Join(dataPath, "zipcodes.csv")
	streetFile = filepath.Join(dataPath, "streets.csv")

	downloadIfNot("http://www.census.gov/genealogy/www/data/1990surnames/dist.male.first", mNameFile)
	downloadIfNot("http://www.census.gov/genealogy/www/data/1990surnames/dist.female.first", fNameFile)
	downloadIfNot("http://www.census.gov/genealogy/www/data/1990surnames/dist.all.last", lNameFile)
	downloadIfNot("http://www.unitedstateszipcodes.org/zip_code_database.csv", zipCodeFile)
	downloadIfNot("https://data.cityofchicago.org/api/views/i6bp-fvbx/rows.csv?accessType=DOWNLOAD", streetFile)
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
