// handler.go
package main

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        w.Write([]byte("Hello, Guest!"))
        return
    }
    w.Write([]byte("Hello, " + name + "!"))
}
