package ft

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

//MergeStatus represents the result of merge
type MergeStatus struct {
	NewKeys map[string]string
	MissingKeys map[string]string
	DiffKeys []Diff
	final string
}

func (ms MergeStatus) FinalJSON() string {
	for _, diff := range ms.DiffKeys {
		rgx := fmt.Sprintf(`("%s".*:.*")%s`,escape(diff.Key),escape(diff.WeakValue))
		m := regexp.MustCompile(rgx)
		repl := fmt.Sprintf("${1}%s",preserveBreakLine(diff.StrongValue))
		ms.final = m.ReplaceAllString(ms.final, repl)
	}
	return ms.final
}
//TODO make a independent api
func escape(s string) string {
	chars := []string{"(",")","[","]","|","?",".","*","$"}
	esc := s

	for _,c := range chars {
		esc = strings.ReplaceAll(esc,c,fmt.Sprintf("\\%s",c))
	}
	return esc
}

func preserveBreakLine(s string) string {
	return strings.ReplaceAll(s,"\n","\\n")	
}

//Diff reprents when two values are different on given values
type Diff struct {
	Key string
	WeakValue string
	StrongValue string
}

//MergeJson receive 2 json content  and try merge it. It is implemented for simple 
//JSON formats. key value only
func MergeJson(jsonStrWeak, jsonStrStrong string) (status MergeStatus, err error) {
	jsonWeak := make(map[string]string)
	err = json.Unmarshal([]byte(jsonStrWeak), &jsonWeak)
	if err != nil {
		return MergeStatus{}, err
	}

	jsonStrong := make(map[string]string)
	err = json.Unmarshal([]byte(jsonStrStrong), &jsonStrong)
	if err != nil {
		return MergeStatus{}, err
	}
	status = compare(jsonWeak,jsonStrong)
	status.final = jsonStrWeak
	return status,nil
}



func compare(weak, strong map[string]string) (status MergeStatus){
	status = MergeStatus{}
	status.NewKeys = make(map[string]string)
	status.MissingKeys = make(map[string]string)
	status.DiffKeys = []Diff{}
	for k, wv := range weak {
		sv, ok := strong[k]
		if !ok {
			status.NewKeys[k] = wv
		}else if sv != wv{
			status.DiffKeys = append(status.DiffKeys,Diff{k,wv,sv})
		}
	}

	for k, v := range strong {
		_, ok := weak[k]
		if !ok {
			status.MissingKeys[k] = v
		}
	}

	return status
}
