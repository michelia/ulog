package ulog

import (
	"testing"
	"time"
)

func TestNewLog(t *testing.T) {
	c := Config{
		Filename: "log",
		IsJSON:   true,
		Fields: map[string]string{
			"f1": "foo",
			"f2": "bar",
		},
	}
	NewLog(c).Info().Dur("dur", time.Second).Msg("test")
}
func TestNewLogConsole(t *testing.T) {
	log := NewLogConsole()
	log.Print("test print")
	log.Error().Str("f1", "foo").Msg("test")
	NewLogConsole(true).Print("noColor test print")
	NewLogConsole(false).Print("test print")
}
