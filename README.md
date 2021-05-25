# Autostarter

[![GoDoc](https://godoc.org/github.com/danieloliveira085/autostarter?status.svg)](https://godoc.org/github.com/danieloliveira085/autostarter) [![license](https://img.shields.io/github/license/danieloliveira085/autostarter.svg?style=flat)](https://github.com/danieloliveira085/autostarter/blob/main/LICENSE) 
  
Autostarter is a Go library that creates a shortcut to run automatically at startup and supports cross-compilation between Windows and Linux 

## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see [this page](https://golang.org/doc/install) first.

```sh
go get github.com/danieloliveira085/autostarter
```

### Usage

Import the package into your project.

```go
import "github.com/danieloliveira085/autostarter"
```

Construct a new autostart that can be used to access the main functions of the autostart created

```go
a := autostarter.NewAutostart(
	autostarter.Shortcut{
		Name:    "Shortcut name",
		Exec:    "Exec",
		Args:    []string{}, //Arguments, can be empty
		StartIn: "Path where exec starts",
	},
	autostarter.DefaultIcon, //Icon, for a custom icon, use SetIcon()
)
```

See Documentation on GoDoc for more detailed information.
