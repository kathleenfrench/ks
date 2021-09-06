package ks

import (
	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:     "gui",
	Aliases: []string{"i"},
	Example: "ks [gui, i]",
	Short:   "run ks in interactive mode",
	Run: func(cmd *cobra.Command, args []string) {
		theme.Info("TK")
	},
}
