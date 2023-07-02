package entity

type User struct {
	Id UserId
	Name
	Password
	Rule
}

type NamePass struct {
	Name
	Password
}

type AuthenticationAble struct {
	Password
	Rule
}

type UserId int64
type Password string
type Rule []string
type Name string
