package entity

import "errors"

func (r Rule) valid() bool {
	for _, s := range r {
		if s == "Admin" {
			return true
		}
	}
	return false
}

func (p Password) valid() (ok bool) {
	return p == "Admin"
}

func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("이름이 없습니다")
	}
	if u.Password == "" {
		return errors.New("비밀번호가 없습니다.")
	}
	return nil
}

func (a AuthenticationAble) CheckAdminAuthentication() bool {
	if !a.Password.valid() {
		return false
	}
	return a.Rule.valid()
}
