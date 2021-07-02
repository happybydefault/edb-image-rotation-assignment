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

// Flip returns the rotated image of an ASCII encoded PBM. The degrees should be a multiple of a quarter turn (e.g. 90, 180, -270, etc.), otherwise it
// returns a non-nil error.
func Flip(image io.Reader, degrees int, ccw bool) ([]byte, error) {
	if image == nil {
		return nil, errors.New("image is nil")
	}

	quarterTurn := 90
	if degrees%quarterTurn > 0 {
		return nil, fmt.Errorf("number of degrees is not multiple of a quarter turn")
	}

	r := bufio.NewReader(image)

	sizeStr, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	sizeStr = strings.TrimSuffix(sizeStr, "\n")

	size := strings.Split(sizeStr, " ")
	if len(size) < 2 {
		return nil, fmt.Errorf("invalid size string")
	}

	width, err := strconv.Atoi(size[0])
	if err != nil {
		return nil, fmt.Errorf("invalid width: %w", err)
	}

	height, err := strconv.Atoi(size[1])
	if err != nil {
		return nil, fmt.Errorf("invalid height: %w", err)
	}

	log.Printf("width: %d, height: %d", width, height)

	return nil, nil
}
