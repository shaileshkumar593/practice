package main

type Stringer interface{ String() string }

type User struct{ Name string }

func (u *User) String() string { return u.Name }

var _ Stringer = User{Name: "Ada"}
var _ Stringer = &User{Name: "Ada"}

func main()
