package ui

import (
	"os"

	"github.com/kathleenfrench/ks/internal/theme"
)

func ExitOnErr(errMsg string) {
	theme.Err(errMsg)
	os.Exit(1)
}

func Exit() {
	theme.Info("bye!")
	os.Exit(1)
}
