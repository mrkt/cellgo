//|------------------------------------------------------------------
//|        __
//|     __/  \
//|  __/  \__/_
//| /  \__/    \
//|/\__/CellGo /_
//|\/_/NetFW__/  \
//|  /\__ _/  \__/
//|  \/_/  \__/_/
//|    /\__/_/
//|    \/_/
//| ------------------------------------------------------------------
//| Cellgo Framework tool/jsoncode file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-03

package tool

import (
	"encoding/json"
)

var Json = &jsonTool{}

type jsonTool struct {
}

// encode struct 2 json.
func (j *jsonTool) Encode(value interface{}) ([]byte, error) {
	encoded, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

// decode json 2 struct.
func (j *jsonTool) Decode(value []byte, decoded interface{}) error {
	err := json.Unmarshal(value, decoded)
	if err != nil {
		return err
	}
	return nil
}
