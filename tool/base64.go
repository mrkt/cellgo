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
//| Cellgo Framework tool/base64 file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-03

package tool

import (
	"encoding/base64"
)

var Base64 = &base64Tool{}

type base64Tool struct {
}

// encode encodes a value using base64.
func (bt *base64Tool) Encode(value []byte) string {
	encoded := base64.StdEncoding.EncodeToString(value)
	return encoded
}

// decode decodes a value using base64.
func (bt *base64Tool) Decode(value string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
