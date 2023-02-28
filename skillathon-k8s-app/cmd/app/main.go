package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"go.uber.org/automaxprocs/maxprocs"
)

type application struct {
	SetGomaxProcs           bool          `env:"APP_SET_GOMAXPROCS" envDefault:"true"`
	Port                    int           `env:"APP_PORT" envDefault:"8080"`
	ReadTimeout             time.Duration `env:"APP_READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout            time.Duration `env:"APP_WRITE_TIMEOUT" envDefault:"10s"`
	GracefulShutdownTimeout time.Duration `env:"APP_SHUTDOWN_TIMEOUT" envDefault:"20s"`
	ShutdownDelay           time.Duration `env:"APP_SHUTDOWN_DELAY" envDefault:"3s"`
	LogFormat               string        `env:"APP_LOG_FORMAT" envDefault:"pretty"`
	LogNoColor              bool          `env:"APP_LOG_NO_COLOR" envDefault:"false"`
	FailLiveness            bool
	FailReadiness           bool
	LivenessMutex           sync.Mutex
	ReadinessMutex          sync.Mutex
	logger                  *zerolog.Logger
	errorLog                *log.Logger
}

func main() {
	app := &application{}

	err := env.Parse(app)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	app.configureLogging()

	if app.SetGomaxProcs {
		undo, err := maxprocs.Set()
		defer undo()
		if err != nil {
			app.logger.Error().Caller().
				Msgf("failed to set GOMAXPROCS: %v", err)
		}
	}
	app.logger.Info().Caller().
		Msgf("Runtime settings: GOMAXPROCS = %d", runtime.GOMAXPROCS(0))

	err = app.serve()
	if err != nil {
		app.logger.Fatal().Caller().
			Err(err).Msg("")
	}
}
