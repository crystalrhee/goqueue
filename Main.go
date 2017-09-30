// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
  "Models"
)

func createRoom(w http.ResponseWriter, r *http.Request) {
  user := Models.User{ Name: "Jason", ConnectedAccount: true }
  room := Models.Room{ Users: []Models.User{user}, Name: "The Doom Room", Password: "you can't escape the doom room" }
  room.PrintRoom()
}

func main() {
  fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/test/", http.StripPrefix("/test/", fs))

	http.HandleFunc("/room/create/", createRoom)

	http.ListenAndServe(":8080", nil)
}
