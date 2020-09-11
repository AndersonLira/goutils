package recode

import 
(
	"strings"
	"github.com/andersonlira/goutils/io"

)


type Coder struct {
	FilePath string
	Lines []string
}


func MakeCoder(filePath string) (Coder, error) {
	content, err := io.ReadFile(filePath)
	coder := Coder{}
	if err != nil {
		return coder, err
	}

	coder.FilePath = filePath
	coder.Lines = strings.Split(content,"\r\n")
	return coder,nil

}

func (c *Coder) AddAfterLine(after string,before string,content ...string) bool {
	newContent := []string{}
	p0 := false
	added := false
	for _, line := range c.Lines {
		if !p0 && strings.Contains(line,after){
			p0 = true
		}
		if p0 {
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

func (c *Coder) NewCodeContent() string {
	aux := ""
	for _, line := range c.Lines {
		aux = aux + line
	}
	return aux
}