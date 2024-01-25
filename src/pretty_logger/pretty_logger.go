package prettylogger

import (
	"bytes"
	"encoding/json"
)

// To check if the response if correct or not you can use
// this extension for showing json log
func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
