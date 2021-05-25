# Autostarter

[![GoDoc](https://godoc.org/github.com/danieloliveira085/Autostarter?status.svg)](https://godoc.org/github.com/danieloliveira085/Autostarter) [![license](https://img.shields.io/github/license/danieloliveira085/Autostarter.svg?style=flat)](https://github.com/danieloliveira085/Autostarter/blob/main/LICENSE) 
  
Autostarter is a Go library that creates a shortcut to run automatically at startup and supports cross-compilation between Windows and Linux 

## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see [this page](https://golang.org/doc/install) first.

```sh
go get github.com/danieloliveira085/Autostarter
```

### Usage

Import the package into your project.

```go
import "github.com/danieloliveira085/Autostarter"
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
