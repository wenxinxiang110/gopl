package chapter2

import "testing"

func TestPopCount(t *testing.T) {

	tests := []struct {
		name string
		args uint64
		want int
	}{
		{
			"case-1",
			1,
			1,
		},
		{
			"case-2",
			2,
			1,
		},
		{
			"case-3",
			3,
			2,
		},
		{
			"case-7",
			7,
			3,
		},
		{
			"case-31",
			31,
			5,
		},
		{
			"case-256",
			256,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCount(tt.args); got != tt.want {
				t.Errorf("PopCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
