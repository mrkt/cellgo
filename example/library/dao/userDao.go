package dao

import (
	"github.com/mrkt/cellgo"
)

type UserDao struct {
	cellgo.Dao
}

func (this *UserDao) Before() {}

func (this *UserDao) GetUserInfoList(name interface{}) map[string]string {
	userInfo := make(map[string]string)
	if names, _ := name.(string); names == "tommy" {
		userInfo["Username"] = "dao://tommy.jin"
		userInfo["Email"] = "dao://tommy.jin@aliyun.com"
	}
	return userInfo
}
