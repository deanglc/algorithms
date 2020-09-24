package main

import "fmt"

/*
s = "3[a]2[bc]"
aaabcbc

*/
func main() {
	a := "Randal"
	for i := 0; i < len(a); i++ {
		fmt.Printf("%x ", a[i])
		fmt.Printf("%c ", a[i])
	}
	for k, v := range a {
		fmt.Println(k, v)
	}
	b := []rune("公公哦那个")
	for k, v := range b {
		fmt.Println(k, v)
	}
}

// func decodeString(s string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}
// 	stack := make([]byte,0)
// 	for _ , val := range s {
// 		switch val {
// 		case "]":
// 			temp :=
//
// 		}
// 	}
// }
