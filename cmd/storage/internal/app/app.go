package app

import (
	"errors"

	"github.com/powerman/structlog"
)

// Errors.
var (
	ErrNotFound = errors.New("not found")
)

var log = structlog.New()

type (
	// App is an application interface.
	App interface {
		Save(stationID string, key string, value []byte) error
		Load(stationID string, key string) ([]byte, error)
		Info() string
		KasseInfo() string
	}
	// Repo is a DAL interface.
	Repo interface {
		Save(stationID string, key string, value []byte) error
		Load(stationID string, key string) ([]byte, error)
		Info() string
	}
	// KasseSvc is an interface for kasse service.
	KasseSvc interface {
		Info() (string, error)
	}
)

type app struct {
	repo     Repo
	kasseSvc KasseSvc
}

// New creates and returns new App.
func New(repo Repo, kasseSvc KasseSvc) App {
	return &app{
		repo:     repo,
		kasseSvc: kasseSvc,
	}
}
