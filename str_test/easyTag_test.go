package str

import (
	"testing"

	"github.com/andersonlira/goutils/str"
)

func TestEasyTag(t *testing.T) {
	tag := "A,BC,'CDE,FGH',IJK"
	tags := str.EasyTag(tag, ",", "'")

	if len(tags) != 4 {
		t.Errorf("Tag size should be %d but %d", 4, len(tags))
	}
	if tags[0] != "CDE,FGH" || tags[1] != "A" || tags[2] != "BC" || tags[3] != "IJK" {
		t.Errorf("Tags were not separeted properly %v", tags)
	}

	tag = "'ABC,CDE','FGH,IJK'"
	tags = str.EasyTag(tag, ",", "'")

	if len(tags) != 2 {
		t.Errorf("Tag size should be %d but %d", 2, len(tags))
	}
	if tags[0] != "ABC,CDE" || tags[1] != "FGH,IJK" {
		t.Errorf("Tags were not separeted properly %v", tags)
	}

	tag = "'ABC,CDE','FGH,IJK"
	tags = str.EasyTag(tag, ",", "'")

	if len(tags) != 3 {
		t.Errorf("Tag size should be %d but %d", 3, len(tags))
	}
	if tags[0] != "ABC,CDE" || tags[1] != "'FGH" || tags[2] != "IJK" {
		t.Errorf("Tags were not separeted properly %v", tags)
	}

	tag = "|ABC;CDE|;|FGH;IJK"
	tags = str.EasyTag(tag, ";", "|")

	if len(tags) != 3 {
		t.Errorf("Tag size should be %d but %d", 3, len(tags))
	}
	if tags[0] != "ABC;CDE" || tags[1] != "|FGH" || tags[2] != "IJK" {
		t.Errorf("Tags were not separeted properly %v", tags)
	}

}
