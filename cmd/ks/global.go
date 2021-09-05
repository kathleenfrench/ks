package ks

import (
	"github.com/kathleenfrench/ks/pkg/clipboard"
	"github.com/kathleenfrench/ks/pkg/parse"
)

// interfaces
var (
	p    = parse.NewParser()
	clip = clipboard.NewClipboard()
)

// flags
var (
	verbose bool
	silent  bool
)
