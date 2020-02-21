package ch3

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	type args struct {
		n int
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"12345-2",
			args{
				n: 2,
				s: []string{"1", "2", "3", "4", "5"},
			},
			[]string{"4", "5", "1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rotate(tt.args.n, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
