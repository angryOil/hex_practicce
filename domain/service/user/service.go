package user

import (
	"hex_practicce/adaptor/out/repository"
	"hex_practicce/domain/entity"
)

type UseCase struct {
	Repo repository.UserIRepository
}

type IService interface {
	SignOn(name, password string, rule []string) (error, int64)
}

func NewService(userRepository repository.UserIRepository) IService {
	return UseCase{userRepository}
}

func (c UseCase) SignOn(name, password string, rule []string) (error, int64) {
	// 각 필드 값이 있는지 빈문자등 확인 -> 각 필드를 user entity 로 만들고 user.valid 만들어서 사용한다
	u := toEntity(name, password, rule)

	// 회원 가입을 필요로하는 모든 속성을 검증합니다.
	err1 := valid(u)
	if err1 != nil {
		return err1, 0
	}

	// 저장 요청
	err2, saved := c.Repo.Insert(u)

	return err2, int64(saved.Id)
}

func toEntity(name, password string, rule []string) entity.User {
	return entity.User{
		Id:       entity.UserId(0),
		Name:     entity.Name(name),
		Password: entity.Password(password),
		Rule:     entity.Rule(rule),
	}
}

func valid(u entity.User) error {
	return u.Validate()
}
