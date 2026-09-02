package logic

import (
	"testing"
	"time"
)

func TestChrono(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		want     string
	}{
		{name: "under an hour", duration: 5*time.Minute + 9*time.Second, want: "05:09"},
		{name: "over an hour", duration: time.Hour + 5*time.Minute + 9*time.Second, want: "01:05:09"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chrono(tt.duration); got != tt.want {
				t.Fatalf("Chrono(%v) = %q, want %q", tt.duration, got, tt.want)
			}
		})
	}
}
