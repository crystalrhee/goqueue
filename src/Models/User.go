package Models

import "fmt"

type IUser interface {
  printPermissions()
}

type User struct {
  Name string
  ConnectedAccount bool
}

type Moderator struct {
  User
}

func (u User) printPermissions() {
    fmt.Println("Name:", u.Name)
}
