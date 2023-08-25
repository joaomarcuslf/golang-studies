package main

import "testing"

func Test_LCS(t *testing.T) {
	type args struct {
		s1 string
		s2 string
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
				s1: "abcdef",
				s2: "abc",
			},
			want: "abc",
		},
		{
			name: "#2",
			args: args{
				s1: "abcdef",
				s2: "acf",
			},
			want: "acf",
		},
		{
			name: "#3",
			args: args{
				s1: "132535365",
				s2: "123456789",
			},
			want: "12356",
		},
		{
			name: "#4",
			args: args{
				s1: "abcdefghijklmnopq",
				s2: "apcdefghijklmnobq",
			},
			want: "acdefghijklmnoq",
		},
		{
			name: "#5",
			args: args{
				s1: "eidhacgfb",
				s2: "gifbcdaeh",
			},
			want: "idh",
		},

		{
			name: "#6",
			args: args{
				s1: "dgafceihb",
				s2: "adgebichf",
			},
			want: "dgeih",
		},

		{
			name: "#7",
			args: args{
				s1: "fbedgicha",
				s2: "aebchifdg",
			},
			want: "fdg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCS(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("LCS() %s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
