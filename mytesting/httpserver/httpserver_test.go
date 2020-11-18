package httpserver

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	ep := errPinger{}

	srv := Server{dbCli: &ep}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("", "", nil)

	srv.Ping(w, r)
	if w.Code != 200 {
		t.Fatalf("expected code %d, got %d", 200, w.Code)
	}

	ep.err = errors.New("oops")

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("", "", nil)

	srv.Ping(w, r)

	if w.Code != 500 {
		t.Fatalf("expected code %d, got %d", 500, w.Code)
	} else if w.Body.String() != "oops\n" {
		t.Fatalf("expected body %q, got %q", "oops", w.Body.String())
	}
}

type errPinger struct {
	err error
}

func (e *errPinger) Ping() error {
	return e.err
}

func TestRouter(t *testing.T) {
	ms := &Server{dbCli: (*nilPinger)(nil)}
	ts := httptest.NewServer(ms.Handler())
	defer ts.Close()

	pingurl := ts.URL + "/ping"

	res, err := http.Get(pingurl)
	if err != nil {
		t.Fatal(err)
	} else if res.StatusCode != 200 {
		t.Fatalf("expected code %d, got %d", 200, res.StatusCode)
	}
}

type nilPinger struct{}

func (*nilPinger) Ping() error { return nil }
