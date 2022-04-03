package user

import "fmt"

type User struct {
	matricule string
	name      string
	firstname string
}

func (u *User) New(matricule string, name string, firstname string) {
	u.matricule = matricule
	u.name = name
	u.firstname = firstname
}

func (u *User) ToString() string {
	return fmt.Sprintf("matricule : %v, name : %v, firstname : %v", u.matricule, u.name, u.firstname)
}
