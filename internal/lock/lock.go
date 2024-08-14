package lock

import (
	"errors"

	"github.com/gofrs/flock"
	"github.com/patrickap/runr/m/v2/internal/log"
)

var lock = flock.New("/tmp/runr.lock")

func Lock() error {
	locked, err := lock.TryLock()

	if !locked {
		return errors.New("failed to acquire lock: unable to get exclusive lock")
	}

	if err != nil {
		return err
	}

	return nil
}

func Unlock() error {
	return lock.Unlock()
}

func RunWithLock(f func() error) error {
	err := lock.Lock()
	if err != nil {
		log.Instance().Error().Msgf("Failed to acquire lock: %v", err)
		return err
	}

	defer lock.Unlock()

	return f()
}
