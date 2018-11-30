package floydwarshall

import (
	"reflect"
	"testing"
)

func Test_transitiveClosure(t *testing.T) {

	tests := []struct {
		name string
		W    [][]bool
		want [][]bool
	}{
		{
			name: "case1",
			W: [][]bool{
				[]bool{true, false, false, false},
				[]bool{false, true, true, true},
				[]bool{false, true, true, false},
				[]bool{true, false, true, true},
			},
			want: [][]bool{
				[]bool{true, false, false, false},
				[]bool{true, true, true, true},
				[]bool{true, true, true, true},
				[]bool{true, true, true, true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transitiveClosure(tt.W); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transitiveClosure() = %v, want %v", got, tt.want)
			}
		})
	}
}
