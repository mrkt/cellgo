package service

import (
	"library/dao"

	"github.com/mrkt/cellgo"
)

type UserService struct {
	cellgo.Service
}

func (this *UserService) Before() {
	//init dao
	this.GetDao(&dao.UserDao{})
}

func (this *UserService) GetUserInfo(name interface{}) map[string]string {
	userInfo, _ := this.GetDaoFunc("GetUserInfoList", name).(map[string]string)
	for k, v := range userInfo {
		userInfo[k] = "service://" + v
	}
	return userInfo
}
