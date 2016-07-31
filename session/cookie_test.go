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
//|------------------------------------------------------------------
//| Cellgo Framework session/cookie_test file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-31

package session

import (
	"encoding/json"
	"fmt"
	"testing"
)

//cookie base json config
type testConfig struct {
	HashKey    string `json:"hashKey"`    //安全密钥 hash string
	CookieName string `json:"cookieName"` //cookie name
	Secure     bool   `json:"secure"`     //安全与否
	Maxage     int    `json:"maxage"`     //cookie max life time
}

func TestJson(t *testing.T) {
	var cookieName string = "cellcookie"
	var cookieMaxage string = "86400"
	var cookieHashKey string = "9597f4KpYTsJ5tD6"
	testStr := "{\"hashKey\":\"" + cookieHashKey + "\",\"cookieName\":\"" + cookieName + "\",\"secure\":true,\"maxage\":" + cookieMaxage + "}"
	fmt.Println(testStr)
	cf := new(testConfig)
	err := json.Unmarshal([]byte(testStr), cf)
	if err != nil {
		t.Fatal("json.Unmarshal err")
	}
	fmt.Println(cf.Secure)
}
