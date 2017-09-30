// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "Models"
	"net/http"
  "math/rand"
)

const Rooms = map[string]Models.Room
const KEY_LENGTH = 10
const letterRunes = []string("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func uniqueKey(key) {
  recurse := false
  for k := range Rooms {
    if (k == key) {
        recurse := true
      }
  }
  if (recurse) {
    return uniqueKey(genRoomKey(KEY_LENGTH))
  } else {
    return key
  }
}

func genRoomKey(n int) {
  b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return uniqueKey(string(b))
}

func createRoom(w http.ResponseWriter, r *http.Request) {
  // Creates user and room with a randomly generated key
  u := Models.User{ Name: r.UserName, ConnectedAccount: true}
  key := genRoomKey(KEY_LENGTH)
  Rooms[key] = Models.Room{ Name: r.RoomName, Password: r.RoomPassword, Users: []Models.User{u},
    SpotifyPassword: r.SpotifyPassword, SpotfiyUsername: r.SpotfiyUsername}

  roomCookie := http.Cookie{ Name: "goqueueRoomKey", Value: key}
  userCookie := http.Cookie{ Name: "goqueueUName", Value: r.UserName}
  http.SetCookie(w, roomCookie)
  http.SetCookie(w, userCookie)

  http.Redirect(w, r, "/room/", http.StatusFound)
}

func deleteRoom(w http.ResponseWriter, r *http.Request) {
  delete(Rooms, r.RoomKey)
}

func leaveRoom(w http.ResponseWriter, r *http.Request) {
  users := Rooms[r.RoomKey].Users
  if len(users) == 1 {
    deleteRoom(w, r)
  } else {
    for i := range users {
      if users[i].Name == r.UserName {
        users[i] = users[0]
        Rooms[r.RoomKey].Users := users[1:]
      }
    }
  }
}

func getRoom(w http.ResponseWriter, r *http.Request) {
  for _, cookie := range r.Cookies() {
    cookie.Value
    fmt.Fprint(w, cookie.Name)
  }
}

func addUserToRoom(w http.ResponseWriter, r *http.Request) {
  user = Models.User{ Name: r.UserName, ConnectedAccount: false}
  append(Rooms[r.RoomKey].Users, user)
}

func main() {
  rand.Seed(time.Now().UnixNano())
  fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/test/", http.StripPrefix("/test/", fs))

	http.HandleFunc("/room/create/", createRoom)

	http.ListenAndServe(":8080", nil)
}
