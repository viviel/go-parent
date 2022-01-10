package interface1_test

import (
	"fmt"
	"go_parent/syntax/interface1"
	"testing"
)

func Test1(t *testing.T) {
	var u interface1.User
	u = interface1.User{Name: "vv", Age: 12}
	r := u.Info()
	fmt.Println(r)
}

func Test2(t *testing.T) {
	var i interface1.Info
	i = &interface1.User{}
	r := i.Info()
	fmt.Println(r)
}
