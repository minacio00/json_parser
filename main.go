package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func closeHelper(f *os.File) {
	err := f.Close()
	check(err)
}

func main() {
	// cli
	fPath := flag.String("file", "", "-file=<path>")
	flag.Parse()

	// file
	f, err := os.Open(*fPath)
	check(err)
	defer closeHelper(f)

	// read
	r := bufio.NewScanner(f)
	r.Split(bufio.ScanBytes)
	i := 0
	for r.Scan() {
		s := r.Text()
		// rr := string(s)
		// fmt.Printf("%q \n", s)
		Lex(s)

		i++
	}
	if err := r.Err(); err != nil {
		log.Fatalf("scan error: %v", err)
	}
}
