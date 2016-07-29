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

func TestAuthcode(t *testing.T) {

	var (
		value []byte = []byte("TOMMY")
		//value     []byte = []byte("KMKiARpvwqUsw44iFsKWwrvDvQ==")
		operation string = "ENCODE"
		hashkey   string = "9597f4KpYTsJ5tD6"
	)
	auth_key := If(hashkey != "", hashkey, HashKey).(string)

	h := md5.New()
	h.Write([]byte(auth_key)) // md5加密
	cipherStr := h.Sum(nil)
	key := hex.EncodeToString(cipherStr)
	key_length := len(key)

	//str ：= fmt.Sprintf("%d", value)
	var valueStr string
	if operation == "DECODE" {
		//fmt.Println(value)
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
		//fmt.Println(valueStr)
		//fmt.Println(temprune)
		//fmt.Println(valueStr_length)
		//fmt.Println(z)
		//fmt.Println(fmt.Sprintf("%x", temprune[z]))
		tempInt, err := strconv.ParseInt(fmt.Sprintf("%x", temprune[z]), 16, 10)
		if err != nil {
			t.Fatal("strconv.ParseInt", err)
		}
		//fmt.Println(tempInt)
		str := fmt.Sprintf("%c", int(tempInt)^(box[(box[x]+box[y])%256]))
		//fmt.Println(str)
		//temp16 := strconv.FormatInt(int64(int(tempInt)^(box[(box[x]+box[y])%256])), 16)
		//fmt.Println(fmt.Sprintf("%c", temp16))
		result += str
	}

	if operation == "DECODE" {
		//fmt.Println(result)
		/*var tempStr string
		for _, v := range result {
			temp, err := strconv.ParseInt(fmt.Sprintf("%s", v), 16, 10)
			if err != nil {
				fmt.Println("strconv.Atoi", err)
			}
			tempStr += fmt.Sprintf("%s", temp)
		}
		fmt.Println(tempStr)*/
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
	t.Fatal(Authcode([]byte("KMKiARpvwqUsw44iFsKWwrvDvQ=="), "DECODE", "9597f4KpYTsJ5tD6"))
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
