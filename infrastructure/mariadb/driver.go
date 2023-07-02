package mariadb

import "hex_practicce/domain/entity"

type Mariadb struct {
}

func NewMaria() Mariadb {
	return Mariadb{}
}

func (receiver Mariadb) Insert(user entity.User) (error, entity.User) {
	return nil, entity.User{}
}
