package entity

func (u User) GetAuthenticationAble() AuthenticationAble {
	return AuthenticationAble{
		Password: u.Password,
		Rule:     u.Rule,
	}
}
