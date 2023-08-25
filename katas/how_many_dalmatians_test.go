package main

import (
	"testing"
)

func Test_HowManyDalmatians(t *testing.T) {
	type args struct {
		number int
	}

	type test struct {
		name string
		args args
		want string
	}

	tests := []test{
		{
			name: "26 dogs",
			args: args{
				number: 26,
			},
			want: "More than a handful!",
		},
		{
			name: "8 dogs",
			args: args{
				number: 8,
			},
			want: "Hardly any",
		},
		{
			name: "14 dogs",
			args: args{
				number: 14,
			},
			want: "More than a handful!",
		},
		{
			name: "80 dogs",
			args: args{
				number: 80,
			},
			want: "Woah that's a lot of dogs!",
		},
		{
			name: "100 dogs",
			args: args{
				number: 100,
			},
			want: "Woah that's a lot of dogs!",
		},
		{
			name: "50 dogs",
			args: args{
				number: 50,
			},
			want: "More than a handful!",
		},
		{
			name: "10 dogs",
			args: args{
				number: 10,
			},
			want: "Hardly any",
		},
		{
			name: "101 dogs",
			args: args{
				number: 101,
			},
			want: "101 DALMATIONS!!!",
		},
	}

	for _, tt := range tests {
		if got := HowManyDalmatians(tt.args.number); got != tt.want {
			t.Errorf("HowManyDalmatians() %s = %v, want %v", tt.name, got, tt.want)
		}
	}
}
