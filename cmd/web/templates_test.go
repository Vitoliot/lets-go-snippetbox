package main

import (
	"lets-go-snippetbox/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	cases := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 7, 1, 10, 15, 0, 0, time.UTC),
			want: "01 Jul 2024 at 10:15",
		},
		{
			name: "Blank",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 7, 1, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "01 Jul 2024 at 09:15",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, humanDate(tt.tm), tt.want)
		})
	}
}
