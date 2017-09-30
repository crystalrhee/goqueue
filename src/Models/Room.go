package Models

import (
  "fmt"
)

type IRoom interface {
  printPermissions()
}

type Room struct {

  Users           []User
  RoomName        string
  RoomPassword    string
  SpotfiyUsername string
  SpotifyPassword string
}

func (r Room) PrintRoom() {
  fmt.Println("Name:", r.Name)
}
