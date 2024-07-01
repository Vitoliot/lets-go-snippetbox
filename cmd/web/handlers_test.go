package main

import (
	"bytes"
	"io"
	"lets-go-snippetbox/internal/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	ping(rr, r)

	res := rr.Result()

	assert.Equal(t, res.StatusCode, http.StatusOK)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(bytes.TrimSpace(body)), "OK")
}
