package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("races two servers", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL, 1*time.Second)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if err != nil {
			t.Fatalf("it shouldn't have an error %v", err)
		}
	})

	t.Run("return an error if a server doesn't response within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Millisecond)
		serverB := makeDelayedServer(12 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL, 5*time.Millisecond)

		if err == nil {
			t.Errorf("expect an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
