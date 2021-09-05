package ks

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/kathleenfrench/ks/pkg/clipboard"
	"github.com/kathleenfrench/ks/pkg/parse"
	"github.com/spf13/cobra"
)

var (
	p    = parse.NewParser()
	clip = clipboard.NewClipboard()
)

var rootCmd = &cobra.Command{
	Use:   "ks",
	Short: "a simple utility for base64 encoding secrets for k8s and copying them to the clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		theme.PrintLogo()
	},
}

var encodeCmd = &cobra.Command{
	Use:     "encode",
	Aliases: []string{"e"},
	Short:   "base64 encode a secret",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("must provide a secret to encode")
		}

		secret := args[0]
		encoded, err := p.Encode(strings.TrimSpace(secret))
		if err != nil {
			log.Fatal(err)
		}

		err = clip.Write(encoded)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(encoded)
	},
}

var decodeCmd = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"d"},
	Short:   "decode a base64 encoded secret",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("must provide a secret to decode")
		}

		secret := args[0]
		decoded, err := p.Decode(strings.TrimSpace(secret))
		if err != nil {
			log.Fatal(err)
		}

		err = clip.Write(decoded)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(decoded)
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	rootCmd.AddCommand(decodeCmd)
	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd: true,
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
