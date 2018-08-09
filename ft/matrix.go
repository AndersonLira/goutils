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
}

//Println prints content formatted
func (m Matrix) Println(s string) {
	s = strPad(s, m.Width, " ", "RIGHT")[0 : m.Width-3]
	fmt.Print(m.TextColor)
	fmt.Println(fmt.Sprintf("%s %s%s", string(m.Separator), s, string(m.Separator)))
}

//Line prints a line with Width width
func (m Matrix) Line() {
	fmt.Print(m.LineColor)
	fmt.Println(m.getLine())
}

func (m Matrix) getLine() string {
	if m.line == nil {
		s := string(m.Separator)
		line := strPad(s, m.Width, s, "LEFT")
		m.line = &line
	}
	return *m.line
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
