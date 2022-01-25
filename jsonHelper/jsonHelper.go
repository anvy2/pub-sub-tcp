package jsonHelper

import (
	"encoding/json"
	"time"
)

type JSON json.RawMessage

type ChannelData struct {
	Data      string
	Timestamp time.Time
}

// EncodeToJson Takes a interface and convert it to JSON type which is represented in []byte by json.RawMessage
func EncodeToJson(o interface{}) (JSON, error) {
	result, err := json.Marshal(o)
	return result, err
}

// DecodeJson Takes json.RawMessage and export it to interface type given
func DecodeJson(jsonData JSON, o interface{}) error {
	return json.Unmarshal(jsonData, o)
}

// IsJSON Check if the given string is valid json
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
