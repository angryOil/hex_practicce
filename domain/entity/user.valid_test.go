package entity

import "testing"

func TestAuthenticationAble_CheckAdminAuthentication(t *testing.T) {
	type fields struct {
		Password Password
		Rule     Rule
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				Password: "Admin",
				Rule:     Rule{"Admin"},
			},
			want: true,
		},
		{
			name: "passNotValid",
			fields: fields{
				Password: "hello",
				Rule:     Rule{"Admin"},
			},
			want: false,
		}, {
			name: "ruleNotValid",
			fields: fields{
				Password: "Admin",
				Rule:     Rule{"아무거나"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AuthenticationAble{
				Password: tt.fields.Password,
				Rule:     tt.fields.Rule,
			}
			if got := a.CheckAdminAuthentication(); got != tt.want {
				t.Errorf("CheckAdminAuthentication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword_valid(t *testing.T) {
	tests := []struct {
		name   string
		p      Password
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := tt.p.valid(); gotOk != tt.wantOk {
				t.Errorf("valid() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestRule_valid(t *testing.T) {
	tests := []struct {
		name string
		r    Rule
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.valid(); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {
	type fields struct {
		Id       UserId
		Name     Name
		Password Password
		Rule     Rule
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Id:       0,
				Name:     "정온유",
				Password: "joyworld",
				Rule:     Rule{"Admin"},
			},
			wantErr: false,
		}, {
			name: "noName",
			fields: fields{
				Id:       0,
				Name:     "",
				Password: "123",
				Rule:     Rule{"ny"},
			},
			wantErr: true,
		}, {
			name: "noPassword",
			fields: fields{
				Id:       0,
				Name:     "joy",
				Password: "",
				Rule:     Rule{"룰!"},
			},
			wantErr: true,
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
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
