package repository

import (
	"hex_practicce/domain/entity"
	"hex_practicce/infrastructure/mariadb"
)

type UserIRepository interface {
	Insert(user entity.User) (error, entity.User)
}

type UserRepository struct {
}

func NewRepo() UserIRepository {
	return mariadb.NewMaria()
}

func (u UserRepository) Insert(user entity.User) (error, entity.User) {
	return nil, entity.User{}
}
