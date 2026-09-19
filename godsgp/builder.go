package main

import "fmt"

type User struct {
	Name     string
	Age      int
	Email    string
	City     string
	IsActive bool
}

type UserBuilder struct {
	user User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: User{
			IsActive: true,
		},
	}
}

func (b *UserBuilder) Name(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) Age(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) Email(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) City(city string) *UserBuilder {
	b.user.City = city
	return b
}

func (b *UserBuilder) Active(active bool) *UserBuilder {
	b.user.IsActive = active
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}

func main() {
	user := NewUserBuilder().
		Name("John").
		Age(30).
		Email("john@example.com").
		City("Delhi").
		Build()

	fmt.Printf("%+v\n", user)
}

/*

Builder is useful when an object has many optional parameters.

Instead of:

User("John", 30, "Delhi", "123", true, ...)
we use:

NewUserBuilder().
	Name("John").
	Age(30).
	City("Delhi").
	Build()


*/
