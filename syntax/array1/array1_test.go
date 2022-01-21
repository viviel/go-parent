package array1_test

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	type a []string
	as := make(a, 0, 2)
	fmt.Println(len(as))
	as = append(as, "1")
	fmt.Println(len(as))
}
