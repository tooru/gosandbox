package main

//go:generate statik

import (
	"io"
	"log"
	"os"

	"github.com/rakyll/statik/fs"
	_ "github.com/tooru/gosandbox/mingo/example-statik/statik"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	f, err := statikFS.Open("/index.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
}
