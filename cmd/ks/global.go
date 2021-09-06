package ks

import (
	"github.com/kathleenfrench/ks/internal/secret"
	"github.com/kathleenfrench/ks/pkg/file"
	"github.com/kathleenfrench/ks/pkg/parse"
)

// interfaces
var (
	p  = parse.NewParser()
	fm = file.NewManager()
	sm = secret.NewManager()
)

// global flags
var (
	verbose    bool
	silent     bool
	targetFile string
)
