package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"flag"
)

var (
	passwordAddrOpt = flag.String("p", "123", "password for VerneMQ broker")
)

func main() {
	flag.Parse()

	fmt.Printf("Password: %s\n", *passwordAddrOpt)
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(*passwordAddrOpt), 12)
	fmt.Printf("Encrypted: %s\n", encryptedPassword)
}