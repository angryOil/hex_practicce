package user

import (
	"errors"
	"hex_practicce/adaptor/out/repository"
	"hex_practicce/domain/entity"
	"reflect"
	"testing"
)

type mockUserRepo struct {
}

func (mockUserRepo) Insert(user entity.User) (error, entity.User) {
	return nil, entity.User{
		Id:       100,
		Name:     "tes",
		Password: "asdf1234",
		Rule: entity.Rule{
			"Admin",
		},
	}
}

func TestUseCase_signOn(t *testing.T) {
	var mock mockUserRepo

	type fields struct {
		userRepo repository.UserIRepository
	}
	type args struct {
		name     string
		password string
		rule     []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
		want1  int64
	}{
		{"success", fields{mock}, args{
			name:     "Admin",
			password: "asdf1234",
			rule: entity.Rule{
				"Admin",
			},
		}, nil, 100},
		{
			"noName", fields{mock}, args{
				name:     "",
				password: "1234",
				rule:     entity.Rule{"nillll"},
			},
			errors.New("이름이 없습니다"), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := UseCase{
				Repo: tt.fields.userRepo,
			}
			got, got1 := c.SignOn(tt.args.name, tt.args.password, tt.args.rule)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignOn() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SignOn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
