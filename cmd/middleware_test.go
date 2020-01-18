package main

import (
    "testing"
)

func TestSecureHeaders(t *testing.T) {
    app := newTestApplication(t)
    ts := newTestServer(t, app.routes())
    defer ts.Close()

    code, header, _ := ts.get(t, "/")

    frameOptions := header.Get("X-Frame-Options")
    if frameOptions != "deny" {
        t.Errorf("want %q; got %q", "deny", frameOptions)
    }

    xssProtection := header.Get("X-XSS-Protection")
    if xssProtection != "1; mode=block" {
        t.Errorf("want %q; got %q", "1; mode=block", xssProtection)
    }

    // Check that the middleware has correctly called the next handler in line
    // and the response status code and body are as expected.
    if code != 405 {
       t.Errorf("want %d; got %d", 405, code)
    }

}
