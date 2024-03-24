package models

import (
	"testing"

	"github.com/noloman/snippetbox/internal/assert"
)

func TestUserExists(t *testing.T) {
	// Skip the test if the "-short" flag is provided when running the test.
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}
	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "User exists",
			userID: 1,
			want:   true,
		},
		{
			name:   "User does not exist",
			userID: 2,
			want:   false,
		},
		{
			name:   "Zero ID",
			userID: 0,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDatabase(t)
			m := UserModel{DB: db}
			got, err := m.Exists(tt.userID)
			assert.Equal(t, got, tt.want)
			assert.NilError(t, err)
		})
	}
}
