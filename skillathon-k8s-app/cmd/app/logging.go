package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type stdErrorLogWrapper struct {
	logger *zerolog.Logger
}

// Write implements io.Writer interface to redirect standard logger entries to zerolog.
// Also, it cuts caller from a log entry and passes it to zerolog's caller
func (s stdErrorLogWrapper) Write(p []byte) (n int, err error) {
	caller, errorMsg, _ := strings.Cut(string(p), " ")
	caller = strings.TrimRight(caller, ":")

	s.logger.Error().
		Str("caller", caller).
		Str("error", errorMsg).
		Msg("")

	return len(p), nil
}

// configureLogging configures zerolog and sets the respective fields in the application struct
func (app *application) configureLogging() {
	zlog.Logger = zlog.Output(os.Stdout)
	app.logger = &zlog.Logger
	logWrapper := stdErrorLogWrapper{logger: app.logger}
	app.errorLog = log.New(logWrapper, "", log.Lshortfile)

	zerolog.CallerMarshalFunc = app.lshortfile
	zerolog.DurationFieldUnit = time.Second

	if app.LogFormat == "pretty" {
		zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: app.LogNoColor})
	}
}

// lshortfile implements Lshortfile equivalent for zerolog's CallerMarshalFunc
func (app *application) lshortfile(pc uintptr, file string, line int) string {
	// Copied from the standard library: https://cs.opensource.google/go/go/+/refs/tags/go1.17.8:src/log/log.go;drc=926994fd7cf65b2703552686965fb05569699897;l=134
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return file + ":" + strconv.Itoa(line)
}
