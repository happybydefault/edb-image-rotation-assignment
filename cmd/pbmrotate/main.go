package main

import (
	"bufio"
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

		degrees    int
		ccw        bool
		resultName string
	)

	flagSet := flag.NewFlagSet("pbmrotate", flag.ExitOnError)

	flagSet.BoolVar(&printHelp, "h", false, "Print this help text")

	flagSet.IntVar(&degrees, "d", 90, "Number of degrees. Possible values are only 90, 180, and 270")
	flagSet.BoolVar(&ccw, "c", false, "Counterclockwise")
	flagSet.StringVar(&resultName, "o", "", "Write the result to file instead of stdout")

	flagSet.Usage = func() {
		fmt.Fprintf(flagSet.Output(), "\nUsage: %s [OPTIONS] FILE\n\nOptions:\n\n", flagSet.Name())
		flagSet.PrintDefaults()
		fmt.Fprint(flagSet.Output(), "\n")
		printExamples(flagSet.Output())
	}

	// On error, it executes flag set's Usage() and exits, because of flag.ExitOnError
	flagSet.Parse(os.Args[1:])

	if printHelp {
		flagSet.SetOutput(os.Stdout)
		flagSet.Usage()
		os.Exit(0)
	}

	inputName := flagSet.Arg(0)

	err := run(inputName, resultName, degrees, ccw)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(inputName string, resultName string, degrees int, ccw bool) error {
	var (
		r io.Reader
		w io.Writer
	)

	stat, err := os.Stdin.Stat()
	if err != nil {
		return fmt.Errorf("could not get stats from stdin: %w", err)
	}

	pipe := stat.Mode()&os.ModeCharDevice != os.ModeCharDevice
	if pipe {
		r = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(inputName)
		if err != nil {
			return fmt.Errorf("could not open file %q: %s", inputName, err)
		}
		defer f.Close()

		r = bufio.NewReader(f)
	}

	if resultName == "" {
		buf := bufio.NewWriter(os.Stdout)
		defer buf.Flush()
		w = buf
	} else {
		f, err := os.Create(resultName)
		if err != nil {
			return fmt.Errorf("could not copy to file %q: %w", resultName, err)
		}
		defer f.Close()

		buf := bufio.NewWriter(f)
		defer buf.Flush()
		w = buf
	}

	err = pbm.Rotate(w, r, degrees, ccw)
	if err != nil {
		return fmt.Errorf("could not rotate image: %w", err)
	}

	return nil
}

func printExamples(w io.Writer) error {
	examples := `# Rotate an image 270 degrees clockwise and write the result to a file
pbmrotate -d=270 -o="example-image-rotated.pbm" example-image.pbm

# Rotate an image 90 degrees counterclockwise and write the result to stdout
pbmrotate -d=90 -c example-image.pbm

# Rotate an image 180 degrees from stdin and write the result to a file
curl "https://example.com/internet-image.pbm" | pbmrotate -d=180 -o="internet-image-rotated.pbm"`

	_, err := fmt.Fprintf(w, "Examples:\n\n%+20v\n", examples)
	return err
}
