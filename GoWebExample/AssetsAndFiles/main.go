// static-files.go
package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("/Users/wyy/code/github/go/SGG-Go-lesson/GoWebExample/AssetsAndFiles/assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}