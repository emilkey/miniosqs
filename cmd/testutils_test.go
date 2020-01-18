package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func newTestApplication(t *testing.T) *application {
    return &application{
        errorLog: log.New(ioutil.Discard, "", 0),
        infoLog:  log.New(ioutil.Discard, "", 0),
    }
}

type testServer struct {
    *httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
    ts := httptest.NewServer(h)
    return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
    rs, err := ts.Client().Get(ts.URL + urlPath)
    if err != nil {
        t.Fatal(err)
    }

    defer rs.Body.Close()
    body, err := ioutil.ReadAll(rs.Body)
    if err != nil {
        t.Fatal(err)
    }

    return rs.StatusCode, rs.Header, body
}

func (ts *testServer) post(t *testing.T, urlPath string, data string) (int, http.Header, []byte) {
    r := strings.NewReader(data)
    rs, err := ts.Client().Post(ts.URL + urlPath, "application/json", r)
    if err != nil {
        t.Fatal(err)
    }

    defer rs.Body.Close()
    body, err := ioutil.ReadAll(rs.Body)
    if err != nil {
        t.Fatal(err)
    }

    return rs.StatusCode, rs.Header, body
}
