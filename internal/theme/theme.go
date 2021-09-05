package theme

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const yellow = "#ffee00"
const green = "#42f590"

var (
	subtle  = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	special = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	logoStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#42f590")).
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			SetString("ks")

	url = lipgloss.NewStyle().Foreground(special).Render

	descStyle = lipgloss.NewStyle()

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)

	helpStyle = lipgloss.NewStyle().Italic(true).PaddingTop(1)

	desc = lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("a simple utility for encoding/decoding k8s secrets"),
		infoStyle.Render("by kathleen french"+divider+url("https://github.com/kathleenfrench/ks")),
		helpStyle.Render("run ks -h for more info on available commands"),
	)

	logoWrapperStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)

	stdoutStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(yellow)).
			MarginLeft(1).
			PaddingTop(1).
			PaddingBottom(1)

	resultStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(green)).
			MarginLeft(1).
			PaddingTop(1).
			PaddingBottom(1)
)

func Info(text string) {
	fmt.Println(stdoutStyle.Render(text))
}

func Result(text string) {
	fmt.Println(resultStyle.Render(text))
}

func PrintLogo() {
	termWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}
	row := lipgloss.JoinHorizontal(lipgloss.Top, logoStyle.String(), desc)
	doc.WriteString(row + "\n\n")

	if termWidth > 0 {
		logoWrapperStyle = logoWrapperStyle.MaxWidth(termWidth)
	}

	fmt.Println(logoWrapperStyle.Render(doc.String()))
}
