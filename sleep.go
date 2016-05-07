package main

import (
	"time"
)

func Sleep(duration string) error {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}
	time.Sleep(dur)
	return nil
}
