package user

import (
	"encoding/json"
	"hex_practicce/domain/service/user"
	"net/http"
	"strconv"
)

type Controller struct {
	UserService user.IService
}

type IController interface {
}

func NewController(service user.IService) Controller {
	return Controller{service}
}

func (c Controller) SignOn(w http.ResponseWriter, r *http.Request) {
	json.Marshal(r.Body)
	err1, id := c.UserService.SignOn("joy", "1234", nil)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Success"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(int(id))))
}
