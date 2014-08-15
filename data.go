package main

import (
	"bufio"
	//"compress/gzip"
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

type stringSlice []string

func (ss stringSlice) getOne(capitalize bool) string {
	idx := rand.Intn(len(ss))
	if capitalize {
		return ss[idx]
	} else {
		return strings.ToLower(ss[idx])
	}
}

var (
	mNameFile    string
	fNameFile    string
	lNameFile    string
	zipCodesFile string
	streetFile   string
)

/* TODO: Either run this with a flag and quit or run this in lieu of files then quit.
There's no need to yank down a bunch of data, put it in memory, process it and THEN
run. */

func dataInit() {
	binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln("kraken cannot determine the location of its binary; terminating")
	}

	dataPath := filepath.Join(binPath, "data")

	_, err = os.Stat(dataPath)
	if os.IsNotExist(err) {
		log.Println("data directory doesn't exist, creating")
		if err = os.Mkdir(dataPath, os.ModePerm); err != nil {
			log.Fatalf("can't create data directory %s\n", dataPath)
		}
	}
	mNameFile = filepath.Join(dataPath, "mnames")
	fNameFile = filepath.Join(dataPath, "fnames")
	lNameFile = filepath.Join(dataPath, "lnames")
	zipCodesFile = filepath.Join(dataPath, "zipcodes")
	streetFile = filepath.Join(dataPath, "streets")

	if noFile(mNameFile) {
		downloadAndProcessNamefile("http://www.census.gov/genealogy/www/data/1990surnames/dist.male.first", mNameFile)
	}
	if noFile(fNameFile) {
		downloadAndProcessNamefile("http://www.census.gov/genealogy/www/data/1990surnames/dist.female.first", fNameFile)
	}
	if noFile(lNameFile) {
		downloadAndProcessNamefile("http://www.census.gov/genealogy/www/data/1990surnames/dist.all.last", lNameFile)
	}

	if noFile(zipCodesFile) {
		downloadAndProcessZipCodesfile("http://www.unitedstateszipcodes.org/zip_code_database.csv", zipCodesFile)
	}

	if noFile(streetFile) {
		downloadAndProcessStreetfile("https://data.cityofchicago.org/api/views/i6bp-fvbx/rows.csv?accessType=DOWNLOAD", streetFile)
	}

	fNames = getNameSlice(fNameFile)
	mNames = getNameSlice(mNameFile)
	lNames = getNameSlice(lNameFile)
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

func noFile(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return true
		} else {
			log.Fatalf("failed to check for filename '%s': %s", filename, err)
		}
	}
	return false
}

//TODO: Write to tempfile first and then mv file to final location
func downloadAndProcessNamefile(url string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file '%s': %s", filename, err)
	}
	defer file.Close()
	//gz_out := gzip.NewWriter(file)
	//defer gz_out.Close()
	csv_out := csv.NewWriter(file)
	defer csv_out.Flush()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("can't download %s: %s\n", url, err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		csv_out.Write(strings.FieldsFunc(toName(scanner.Text()), unicode.IsSpace)[0:2])
	}

	log.Println("done")
}

//TODO Can probably generify this somewhat
func downloadAndProcessZipCodesfile(url string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file '%s': %s", filename, err)
	}
	defer file.Close()
	//gz_out := gzip.NewWriter(file)
	//defer gz_out.Close()
	csv_out := csv.NewWriter(file)
	defer csv_out.Flush()

	log.Println("Downloading zip codes (this may take a while)")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("can't download %s: %s", url, err)
	}
	defer resp.Body.Close()

	input := csv.NewReader(resp.Body)

	for {
		zipcode, err := input.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Error parsing %s: %s", url, err)
		}

		// Don't care about PO Boxes (yet) or nonstandard zips.
		if zipcode[1] != "STANDARD" {
			continue
		}
		// Only take codes of estimated population of >5,000
		if len(zipcode) >= 15 {
			pop, _ := strconv.Atoi(zipcode[14])
			if pop >= 5000 {
				csv_out.Write([]string{zipcode[0], toName(zipcode[2]), zipcode[5]})
			}
		}
	}
	log.Println("done")
	return
}

func downloadAndProcessStreetfile(url string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file '%s': %s", filename, err)
	}
	defer file.Close()
	//gz_out := gzip.NewWriter(file)
	//defer gz_out.Close()
	csv_out := csv.NewWriter(file)
	defer csv_out.Flush()

	log.Println("Downloading zip codes (this may take a while)")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("can't download %s: %s", url, err)
	}
	defer resp.Body.Close()

	input := csv.NewReader(resp.Body)

	//Skip the first line
	_, _ = input.Read()
	for {
		streets, err := input.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Error parsing %s: %s", url, err)
		}

		csv_out.Write([]string{toName(streets[2] + " " + streets[3]), streets[5], streets[6]})
	}
	log.Println("done")
	return
}

func toName(s string) string {
	rs := []rune(s)
	last := rune(' ')

	for i, r := range rs {
		if unicode.IsLetter(r) {
			if unicode.IsSpace(last) {
				rs[i] = unicode.ToUpper(rs[i])
			} else {
				rs[i] = unicode.ToLower(rs[i])
			}
		}
		last = rs[i]
	}

	return string(rs)
}
