package main

import "fmt"

type user struct {
	name  string
	email string
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)

	u3 := &u2
	fmt.Println("V#", u3, &u3)
	u4 := testU(u2)
	fmt.Printf("U44444====%p \n", &u4)

}

func testU(u *user) user {
	fmt.Println("======", &u, u)
	fmt.Printf("======%p \n", &u)
	return *u
}

//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)
	return u
}

//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)
	return &u
}
