package main

import "fmt"

type Permissions struct {
  Name              string
  Muted             bool
  ConnectedAccount  bool
  CanSkip           bool
}

type IUser interface {
  printPermissions()
}

type User struct {
  Permissions
}

type Moderator struct {
  User
}

func (u User) printPermissions() {
  for
    return fmt.Println("%s: %s", part.Name, part.Description)
}

func main() {
    fmt.Printf("hello, world\n")
}
nn
