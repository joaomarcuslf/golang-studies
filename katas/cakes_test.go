package main

import "testing"

func Test_Cakes(t *testing.T) {
	type args struct {
		recipe    map[string]int
		available map[string]int
	}

	type test struct {
		name string
		args args
		want int
	}

	tests := []test{
		{
			name: "#1",
			args: args{
				recipe:    map[string]int{"flour": 500, "sugar": 200, "eggs": 1},
				available: map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200},
			},
			want: 2,
		},
		{
			name: "#2",
			args: args{
				recipe:    map[string]int{"flour": 500, "sugar": 200, "eggs": 1},
				available: map[string]int{"flour": 200, "sugar": 1200, "eggs": 5, "milk": 200},
			},
			want: 0,
		},
		{
			name: "#3",
			args: args{
				recipe:    map[string]int{"flour": 500, "sugar": 200, "eggs": 1},
				available: map[string]int{"flour": 2500, "sugar": 1200, "eggs": 1, "milk": 200},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cakes(tt.args.recipe, tt.args.available); got != tt.want {
				t.Errorf("Cakes() %s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
