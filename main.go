package main

import (
	"flag"
	"fmt"
	"link/href"
	"log"
	"os"
)

func main() {
	filename := flag.String("f", "ex3.html", "HTML file to parse from")
	flag.Parse()
	f, err := os.Open(*filename)
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(href.Parse(f))

}
