package str

import (
	"testing"

	"github.com/andersonlira/goutils/str"
)

func TestEasyTag(t *testing.T) {
	tag := "ABC,'CDE,FGH',IJK"
	tags := str.EasyTag(tag, ",", "'")

	if len(tags) != 3 {
		t.Errorf("Tag size should be %d but %d", 3, len(tags))
	}
	if tags[0] != "ABC" || tags[1] != "CDE,FGH" || tags[2] != "IJK" {
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
		t.Errorf("Tag size should be %d but %d", 2, len(tags))
	}
	if tags[0] != "ABC,CDE" || tags[1] != "FGH,IJK" {
		t.Errorf("Tags were not separeted properly %v", tags)
	}

}
