package Models

import "fmt"

type IUser interface {
  printPermissions()
}

type Permissions struct {
  Name              string
  Muted             bool
  ConnectedAccount  bool
  CanSkip           bool
}

type User struct {
  Permissions
}

type Moderator struct {
  User
}

func (u User) printPermissions() {
    fmt.Println("Name:", u.Name)
}
