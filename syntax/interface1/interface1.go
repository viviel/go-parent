package interface1

import (
	"go_parent/syntax/struct1"
	"strconv"
)

type User struct1.User

type Info interface {
	Info() string
}

func (u *User) Info() string {
	return u.Name + strconv.Itoa(u.Age)
}
