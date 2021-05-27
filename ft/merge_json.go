package ft

import (
	"encoding/json"
)

//MergeStatus represents the result of merge
type MergeStatus struct {
	NewKeys map[string]interface{}
	MissingKeys map[string]interface{}
	
}

//MergeJson receive 2 json content  and try merge it
func MergeJson(jsonStrWeak, jsonStrStrong string) (status MergeStatus, err error) {
	jsonWeak := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonStrWeak), &jsonWeak)
	if err != nil {
		return MergeStatus{}, err
	}

	jsonStrong := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonStrStrong), &jsonStrong)
	if err != nil {
		return MergeStatus{}, err
	}
	return compare(jsonWeak,jsonStrong), nil
}



func compare(weak, strong map[string]interface{}) (status MergeStatus){
	status = MergeStatus{}
	status.NewKeys = make(map[string]interface{})
	status.MissingKeys = make(map[string]interface{})
	for k, v := range weak {
		_, ok := strong[k]
		if !ok {
			status.NewKeys[k] = v
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
