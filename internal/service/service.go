package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(text string) string {
	convText := newConvertedText(text)
	return convText.res
}

type convertedText struct {
	sourceText string
	res        string
}

func newConvertedText(text string) *convertedText {
	convText := &convertedText{
		sourceText: text,
	}

	convText.convertText()

	return convText
}

func (convText *convertedText) convertText() {
	f := func(r rune) bool {
		return r == '.' || r == '-'
	}
	if strings.ContainsFunc(convText.sourceText, f) {
		convText.res = morse.ToText(convText.sourceText)
	} else {
		convText.res = morse.ToMorse(convText.sourceText)
	}
}
