package pbm

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestFlip(t *testing.T) {
	image := []byte(`6 10
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
`)
	imageCW90 := []byte(`10 6
0 0 0 1 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0 0
0 0 0 1 1 1 1 1 1 1
0 0 0 0 0 0 0 0 0 0
`)

	type args struct {
		image   io.Reader
		degrees int
		ccw     bool
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "assignment example, 90 degrees clockwise",
			args: args{
				image:   bytes.NewReader(image),
				degrees: 90,
				ccw:     false,
			},
			want:    imageCW90,
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Flip(test.args.image, test.args.degrees, test.args.ccw)
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
