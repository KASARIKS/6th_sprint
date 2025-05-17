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
	sourceText  string
	checkString string
	res         string
}

func newConvertedText(text string) *convertedText {
	convText := &convertedText{
		sourceText: text,
	}

	convText.convertText()

	return convText
}

func (convText *convertedText) convertText() {
	convText.makeCheckString()
	convText.makeRes()
}

func (convText *convertedText) makeCheckString() {
	convText.checkString = strings.Replace(convText.sourceText, ".", "", -1)
	convText.checkString = strings.Replace(convText.checkString, "-", "", -1)
	convText.checkString = strings.Replace(convText.checkString, " ", "", -1)
}

func (convText *convertedText) makeRes() {
	if convText.checkString == "" {
		convText.res = morse.ToText(convText.sourceText)
	} else {
		convText.res = morse.ToMorse(convText.sourceText)
	}
}
