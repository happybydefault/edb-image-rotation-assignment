package pbm

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestFlip(t *testing.T) {
	image := `P1
# This is an example bitmap of the letter "J"
6 10
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
1 0 0 0 1 0
0 1 1 1 0 0
0 0 0 0 0 0
0 0 0 0 0 0
`
	imageCW90 := `P1
# This is an example bitmap of the letter "J"
10 6
0 0 0 1 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0 0
0 0 0 1 1 1 1 1 1 1
0 0 0 0 0 0 0 0 0 0
`

	type args struct {
		r       io.Reader
		degrees int
	}

	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		{
			name: "assignment example, 90 degrees clockwise",
			args: args{
				r:       bytes.NewReader([]byte(image)),
				degrees: 90,
			},
			want:    bytes.NewReader([]byte(imageCW90)),
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Flip(test.args.r, test.args.degrees)
			if (err != nil) != test.wantErr {
				t.Errorf("Flip() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Flip() got = %v, want %v", got, test.want)
			}
		})
	}
}
