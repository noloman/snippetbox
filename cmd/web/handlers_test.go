package main

import (
	"net/http"
	"testing"

	"github.com/noloman/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	t.Parallel()
	app := newApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, string(body), "OK")
}
