package conf

import (
	"controllers"

	"github.com/mrkt/cellgo"
)

func SetController() {
	cellgo.CellCore.RegisterController("user", &controllers.UserController{}, []string{"Run", "Add"})
}
