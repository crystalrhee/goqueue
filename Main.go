// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
  "Models"
  "fmt"
)

type Page struct {
	Title string
	Body  []byte
}

func createRoom(w http.ResponseWriter, r *http.Request) {
  room = Models.Room{ Model.Users: []User{u}, Name: "The Doom Room", Password: "you can't escape the doom room" }
  room.printRoom()
}

func main() {
  fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/test/", http.StripPrefix("/test/", fs))

	http.HandleFunc("/createRoom/", testHandler)

	http.ListenAndServe(":8080", nil)
}
