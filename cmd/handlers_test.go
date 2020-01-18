package main

import (
    "net/http"
    "testing"
)

func Test405(t *testing.T) {
    app := newTestApplication(t)
    ts := newTestServer(t, app.routes())
    defer ts.Close()

    code, _, _ := ts.get(t, "/")

    if code != 405 {
        t.Errorf("want %d; got %d", http.StatusOK, code)
    }
}

func Test404(t *testing.T) {
    app := newTestApplication(t)
    ts := newTestServer(t, app.routes())
    defer ts.Close()

    code, _, _ := ts.get(t, "/doesnotexist")

    if code != 404 {
        t.Errorf("want %d; got %d", http.StatusOK, code)
    }
}

func TestEmpty(t *testing.T) {
    app := newTestApplication(t)
    ts := newTestServer(t, app.routes())
    defer ts.Close()

    code, _, body := ts.post(t, "/", "")

    if code != 200 {
        t.Errorf("want %d; got %d", http.StatusOK, code)
    }

    desiredBodyText := "Empty body"
    if string(body) != desiredBodyText {
        t.Errorf("want body to equal '%s'; got '%s'", desiredBodyText, body)
    }
}
