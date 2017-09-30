package main

import "fmt"
import "User.go"

type IRoom interface {
  printPermissions()
}

type Room struct {
  Users     []IUser
  Name      string
  Password  string
}

func (r Room) printRoom() {
    fmt.Println("Name:", u.Name)
}

func main() {
    u := User{ Permissions { Name: "Jason", Muted: true, ConnectedAccount: true, CanSkip: true }}
    r := Room{ Users: [1]Users{u}, Name: "The Doom Room", Password: "you can't escape the doom room" }
    r.printPermissions()
}
