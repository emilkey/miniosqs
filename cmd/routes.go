package main

import (
    "net/http"
    "github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
    standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    mux := http.NewServeMux()
    mux.HandleFunc("/", app.notifySqs)
    mux.HandleFunc("/bucket/setup", app.setupMinioBucket)

    return standardMiddleware.Then(mux)
}
