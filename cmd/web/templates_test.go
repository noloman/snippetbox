package main

import (
	"github.com/noloman/snippetbox/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name,
	// input to our humanDate() function (the tm field), and expected output
	// (the want field).
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2021, 02, 01, 0, 0, 0, 0, time.UTC),
			want: "01 Feb 2021 at 00:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2021, 02, 01, 0, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "31 Jan 2021 at 23:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			assert.Equal(t, hd, tt.want)
		})
	}
}
