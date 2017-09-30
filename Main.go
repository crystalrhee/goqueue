// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func testHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
  fs := http.FileServer(http.Dir("FrontEnd"))
	http.Handle("/test/", http.StripPrefix("/test/", fs))

	http.HandleFunc("/createRoom/", testHandler)

	http.ListenAndServe(":8080", nil)
}
