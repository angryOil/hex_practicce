package application

import "hex_practicce/adaptor/in/controller/rest/user"

func Run() user.Controller {
	a := InitializeController()
	return a
}
