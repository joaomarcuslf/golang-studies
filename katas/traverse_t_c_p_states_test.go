package main

import "testing"

func Test_TraverseTCPStates(t *testing.T) {
	type args struct {
		events []string
	}

	type test struct {
		name string
		args args
		want string
	}

	tests := []test{
		{
			name: "CLOSE_WAIT",
			args: args{
				events: []string{"APP_ACTIVE_OPEN", "RCV_SYN_ACK", "RCV_FIN"},
			},
			want: "CLOSE_WAIT",
		},
		{
			name: "ESTABLISHED",
			args: args{
				events: []string{"APP_PASSIVE_OPEN", "RCV_SYN", "RCV_ACK"},
			},
			want: "ESTABLISHED",
		},
		{
			name: "LAST_ACK",
			args: args{
				events: []string{"APP_ACTIVE_OPEN", "RCV_SYN_ACK", "RCV_FIN", "APP_CLOSE"},
			},
			want: "LAST_ACK",
		},
		{
			name: "SYN_SENT",
			args: args{
				events: []string{"APP_ACTIVE_OPEN"},
			},
			want: "SYN_SENT",
		},
		{
			name: "ERROR",
			args: args{
				events: []string{"APP_PASSIVE_OPEN", "RCV_SYN", "RCV_ACK", "APP_CLOSE", "APP_SEND"},
			},
			want: "ERROR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TraverseTCPStates(tt.args.events); got != tt.want {
				t.Errorf("TraverseTCPStates() %s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
