// Package pbm is a library for manipulating PBM images.
package pbm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

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

	sizeStr, err := r.ReadString('\n')
	if err != nil {
		return fmt.Errorf("could not read string: %w", err)
	}
	sizeStr = strings.TrimSuffix(sizeStr, "\n")

	size := strings.Split(sizeStr, " ")
	if len(size) < 2 {
		return errors.New("invalid size string")
	}

	width, err := strconv.Atoi(size[0])
	if err != nil {
		return fmt.Errorf("invalid width: %w", err)
	}

	height, err := strconv.Atoi(size[1])
	if err != nil {
		return fmt.Errorf("invalid height: %w", err)
	}

	log.Printf("width: %d, height: %d", width, height)

	// TODO
	_, err = io.Copy(output, image)
	if err != nil {
		return fmt.Errorf("could not copy: %w", err)
	}
	_, err = fmt.Fprintln(output, "# Flipped")
	if err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}

	return nil
}
