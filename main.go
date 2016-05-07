package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %v SOURCE TARGET DURATION\n", os.Args[0])
		os.Exit(2)
	}
	source, target, duration := os.Args[1], os.Args[2], os.Args[3]
	err := Exists(target)
	if err != nil {
		// not exists or is not ok
		err = CopyFile(source, target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CopyFile failed %q\n", err)
			os.Exit(1)
		}
	}
	err = Sleep(duration)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Sleeping failed %q\n", err)
		os.Exit(1)
	}
}
