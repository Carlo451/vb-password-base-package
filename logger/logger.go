package logger

import (
	"log/slog"
	"os"
)

var Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
var ApiLogger = Logger.With("[passwordstore api]")
var IoLogger = Logger.With("[passwordstore io]")
