package main

import (
	"experiment/user"
	"fmt"
)

func main() {
	var u = new(user.User)

	u.New("123456", "Toto", "Titi")
	fmt.Println(u.ToString())
}
