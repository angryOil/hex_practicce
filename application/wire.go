package application

import (
	"hex_practicce/adaptor/in/controller/rest/user"
	"hex_practicce/adaptor/out/repository"
	user2 "hex_practicce/domain/service/user"
)

func InitializeController() user.Controller {
	userIRepository := repository.NewRepo()
	iService := user2.NewService(userIRepository)
	controller := user.NewController(iService)
	return controller
}
