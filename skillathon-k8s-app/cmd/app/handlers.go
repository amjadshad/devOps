package main

import (
	"fmt"
	"net/http"
	"time"
)

// home is the default hello world-like handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Skillathon demo app")
}

// slow is a very slooow endpoint
func (app *application) slow(w http.ResponseWriter, r *http.Request) {
	time.Sleep(7 * time.Second)
	fmt.Fprintln(w, "Hello from Zootopia sloths :)")
}

// switchReadiness essentially makes the readiness probe endpoint fail
func (app *application) switchReadiness(w http.ResponseWriter, r *http.Request) {
	app.ReadinessMutex.Lock()
	defer app.ReadinessMutex.Unlock()
	app.FailReadiness = !app.FailReadiness
	fmt.Fprintf(w, "FailReadiness is set to %t", app.FailReadiness)
}

// switchLiveness essentially makes the liveness probe endpoint fail
func (app *application) switchLiveness(w http.ResponseWriter, r *http.Request) {
	app.LivenessMutex.Lock()
	defer app.LivenessMutex.Unlock()
	app.FailLiveness = !app.FailLiveness
	fmt.Fprintf(w, "FailLiveness is set to %t", app.FailLiveness)
}
