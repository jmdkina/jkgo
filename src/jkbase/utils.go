package jkbase

import (
    "fmt"
    "os/exec"
    "runtime"
)

var commands = map[string][]string{
    "windows": {"cmd", "/c", "start"},
    "darwin":  {"open"},
    "linux":   {"xdg-open"},
}

// Open calls the OS default program for uri
func JKOpenBrowser(uri string) error {
    run, ok := commands[runtime.GOOS]
    if !ok {
        return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
    }

	run = append(run, uri)
    cmd := exec.Command(run[0], run[1:]...)
    return cmd.Start()
}
