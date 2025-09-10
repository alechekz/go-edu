package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// dir is a custom type for certain directory in filesystem
type dir struct {
	fullname string
	files    []os.DirEntry
}

// init flag
var displayFiles bool

// newDir returns new directory *dir
func newDir(name string) *dir {
	return &dir{
		fullname: name,
	}
}

// read reads directory files
func (d *dir) read() {
	files, err := os.ReadDir(d.fullname)
	if err != nil {
		log.Fatalln(err)
	}
	d.files = files
}

// display prints directory files
func (d *dir) display() {

	//nothing to print exception
	if len(d.files) == 0 {
		return
	}

	//print files
	for _, file := range d.files {
		if file.Name()[0] == '.' {
			continue
		}
		name := filepath.Join(d.fullname, file.Name())

		//continue with next directory
		if file.IsDir() {
			fmt.Printf("%v\n", name)
			nextD := newDir(name)
			nextD.read()
			nextD.display()
		}
		if displayFiles {
			fmt.Printf("%v\n", name)
		}
	}
}

// init
func init() {
	flag.BoolVar(
		&displayFiles,
		"f",
		false,
		"display files if the -f option is specified",
	)
}

// main
func main() {

	flag.Parse()

	//get dir name
	name := "."
	if len(os.Args) > 1 && os.Args[1] != "-f" {
		name = os.Args[1]
	} else if len(os.Args) > 2 {
		name = os.Args[2]
	}

	//get new dir
	dir := newDir(name)
	dir.read()
	dir.display()
}
