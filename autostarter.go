package autostarter

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Default icon, defined by OS
const DefaultIcon icon = ""

type Autostart struct {
	sc   Shortcut
	ic   icon
	path string
}

type Shortcut struct {
	// Executable name
	Name string
	// Executable path, ex: windows: C:/Windows/explorer.exe linux: /usr/bin/gcc
	Exec string
	// The arguments for executable, can be empty
	Args []string
	// The directory where the executable starts
	StartIn string
}

type icon string

// Windows: Accepted formats: ico
// Linux: Accepted formats: ico, png, jpg
func SetIcon(path string) (icon, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errors.New("path invalid: " + path)
	}
	if err := checkIcon(path); err != nil {
		return "", err
	}
	return icon(filepath.Clean(path)), nil
}

// Create a new autostart
// The startup with system is disabled by default, use Enable() to enable it
// or use IsEnabled () to check if this autostart has been enabled previously
func NewAutostart(sc Shortcut, ic icon) *Autostart {
	return &Autostart{
		sc: sc,
		ic: ic,
	}
}

// Return TRUE if the autostart is enabled
func (a *Autostart) IsEnabled() bool {
	if _, err := os.Stat(a.path); !os.IsNotExist(err) {
		return true
	}
	return false
}

// Enable the autostart
func (a *Autostart) Enable() error {
	var err error
	a.path, err = createShortcut(a.sc, a.ic)
	if err != nil {
		return err
	}
	return nil
}

// Disable the autostart
func (a *Autostart) Disable() error {
	return os.Remove(a.path)
}

// If autostart is enabled, it deactivates and vice versa
func (a *Autostart) Trigger() bool {
	if _, err := os.Stat(a.path); !os.IsNotExist(err) {
		a.Disable()
		return false
	}
	a.Enable()
	return true
}

func (s *Shortcut) getArgsString() string {
	var args []string
	for i := range s.Args {
		args = append(args, strconv.Quote(s.Args[i]))
	}
	return strings.Join(args, " ")
}
