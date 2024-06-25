package utils

import (
	"encoding/json"
)

func ParseBody(body []byte, x interface{}) {
	if err := json.Unmarshal([]byte(body), x); err != nil {
		return
	}
}
