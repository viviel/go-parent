package struct1_test

import (
	"fmt"
	"go_parent/syntax/struct1"
	"testing"
)

func Test1(t *testing.T) {
	u := struct1.User{
		Name: "vv",
		Age:  10,
	}
	fmt.Println(u)
	u.M1()
	fmt.Println(u)
	u.M2()
	fmt.Println(u)
}
