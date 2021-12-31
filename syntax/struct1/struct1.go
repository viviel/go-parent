package struct1

type User struct {
	Name string
	Age  int
}

func (u *User) M1() {
	u.Name = u.Name + "_m1"
	u.Age = u.Age + 1
}

func (u User) M2() {
	u.Name = u.Name + "_m2"
	u.Age = u.Age + 1
}
