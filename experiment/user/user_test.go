package user

import (
	"reflect"
	"strings"
	"testing"
)

func TestUserCreation(t *testing.T) {
	user := User{
		matricule: "123456",
		name:      "name",
		firstname: "firstname",
	}

	userTest := new(User)
	userTest.New("123456", "name", "firstname")

	want := reflect.DeepEqual(&user, userTest)
	if !want {
		t.Fatalf("Error not equal")
	}
}

func TestUserToString(t *testing.T) {
	user := User{
		matricule: "123456",
		name:      "name",
		firstname: "firstname",
	}

	want := strings.Compare("matricule : 123456, name : name, firstname : firstname", user.ToString()) == 0
	if !want {
		t.Fatalf("Error not equal")
	}
}
