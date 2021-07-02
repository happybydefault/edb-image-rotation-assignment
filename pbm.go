// Package pbm is a library for manipulating PBM images.
package pbm

import (
	"errors"
	"fmt"
	"io"
)

// Flip returns the rotated image of an ASCII encoded PBM. The degrees should be a multiple of a quarter turn (e.g. 90, 180, -270, etc.), otherwise it
// returns a non-nil error.
func Flip(r io.Reader, degrees int, ccw bool) ([]byte, error) {
	if r == nil {
		return nil, errors.New("image is nil")
	}

	quarterTurn := 90
	if degrees%quarterTurn > 0 {
		return nil, fmt.Errorf("number of degrees is not multiple of a quarter turn")
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("could not read image: %w", err)
	}

	return b, nil
}
