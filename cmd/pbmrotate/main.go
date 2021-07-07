package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	pbm "github.com/happybydefault/edb-image-rotation-assignment"
)

func main() {
	var (
		printHelp bool

		degrees     int
		ccw         bool
		filenameOut string
	)

	flagSet := flag.NewFlagSet("pbmrotate", flag.ExitOnError)

	flagSet.BoolVar(&printHelp, "h", false, "Print help")

	flagSet.IntVar(&degrees, "d", 90, "Number of degrees. Possible values are only 90, 180, and 270")
	flagSet.BoolVar(&ccw, "c", false, "Counterclockwise")
	flagSet.StringVar(&filenameOut, "o", "", "Write to file instead of stdout")

	flagSet.Usage = func() {
		fmt.Fprintf(flagSet.Output(), "\nUsage: %s [OPTIONS] FILE\n\nOptions:\n", flagSet.Name())
		flagSet.PrintDefaults()
	}

	// on error, it executes flag set Usage() and exists (because of flag.ExitOnError)
	flagSet.Parse(os.Args[1:])

	if printHelp {
		flagSet.SetOutput(os.Stdout)
		flagSet.Usage()
		os.Exit(0)
	}

	filenameIn := flagSet.Arg(0)

	err := run(filenameIn, filenameOut, degrees, ccw)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(filenameIn string, filenameOut string, degrees int, ccw bool) error {
	var (
		fileIn  io.Reader
		fileOut io.Writer
	)

	stdinInfo, err := os.Stdin.Stat()
	if err != nil {
		return fmt.Errorf("could not get stats from stdin: %w", err)
	}

	pipe := stdinInfo.Mode()&os.ModeNamedPipe == os.ModeNamedPipe
	if pipe {
		fileIn = os.Stdin
	} else if filenameIn == "" || filenameIn == "-" {
		buf := &bytes.Buffer{}
		_, err := buf.ReadFrom(os.Stdin)
		if err != nil {
			return fmt.Errorf("could not read from stdin: %w", err)
		}
		fileIn = buf
	} else {
		f, err := os.Open(filenameIn)
		if err != nil {
			log.Printf("could not open file %q: %s", filenameIn, err)
			os.Exit(1)
		}
		defer f.Close()
		fileIn = f
	}

	if filenameOut == "" {
		fileOut = os.Stdout
	} else {
		f, err := os.Create(filenameOut)
		if err != nil {
			return fmt.Errorf("could not copy to file %q: %w", filenameOut, err)
		}
		defer f.Close()
		fileOut = f
	}

	err = pbm.Rotate(fileOut, fileIn, degrees, ccw)
	if err != nil {
		return fmt.Errorf("could not rotate image: %w", err)
	}

	return nil
}
