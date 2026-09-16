package gen

import "testing"

func TestIsArrayValid(t *testing.T) {
	tests := []struct {
		name string
		data [9]int
		want bool
	}{
		{name: "partial row", data: [9]int{1, 2, 0, 4, 5, 0, 7, 8, 9}, want: true},
		{name: "duplicate value", data: [9]int{1, 2, 2, 0, 5, 6, 7, 8, 9}, want: false},
		{name: "value above nine", data: [9]int{1, 2, 3, 4, 5, 6, 7, 8, 10}, want: false},
		{name: "negative value", data: [9]int{1, 2, 3, 4, 5, 6, 7, 8, -1}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsArrayValid(tt.data); got != tt.want {
				t.Fatalf("IsArrayValid(%v) = %t, want %t", tt.data, got, tt.want)
			}
		})
	}
}
