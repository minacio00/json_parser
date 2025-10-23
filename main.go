package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// cli
	fPath := flag.String("file", "", "-file=<path>")
	flag.Parse()

	// file
	f, err := os.Open(*fPath)
	check(err)
	defer f.Close()

	// read
	r := bufio.NewScanner(f)
	r.Split(bufio.ScanRunes)
	i := 0
	for r.Scan() {
		s := r.Text()
		rr := []rune(s)[0]
		fmt.Printf("%q \n", rr)
		i++
	}
}
