package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	// Create string flag
	s := parser.String("s", "string", &argparse.Options{Required: true, Help: "String to print"})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	// Finally print the collected string
	fmt.Println(*s)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read(path string) string {
	if path[len(path)-5:] == "fasta" {
		rawData, err := os.ReadFile(path)
		check(err)
		return string(rawData)
	}
	panic("wrong file format")
}
