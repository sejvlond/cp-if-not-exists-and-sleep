package main

import (
	"fmt"
	"os"
)

func Exists(src string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("File is not regular")
	}
	return nil
}
