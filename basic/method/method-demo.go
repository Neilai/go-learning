package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

//若是值调用的话  不会改变原来数据
// func (u user) changeEmail(email string) {
// 	u.email = email
// }

// main is the entry point for the application.
func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	bill.changeEmail("bill@newdomain.com")
	bill.notify()
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
