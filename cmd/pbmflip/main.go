package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	pbm "github.com/happybydefault/edb-image-rotation-assignment"
)

func main() {
	f := flag.NewFlagSet("pbmflip", flag.ExitOnError)

	var printHelp bool
	var ccw bool
	var degrees int

	f.BoolVar(&printHelp, "h", false, "Print help")
	f.IntVar(&degrees, "d", 90, "Number of degrees")
	f.BoolVar(&ccw, "r", false, "Counterclockwise")

	f.Usage = func() {
		fmt.Fprintf(f.Output(), "\nUsage: %s [OPTIONS] FILE\n\nOptions:\n", f.Name())
		f.PrintDefaults()
	}

	// on error, it executes flag set Usage() and exists (because of flag.ExitOnError)
	f.Parse(os.Args[1:])

	if printHelp {
		f.SetOutput(os.Stdout)
		f.Usage()
		return
	}

	filename := f.Arg(0)
	if filename == "" {
		f.Usage()
		os.Exit(2)
	}

	err := run(filename, degrees, ccw)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(filename string, degrees int, ccw bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file %q: %w", filename, err)
	}

	result, err := pbm.Flip(file, degrees, ccw)
	if err != nil {
		return fmt.Errorf("could flip image: %w", err)
	}

	direction := "cw"
	if ccw {
		direction = "ccw"
	}
	name := fmt.Sprintf("%s-%s%d.pbm", strings.TrimSuffix(filename, ".pbm"), direction, degrees)

	err = os.WriteFile(name, result, 0644)
	if err != nil {
		return fmt.Errorf("could not write to file %q: %w", name, err)
	}
	log.Printf("created file %q", name)

	return nil
}
