package utils

import "time"

// TODO: There's might be a better way to do this
func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}

		return nil
	}

	return
}
