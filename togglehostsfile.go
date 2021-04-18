package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// file we're modifying
	name := "test.txt"

	// open original file
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// create temp file
	tmp, err := ioutil.TempFile("", "tmp-text")
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.Close()

	// replace while copying from f to tmp
	if err := replace(f, tmp); err != nil {
		log.Fatal(err)
	}

	// make sure the tmp file was successfully written to
	if err := tmp.Close(); err != nil {
		log.Fatal(err)
	}

	// close the file we're reading from
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	// overwrite the original file with the temp file
	if err := os.Rename(tmp.Name(), name); err != nil {
		log.Fatal(err)
	}

	fmt.Println("We're done over here Captain, check your file for errors")
}

func replace(r io.Reader, w io.Writer) error {

	// domains := []string{"0.0.0.0 google.com","0.0.0.0 youtube.com"}

	// use scanner to read line by line
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		line := sc.Text()
		/*
			for domain := range domains {

			}  */

		if line == "0.0.0.0 google.com" {
			line = "#0.0.0.0 google.com"

		} else if line == "#0.0.0.0 google.com" {
			line = "0.0.0.0 google.com"

		}

		if line == "0.0.0.0 youtube.com" {
			line = "#0.0.0.0 youtube.com"

		} else if line == "#0.0.0.0 youtube.com" {
			line = "0.0.0.0 youtube.com"

		}

		if _, err := io.WriteString(w, line+"\n"); err != nil {
			return err
		}

	}
	return sc.Err()
}
