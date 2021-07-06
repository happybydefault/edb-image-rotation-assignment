// Package pbm is a library for manipulating PBM images.
package pbm

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

var magicNumASCII = []byte("P1")

// Flip writes the rotated image of an ASCII encoded PBM to out. The degrees should be a multiple of a quarter turn (e.g. 90, 180, -270, etc.), otherwise it
// returns a non-nil error.
func Flip(output io.Writer, image io.Reader, degrees int, ccw bool) error {
	if image == nil {
		return errors.New("image is nil")
	}

	quarterTurn := 90
	if degrees%quarterTurn > 0 {
		return errors.New("number of degrees is not multiple of a quarter turn")
	}

	r := bufio.NewReader(image)
	magicNum := make([]byte, 2)
	_, err := r.Read(magicNum)
	if err != nil {
		return fmt.Errorf("could not read magic number: %w", err)
	}
	if !bytes.Equal(magicNum, magicNumASCII) {
		return fmt.Errorf("magic number does not correspond to an ASCII PBM file: %w", err)
	}

	var (
		sizeStr  string
		comments = &bytes.Buffer{}
	)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			return fmt.Errorf("could not read string: %w", err)
		}
		s = strings.TrimSpace(s)

		if s == "" {
			continue
		} else if strings.HasPrefix(s, "#") {
			fmt.Fprintln(comments, s)
			continue
		} else {
			sizeStr = s
			break
		}
	}

	size := strings.Split(sizeStr, " ")
	if len(size) < 2 {
		return errors.New("invalid size string")
	}

	originalWidth, err := strconv.Atoi(size[0])
	if err != nil {
		return fmt.Errorf("invalid width: %w", err)
	}
	originalHeight, err := strconv.Atoi(size[1])
	if err != nil {
		return fmt.Errorf("invalid height: %w", err)
	}
	log.Printf("width: %d, height: %d", originalWidth, originalHeight)

	// TODO

	_, err = fmt.Fprintln(output, string(magicNumASCII))
	if err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}

	_, err = fmt.Fprint(output, comments)
	if err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}

	_, err = fmt.Fprintln(output, "# Flipped")
	if err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}

	_, err = fmt.Fprintln(output, originalHeight, originalWidth)
	if err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}

	width, height := originalHeight, originalWidth
	matrix := make([][]byte, height)
	for i := range matrix {
		matrix[i] = make([]byte, width)
	}

	var count int
	for {
		b, err := r.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("could not read byte: %w", err)
		}

		space := false
		switch b {
		case ' ', '\t', '\n', '\r':
			space = true
		}
		if space {
			continue
		}

		// TODO
		originalX := count % originalWidth
		originalY := count / originalWidth
		log.Println(originalX, originalY)

		count++
	}
	log.Printf("%+v", matrix)

	return nil
}
