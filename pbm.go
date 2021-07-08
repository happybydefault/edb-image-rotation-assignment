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
	"unicode"
)

var magicNumASCII = []byte("P1")

// Rotate writes the rotated image of an ASCII encoded PBM to out. The number of degrees should be a multiple of a quarter turn (e.g. 90, 180, -270, etc.), otherwise it
// returns a non-nil error.
func Rotate(output io.Writer, image io.Reader, degrees int, ccw bool) error {
	if output == nil {
		return errors.New("output is nil")
	}

	if image == nil {
		return errors.New("image is nil")
	}

	if degrees%90 != 0 {
		return errors.New("number of degrees should be a multiple of a quarter turn")
	}

	if ccw {
		degrees *= -1
	}

	rotations := (degrees/90%4 + 4) % 4

	if rotations == 0 {
		_, err := io.Copy(output, image)
		if err != nil {
			return fmt.Errorf("could not copy: %w", err)
		}
		return nil
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

	var width, height int
loop:
	for {
		b, err := r.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("could not read byte: %w", err)
		}

		if unicode.IsSpace(rune(b)) {
			continue
		} else if b == '#' {
			_, err := r.ReadBytes('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					break loop
				}
				return fmt.Errorf("could not read bytes: %w", err)
			}

			continue
		}

		err = r.UnreadByte()
		if err != nil {
			return fmt.Errorf("could not unread byte: %w", err)
		}

		sizeStr, err := r.ReadString('\n')
		if err != nil {
			return fmt.Errorf("could not read string: %w", err)
		}
		sizeStr = strings.TrimSpace(sizeStr)

		size := strings.Split(sizeStr, " ")
		if len(size) < 2 {
			return errors.New("invalid width or height")
		}
		widthStr := size[0]
		heightStr := size[1]

		widthN, err := strconv.Atoi(widthStr)
		if err != nil {
			return fmt.Errorf("invalid width: %w", err)
		}
		width = widthN

		heightN, err := strconv.Atoi(heightStr)
		if err != nil {
			return fmt.Errorf("invalid height: %w", err)
		}
		height = heightN

		break
	}

	w := bufio.NewWriter(output)
	defer w.Flush()

	_, err = fmt.Fprintln(w, string(magicNumASCII))
	if err != nil {
		return fmt.Errorf("could not write: %w", err)
	}

	switch rotations {
	case 1:
		_, err = fmt.Fprintln(w, height, width)
		if err != nil {
			return fmt.Errorf("could not write: %w", err)
		}

		err = rotate90(r, w, width, height)
		if err != nil {
			return fmt.Errorf("could not rotate 90 degrees: %w", err)
		}
	case 2:
		_, err = fmt.Fprintln(w, width, height)
		if err != nil {
			return fmt.Errorf("could not write: %w", err)
		}

		err = rotate180(r, w, width, height)
		if err != nil {
			return fmt.Errorf("could not rotate 180 degrees: %w", err)
		}
	case 3:
		_, err = fmt.Fprintln(w, height, width)
		if err != nil {
			return fmt.Errorf("could not write: %w", err)
		}

		err = rotate270(r, w, width, height)
		if err != nil {
			return fmt.Errorf("could not rotate 270 degrees: %w", err)
		}
	}

	_, err = fmt.Fprint(w, "\n")
	if err != nil {
		return fmt.Errorf("could not write: %w", err)
	}

	return nil
}

func rotate90(r io.Reader, w io.Writer, width, height int) error {
	width, height = height, width

	matrix := make([][]string, height)
	for i := range matrix {
		matrix[i] = make([]string, width)
	}

	rd := bufio.NewReader(r)
	var count int
	for {
		b, err := rd.ReadByte()
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

		if y+1 > len(matrix) {
			return errors.New("invalid pixel data")
		}
		if x+1 > len(matrix[y]) {
			return errors.New("invalid pixel data")
		}
		matrix[y][x] = string(b)

		count++
	}

	buf := bufio.NewWriter(w)
	defer buf.Flush()

	lines := make([]string, len(matrix))
	for i, v := range matrix {
		lines[i] = strings.Join(v, " ")
	}
	_, err := buf.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		return fmt.Errorf("could not write string: %w", err)
	}

	return nil
}

func rotate180(r io.Reader, w io.Writer, width, height int) error {
	matrix := make([][]string, height)
	for i := range matrix {
		matrix[i] = make([]string, width)
	}

	rd := bufio.NewReader(r)
	var count int
	for {
		b, err := rd.ReadByte()
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

		x := (width - 1) - count%width
		y := (height - 1) - (count / width)

		if y+1 > len(matrix) {
			return errors.New("invalid pixel data")
		}
		if x+1 > len(matrix[y]) {
			return errors.New("invalid pixel data")
		}
		matrix[y][x] = string(b)

		count++
	}

	buf := bufio.NewWriter(w)
	defer buf.Flush()

	lines := make([]string, len(matrix))
	for i, v := range matrix {
		lines[i] = strings.Join(v, " ")
	}
	_, err := buf.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		return fmt.Errorf("could not write string: %w", err)
	}

	return nil
}

func rotate270(r io.Reader, w io.Writer, width, height int) error {
	width, height = height, width

	matrix := make([][]string, height)
	for i := range matrix {
		matrix[i] = make([]string, width)
	}

	rd := bufio.NewReader(r)
	var count int
	for {
		b, err := rd.ReadByte()
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

		x := count / height
		y := (height - 1) - count%height

		if y+1 > len(matrix) {
			return errors.New("invalid pixel data")
		}
		if x+1 > len(matrix[y]) {
			return errors.New("invalid pixel data")
		}
		matrix[y][x] = string(b)

		count++
	}

	buf := bufio.NewWriter(w)
	defer buf.Flush()

	lines := make([]string, len(matrix))
	for i, v := range matrix {
		lines[i] = strings.Join(v, " ")
	}
	_, err := buf.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		return fmt.Errorf("could not write string: %w", err)
	}

	return nil
}
