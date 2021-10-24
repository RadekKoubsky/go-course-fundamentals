package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Password \t\t", s)
	fmt.Println("Bcrypted password \t", string(bs))

	loginPassword := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("Invalid password")
		return
	}
	fmt.Println("Successfully logged in")

}
