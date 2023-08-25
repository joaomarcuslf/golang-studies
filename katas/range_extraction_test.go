package main

import "testing"

func Test_RangeExtraction(t *testing.T) {
	type args struct {
		input []int
	}

	type test struct {
		name string
		args args
		want string
	}

	tests := []test{
		{
			name: "#1",
			args: args{
				input: []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20},
			},
			want: "-6,-3-1,3-5,7-11,14,15,17-20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RangeExtraction(tt.args.input); got != tt.want {
				t.Errorf("RangeExtraction() %s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
