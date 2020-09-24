package main

import "runtime"

// package glcHttp
//
// import "net/http"
//
// func main() {
//	http.Client{}
// }

func main() {
	t := 3
	print(t)
	runtime.GC()
}
