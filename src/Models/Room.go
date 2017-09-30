package Models

import (
  "fmt"
)

type IRoom interface {
  printPermissions()
}

type Room struct {
  Users     []User
  Name      string
  Password  string
}

func (r Room) printRoom() {
    fmt.Println("Name:", r.Name)
}

func main() {
    u := User{ Permissions { Name: "Jason", Muted: true, ConnectedAccount: true, CanSkip: true }}
    r := Room{ Users: []User{u}, Name: "The Doom Room", Password: "you can't escape the doom room" }
    r.printRoom()
}
