package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/noloman/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	ping(rr, r)

	rs := rr.Result()

	assert.Equal(t, http.StatusOK, rs.StatusCode)
}
