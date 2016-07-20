package service

import (
	"fmt"
	"library/dao"

	"github.com/mrkt/cellgo"
)

type UserService struct {
	cellgo.Service
}

func (this *UserService) Test() {
	fmt.Println("test service runing...")
	this.GetDao(&dao.UserDao{})

}
