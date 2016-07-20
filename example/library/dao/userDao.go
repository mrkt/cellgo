package dao

import (
	"fmt"

	"github.com/mrkt/cellgo"
)

type UserDao struct {
	cellgo.Dao
}

func (this *UserDao) Test() {
	fmt.Println("test dao runing...")
}
