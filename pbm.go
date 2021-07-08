// Package pbm is a library for manipulating PBM images.
package pbm

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var magicNumASCII = []byte("P1")

// Rotate writes the rotated image of an ASCII encoded PBM to out. The degrees should be a multiple of a quarter turn (e.g. 90, 180, -270, etc.), otherwise it
// returns a non-nil error.
func Rotate(output io.Writer, image io.Reader, degrees int, ccw bool) error {
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
		return errors.New("magic number does not correspond to an ASCII encoded, PBM file")
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

	_, err = fmt.Fprintln(output, string(magicNumASCII))
	if err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}

	_, err = fmt.Fprint(output, comments)
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

		x := (width - 1) - (count / height)
		y := count % height

		matrix[y][x] = b

		count++
	}

	// TODO
	for _, y := range matrix {
		for _, x := range y {
			_, err := fmt.Fprint(output, string(x)+" ")
			if err != nil {
				return fmt.Errorf("could not write to output: %w", err)
			}
		}
		fmt.Fprint(output, "\n")
	}

	return nil
}

// RotateOptimized writes the rotated image of an ASCII encoded PBM to out. The degrees should be a multiple of a quarter turn (e.g. 90, 180, -270, etc.), otherwise it
// returns a non-nil error.
func RotateOptimized(output io.Writer, image io.Reader, degrees int, ccw bool) error {
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
		return errors.New("magic number does not correspond to an ASCII encoded, PBM file")
	}

	var sizeStr string
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			return fmt.Errorf("could not read string: %w", err)
		}
		s = strings.TrimSpace(s)

		if s == "" || strings.HasPrefix(s, "#") {
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

	_, err = fmt.Fprintln(output, string(magicNumASCII))
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

		x := (width - 1) - (count / height)
		y := count % height

		matrix[y][x] = b

		count++
	}

	for _, y := range matrix {
		for _, x := range y {
			_, err := fmt.Fprint(output, string(x))
			if err != nil {
				return fmt.Errorf("could not write to output: %w", err)
			}
		}
	}

	return nil
}
