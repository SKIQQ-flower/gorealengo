package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

func Init() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "06-01-02 @ 15:04:05",
		FormatLevel: func(i interface{}) string {
			switch i {
			case "debug":
				return "🐛"
			case "info":
				return "ℹ️"
			case "warn":
				return "⚠️"
			case "error":
				return "❌"
			case "fatal":
				return "💀"
			default:
				return "🤷"
			}
		},
		FormatMessage: func(i any) string {
			msg, ok := i.(string)
			if !ok {
				msg = ""
			}
			return ".. " + " → " + msg + "\n"
		},
	}
	Logger = zerolog.New(output).With().Timestamp().Logger()

	log.Logger = Logger
}
