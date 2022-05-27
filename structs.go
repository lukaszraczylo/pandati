package pandati

import (
	"encoding/json"

	"github.com/wI2L/jsondiff"
)

type CompareStructReplacedResult struct {
	Key      string
	Value    interface{}
	OldValue interface{}
}

// CompareStructReplaced compares two structs after converting them into json
// and returns:
// - array of differences
// - bool true if there's no differences
// - error if such occurs
func CompareStructsReplaced(a, b interface{}) ([]CompareStructReplacedResult, bool, error) {
	res := []CompareStructReplacedResult{}

	aJson, err := json.Marshal(a)
	if err != nil {
		return res, false, err
	}
	bJson, err := json.Marshal(b)
	if err != nil {
		return res, false, err
	}

	diffTotal, err := jsondiff.CompareJSONOpts(aJson, bJson, jsondiff.Invertible())
	if err != nil {
		return nil, false, err
	}

	for _, diff := range diffTotal {
		if diff.Type == "replace" {
			if IsZero(diff.Path) {
				for k, v := range diff.Value.(map[string]interface{}) {
					res = append(res, CompareStructReplacedResult{
						Key:      k,
						Value:    v,
						OldValue: diff.OldValue.(map[string]interface{})[k],
					})
				}
			} else {
				res = append(res, CompareStructReplacedResult{
					Key:      diff.Path.String(),
					Value:    diff.Value,
					OldValue: diff.OldValue,
				})
			}
		}
	}

	return res, IsZero(res), nil
}
