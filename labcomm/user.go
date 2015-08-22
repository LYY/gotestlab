package labcomm

import (
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (this User) String() string {
	return fmt.Sprintf("%d|%s", this.Id, this.Name)
}

func InitUsers() []User {
	var uss []User
	for i := 0; i < 1000; i++ {
		u := User{i, fmt.Sprintf("%s%d", "name**", i)}
		uss = append(uss, u)
	}
	return uss
}
