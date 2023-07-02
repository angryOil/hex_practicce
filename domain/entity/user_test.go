package entity

import (
	"reflect"
	"testing"
)

func TestUser_GetAuthenticationAble(t *testing.T) {
	type fields struct {
		Id       UserId
		Name     Name
		Password Password
		Rule     Rule
	}
	tests := []struct {
		name   string
		fields fields
		want   AuthenticationAble
	}{
		{
			name: "success",
			fields: fields{
				Id:       0,
				Name:     "joy",
				Password: "1234",
				Rule:     nil,
			},
			want: AuthenticationAble{
				Password("1234"),
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Id:       tt.fields.Id,
				Name:     tt.fields.Name,
				Password: tt.fields.Password,
				Rule:     tt.fields.Rule,
			}
			if got := u.GetAuthenticationAble(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthenticationAble() = %v, want %v", got, tt.want)
			}
		})
	}
}
