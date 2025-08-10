package db

import (
	"sync/atomic"

	"entgo.io/ent/dialect"
)

// keepOpenDriver wraps a dialect.Driver and ignores Close() calls until
// explicitly enabled via allowClose. This prevents accidental closure of the
// global DB pool during runtime.
type keepOpenDriver struct {
	dialect.Driver
	allowClose atomic.Bool
}

func (d *keepOpenDriver) Close() error {
	if d.allowClose.Load() {
		return d.Driver.Close()
	}
	// Ignore premature close attempts during runtime.
	return nil
}

var currentKeepOpen *keepOpenDriver

func wrapKeepOpen(d dialect.Driver) dialect.Driver {
	k := &keepOpenDriver{Driver: d}
	currentKeepOpen = k
	return k
}

// EnableDBClose allows the wrapped driver to actually close.
func EnableDBClose() {
	if currentKeepOpen != nil {
		currentKeepOpen.allowClose.Store(true)
	}
}
