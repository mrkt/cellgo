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
//| Cellgo Framework session/encrypt_test file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-30

package session

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
)

/*func TestAuthcode(t *testing.T) {

	var (
		//value     []byte = []byte("TOMMY")
		value     []byte = []byte("esO8Uk9vwqsmwpk=")
		operation string = "DECODE"
		hashkey   string = "9597f4KpYTsJ5tD6"
	)
	auth_key := If(hashkey != "", hashkey, HashKey).(string)

	h := md5.New()
	h.Write([]byte(auth_key)) // md5加密
	cipherStr := h.Sum(nil)
	key := hex.EncodeToString(cipherStr)
	key_length := len(key)

	var valueStr string
	if operation == "DECODE" {
		temp, err := Decode(value)
		if err != nil {
			t.Fatal("decode error", err)
		}
		valueStr = fmt.Sprintf("%s", temp)

	} else {
		h.Write([]byte(fmt.Sprintf("%s", value) + key)) // md5加密
		cipherStr = h.Sum(nil)
		valueStr = Substr(hex.EncodeToString(cipherStr), 0, 8) + fmt.Sprintf("%s", value)

	}

	var (
		rndkey [256]int
		box    [256]int
		result string
		keys   []rune
	)

	keys = []rune(key)
	for i := 0; i <= 255; i++ {
		tempInt, err := strconv.ParseInt(fmt.Sprintf("%x", keys[i%key_length]), 16, 10)
		if err != nil {
			t.Fatal("strconv.ParseInt", err)
		}
		rndkey[i] = int(tempInt)
		box[i] = i
	}

	for k, j := 0, 0; j < 256; j++ {
		k = (k + box[j] + rndkey[j]) % 256
		box[j], box[k] = box[k], box[j]
	}
	temprune := []rune(valueStr)
	valueStr_length := len(temprune)
	for x, y, z := 0, 0, 0; z < valueStr_length; z++ {
		x = (x + 1) % 256
		y = (y + box[x]) % 256
		box[x], box[y] = box[y], box[x]
		tempInt, err := strconv.ParseInt(fmt.Sprintf("%x", temprune[z]), 16, 10)
		if err != nil {
			t.Fatal("strconv.ParseInt", err)
		}
		str := fmt.Sprintf("%c", int(tempInt)^(box[(box[x]+box[y])%256]))
		result += str
	}

	if operation == "DECODE" {
		h.Write([]byte(Substr(result, 8, 0) + key)) // md5加密
		cipherStr = h.Sum(nil)
		if Substr(result, 0, 8) == Substr(hex.EncodeToString(cipherStr), 0, 8) {
			t.Fatal(Substr(result, 8, 0))
		} else {
			t.Fatal("null value")
		}
	} else {
		t.Fatal(fmt.Sprintf("%s", Encode([]byte(result))))

	}

}*/

func TestAuthcode(t *testing.T) {

	var (
		value []byte = []byte("TOMMY")
		//value     []byte = []byte("esO8Uk9vwqsmwpk=")
		operation string = "ENCODE"
		hashkey   string = "9597f4KpYTsJ5tD6"
	)
	auth_key := If(hashkey != "", hashkey, HashKey).(string)

	h := md5.New()
	h.Write([]byte(auth_key)) // md5加密
	cipherStr := h.Sum(nil)
	key := hex.EncodeToString(cipherStr)
	key_length := len(key)

	var valueStr string
	if operation == "DECODE" {
		temp, err := Decode(value)
		if err != nil {
			t.Fatal("decode error", err)
		}
		valueStr = fmt.Sprintf("%s", temp)

	} else {
		h.Write([]byte(fmt.Sprintf("%s", value) + key)) // md5加密
		cipherStr = h.Sum(nil)
		valueStr = Substr(hex.EncodeToString(cipherStr), 0, 8) + fmt.Sprintf("%s", value)

	}

	var (
		rndkey [256]int
		box    [256]int
		result string
		keys   []rune
	)

	keys = []rune(key)
	for i := 0; i <= 255; i++ {
		tempInt, err := strconv.ParseInt(fmt.Sprintf("%x", keys[i%key_length]), 16, 10)
		if err != nil {
			t.Fatal("strconv.ParseInt", err)
		}
		rndkey[i] = int(tempInt)
		box[i] = i
	}

	for k, j := 0, 0; j < 256; j++ {
		k = (k + box[j] + rndkey[j]) % 256
		box[j], box[k] = box[k], box[j]
	}
	temprune := []rune(valueStr)
	valueStr_length := len(temprune)
	for x, y, z := 0, 0, 0; z < valueStr_length; z++ {
		x = (x + 1) % 256
		y = (y + box[x]) % 256
		box[x], box[y] = box[y], box[x]
		tempInt, err := strconv.ParseInt(fmt.Sprintf("%x", temprune[z]), 16, 10)
		if err != nil {
			t.Fatal("strconv.ParseInt", err)
		}
		str := fmt.Sprintf("%c", int(tempInt)^(box[(box[x]+box[y])%256]))
		result += str
	}

	if operation == "DECODE" {
		h.Write([]byte(Substr(result, 8, 0) + key)) // md5加密
		cipherStr = h.Sum(nil)
		if Substr(result, 0, 8) == Substr(hex.EncodeToString(cipherStr), 0, 8) {
			t.Fatal(Substr(result, 8, 0))
		} else {
			t.Fatal("null value")
		}
	} else {
		t.Fatal(fmt.Sprintf("%s", Encode([]byte(result))))

	}

}

func TestAuthcodeTrue(t *testing.T) {

	t.Fatal(Authcode([]byte("LcO9UEtpw7Etw4YyL8O2wrTDpsKkf8Orw5lkw4JJwqPCv3RMw47CrTnCpcOFw4fCuFYceFPDrsOBSHo1wqMnwq0jHMKSFcKtdAZdVk3Ck3k4wqHDuSjDoFk8wp05YyrCmHjDjMK3wr9hw6vDsQULw7_CqD0XWSbDmcO-w5ckL2stw6DCmsKkwoHDqS3DpcKRwr0xfsKJw6pWw782wq4FfycSwow2wrVhdA7DgcO9XytcKX86wogGZmQqw5VafwNbAAnDlQA6w7zCj8KwCB99w57DpsO9JnHCqSsoGG5RG8KkwpHCog=="), "DECODE", "9597f4KpYTsJ5tD6"))
}

// encode encodes a value using base64.
func Encode(value []byte) []byte {
	encoded := make([]byte, base64.URLEncoding.EncodedLen(len(value)))
	base64.URLEncoding.Encode(encoded, value)
	return encoded
}

// decode decodes a cookie using base64.
func Decode(value []byte) ([]byte, error) {
	decoded := make([]byte, base64.URLEncoding.DecodedLen(len(value)))
	b, err := base64.URLEncoding.Decode(decoded, value)
	if err != nil {
		return nil, err
	}
	return decoded[:b], nil
}
