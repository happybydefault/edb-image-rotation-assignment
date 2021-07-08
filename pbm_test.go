package pbm

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestRotate(t *testing.T) {
	image := `P1
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

	image90 := `P1
10 6
0 0 0 1 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0 0
0 0 0 1 1 1 1 1 1 1
0 0 0 0 0 0 0 0 0 0
`

	image180 := `P1
6 10
0 0 0 0 0 0
0 0 0 0 0 0
0 0 1 1 1 0
0 1 0 0 0 1
0 1 0 0 0 0
0 1 0 0 0 0
0 1 0 0 0 0
0 1 0 0 0 0
0 1 0 0 0 0
0 1 0 0 0 0
`

	image270 := `P1
10 6
0 0 0 0 0 0 0 0 0 0
1 1 1 1 1 1 1 0 0 0
0 0 0 0 0 0 0 1 0 0
0 0 0 0 0 0 0 1 0 0
0 0 0 0 0 0 0 1 0 0
0 0 0 0 0 0 1 0 0 0
`

	imageWithComments := `P1
# Some
# comments
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

	imageWithoutWhitespaces := `P1
# Some
# comments
6 10
000010000010000010000010000010000010100010011100000000000000
`

	imageInvalidSize := `P1
50 -1
000010000010000010000010000010000010100010011100000000000000
`

	type args struct {
		image   io.Reader
		degrees int
		ccw     bool
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "assignment example, 90 degrees clockwise",
			args: args{
				image:   strings.NewReader(image),
				degrees: 90,
				ccw:     false,
			},
			want:    image90,
			wantErr: false,
		},
		{
			name: "assignment example, 180 degrees clockwise",
			args: args{
				image:   strings.NewReader(image),
				degrees: 180,
				ccw:     false,
			},
			want:    image180,
			wantErr: false,
		},
		{
			name: "assignment example, 270 degrees clockwise",
			args: args{
				image:   strings.NewReader(image),
				degrees: 270,
				ccw:     false,
			},
			want:    image270,
			wantErr: false,
		},
		{
			name: "assignment example, 270 degrees counterclockwise",
			args: args{
				image:   strings.NewReader(image),
				degrees: 270,
				ccw:     true,
			},
			want:    image90,
			wantErr: false,
		},
		{
			name: "assignment example with comments, -90 degrees counterclockwise",
			args: args{
				image:   strings.NewReader(imageWithComments),
				degrees: -90,
				ccw:     true,
			},
			want:    image90,
			wantErr: false,
		},
		{
			name: "assignment example with comments and without whitespaces, -270 degrees clockwise",
			args: args{
				image:   strings.NewReader(imageWithoutWhitespaces),
				degrees: -270,
				ccw:     false,
			},
			want:    image90,
			wantErr: false,
		},
		{
			name: "assignment example, 0 degrees clockwise",
			args: args{
				image:   strings.NewReader(image),
				degrees: 0,
				ccw:     false,
			},
			want:    image,
			wantErr: false,
		},
		{
			name: "assignment example, 360 degrees counterclockwise",
			args: args{
				image:   strings.NewReader(image),
				degrees: 360,
				ccw:     true,
			},
			want:    image,
			wantErr: false,
		},
		{
			name: "assignment example, 123 degrees clockwise",
			args: args{
				image:   strings.NewReader(image),
				degrees: 123,
				ccw:     false,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "assignment example with invalid size, 90 degrees clockwise",
			args: args{
				image:   strings.NewReader(imageInvalidSize),
				degrees: 90,
				ccw:     false,
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := Rotate(buf, test.args.image, test.args.degrees, test.args.ccw)
			if (err != nil) != test.wantErr {
				t.Errorf("Rotate() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			got := buf.String()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Rotate() got = %v, want %v", got, test.want)
			}
		})
	}
}
