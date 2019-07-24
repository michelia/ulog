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
	New(c).Info().Dur("dur", time.Second).Msg("test")
}
func TestNewConsole(t *testing.T) {
	log := NewConsole()
	log.Print("test print")
	log.Error().Str("f1", "foo").Msg("test")
	NewConsole().Print("test print")
}
