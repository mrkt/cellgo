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
//| Cellgo Framework tool/urlcode file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-03

package tool

import (
	"net/url"
)

var Urlcode = &urlcodeTool{}

type urlcodeTool struct {
}

func (u *urlcodeTool) Encode(value string) string {
	encoded := url.QueryEscape(value)
	return encoded
}

func (u *urlcodeTool) Decode(value string) (string, error) {
	decoded, err := url.QueryUnescape(value)
	if err != nil {
		return "", err
	}
	return decoded, nil
}
