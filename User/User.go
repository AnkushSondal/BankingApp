package user

import uuid "github.com/satori/go.uuid"

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	UserName  string
}
var allUsers []*User

func NewUser(firstName, lastName, username string) *User {
	newUser := &User{
		ID: uuid.NewV4(),
		FirstName: firstName,
		LastName:  lastName,
		UserName:  username,
	}
	allUsers = append(allUsers, newUser)
	return newUser
}

func findUser(users []*User, username string) (*User,bool){
	for i := 0; i < len(users); i++ {
		// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>", users[i].userName, username)
		if users[i].UserName == username {
			return users[i], true
		}
	}
	return nil, false
}

