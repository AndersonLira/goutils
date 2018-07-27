package str

import "strings"

//EasyTag return a tag array when receive string field
//Example EasyTag("A,'B,C',D",",","'") return "A","B,C","D" values
func EasyTag(s, separator, textMark string) (tags []string) {
	aux := strings.Split(s, textMark)

	return
}
