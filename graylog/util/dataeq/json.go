package dataeq

import (
	"encoding/json"
)

var JSON = New(json.Marshal, json.Unmarshal) //nolint:gochecknoglobals
