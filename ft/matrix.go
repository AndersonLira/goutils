package ft

import (
	"fmt"
	"math"
	"strings"
)

//A Matrix represents struct to print some value formatted
type Matrix struct {
	Width     int
	Separator rune
	LineColor Color
	TextColor Color
	line      *string
	text *string
}

//Println prints content formatted
func (m *Matrix) Println(s string) {
	s = strPad(s, m.Width, " ", "RIGHT")[0 : m.Width-3]
	text := fmt.Sprintf("%s%s %s%s%s\n", string(m.Separator), m.TextColor, s, NONE, string(m.Separator))
	m.appendText(text)
	fmt.Print(text)
}

//PrettyP prints with color 
func (m *Matrix) PrettyP(s string, color Color){
	m.TextColor = color
	m.Println(s)
}

//Line prints a line with Width width
func (m *Matrix) Line() {
	fmt.Print(m.LineColor)
	fmt.Print(m.getLine())
	fmt.Print(NONE)
}

func (m *Matrix) getLine() string {
	if m.line == nil {
		s := string(m.Separator)
		line := strPad(s, m.Width, s, "LEFT")
		line = fmt.Sprintf("%s\n",line)
		m.line = &line
	}
	m.appendText(*m.line)
	return *m.line
}

func (m *Matrix) appendText(text string) {
	aux := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(text,ValueOf(m.TextColor),""),ValueOf(NONE),""),ValueOf(m.LineColor),"")
	if m.text == nil {
		m.text = &aux
		return
	}
	aux = fmt.Sprintf("%s%s",*m.text,aux)
	m.text = &aux
}

//GetText returns consolidate text of matrix
func (m *Matrix) GetText() string {
	return *m.text
}

func strPad(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "BOTH":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}
