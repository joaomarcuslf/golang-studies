package main

import "testing"

func Test_DecodeMorse(t *testing.T) {
	type args struct {
		morseCode string
	}

	type test struct {
		name string
		args args
		want string
	}

	tests := []test{
		{
			name: "HEY JUDE",
			args: args{
				morseCode: ".... . -.--   .--- ..- -.. .",
			},
			want: "HEY JUDE",
		},
		{
			name: "HEY JUDE",
			args: args{
				morseCode: "   .... . -.--   .--- ..- -.. .",
			},
			want: "HEY JUDE",
		},
		{
			name: "SOS",
			args: args{
				morseCode: "...−−−...",
			},
			want: "SOS",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeMorse(tt.args.morseCode); got != tt.want {
				t.Errorf("DecodeMorse() %s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
