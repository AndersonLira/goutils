package recode

import 
(
	"strings"
	"github.com/andersonlira/goutils/io"

)

//Coder is the object that perform code modifications. Use MakeCoder to get a object reference
type Coder struct {
	FilePath string
	Lines []string
}

//MakeCoder creates a object Coder and initialize it.
//Param filePath the filePath
func MakeCoder(filePath string) (Coder, error) {
	content, err := io.ReadFile(filePath)
	coder := Coder{}
	if err != nil {
		return coder, err
	}

	coder.FilePath = filePath
	coder.Lines = strings.Split(content,"\r\n")
	if(len(coder.Lines) == 1) {
		coder.Lines = strings.Split(content,"\n")
	}
	return coder,nil

}

//AddAfterLine add lines of code after giving expression
//after string means the code will be written only after that occurrence is matched
//before if before is diferent of '', the code will be write only that occurrence is matched
//content the lines that should be added
func (c *Coder) AddAfterLine(after string,before string,content ...string) bool {
	newContent := []string{}
	p0 := false
	added := false
	for _, line := range c.Lines {
		line = strings.ReplaceAll(strings.ReplaceAll(line,"\n",""),"\r","")
		if !p0 {
			if strings.Contains(line,after) {
				p0 = true
			}
		} else {
			if !added && (before == "" || strings.Contains(line,before))   {
				added = true
				for _, cont := range content {
					newContent = append(newContent,cont + "\n")
				}
			}
		}
		newContent = append(newContent,line + "\n")

	}
	c.Lines = newContent
	return added
}

//NewCodeContent returns the code changed 
func (c *Coder) NewCodeContent() string {
	aux := ""
	for _, line := range c.Lines {
		aux = aux + line
	}
	return aux
}