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

func (r Room) PrintRoom() {
    fmt.Println("Name:", r.Name)
}
