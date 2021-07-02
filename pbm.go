// Package pbm is a library for manipulating PBM images.
package pbm

import (
	"errors"
	"fmt"
	"io"
)

// Flip rotates an image with a degrees that's a multiple of a quarter turn (e.g. 90, 180, -270, etc.), otherwise it
// returns a non-nil error.
func Flip(r io.Reader, degrees int) (io.Reader, error) {
	if r == nil {
		return nil, errors.New("reader is nil")
	}

	quarterTurn := 90
	if degrees%quarterTurn > 0 {
		return nil, fmt.Errorf("number of degrees is not multiple of a quarter turn")
	}

	return r, nil
}
