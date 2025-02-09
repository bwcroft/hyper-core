package utils

import (
	"fmt"
	"log"
	"strings"
)

func LogInfo(v any) {
	log.Println("Info", v)
}

func LogWarn(v any) {
	log.Println("Warning:", v)
}

func LogError(err error) {
	log.Println("Error:", err)
}

func InfoBox(lines []string) {
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	boxWidth := maxWidth + 2
	fmt.Println("\033[34m", "┌" + strings.Repeat("─", boxWidth) + "┐")
	for _, line := range lines {
		fmt.Printf(" │ %-*s │\n", maxWidth, line)
	}
	fmt.Println(" └" + strings.Repeat("─", boxWidth) + "┘", "\033[0m")
}
