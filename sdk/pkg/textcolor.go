package pkg

import (
	"fmt"
)

// TextBlack ...
const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

// Black ...
func Black(msg string) string {
	return SetColor(msg, 0, 0, TextBlack)
}

// Red ...
func Red(msg string) string {
	return SetColor(msg, 0, 0, TextRed)
}

// Green ...
func Green(msg string) string {
	return SetColor(msg, 0, 0, TextGreen)
}

// Yellow ...
func Yellow(msg string) string {
	return SetColor(msg, 0, 0, TextYellow)
}

// Blue ...
func Blue(msg string) string {
	return SetColor(msg, 0, 0, TextBlue)
}

// Magenta ...
func Magenta(msg string) string {
	return SetColor(msg, 0, 0, TextMagenta)
}

// Cyan ...
func Cyan(msg string) string {
	return SetColor(msg, 0, 0, TextCyan)
}

// White ...
func White(msg string) string {
	return SetColor(msg, 0, 0, TextWhite)
}

// SetColor ...
func SetColor(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}
