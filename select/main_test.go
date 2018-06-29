package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run(" run normally", func(t *testing.T) {
		slowServer := makeDelayServer(5 * time.Millisecond)
		fastServer := makeDelayServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL, 10*time.Millisecond)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

		if err != nil {
			t.Fatalf("%v error is not expected", err)
		}

	})
	t.Run("run with timeout", func(t *testing.T) {
		serverA := makeDelayServer(12 * time.Millisecond)
		serverB := makeDelayServer(11 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("Error is expected but doesn't happen")
		}
	})
}

func makeDelayServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, t *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}
