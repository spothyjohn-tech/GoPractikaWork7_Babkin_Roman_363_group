package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func (user *User) SetPassword(password string) {
	hash := sha256.Sum256([]byte(password))
	user.Password = hex.EncodeToString(hash[:])
}

func (user *User) VerifyPassword(password string) bool {
	hash := sha256.Sum256([]byte(password))
	if hex.EncodeToString(hash[:]) == user.Password {
		return true
	} else {
		return false
	}
}

func main() {
	TestUser := User{
		ID:       1,
		Username: "TestUser1",
		Email:    "TestUser@gmail.com",
	}

	TestUser.SetPassword("Password123")

	if TestUser.VerifyPassword("Password123") {
		fmt.Println("Вы зашли в систему как: " + TestUser.Username)
	} else {
		fmt.Println("Вы ввели неверный пароль")
	}
	if TestUser.VerifyPassword("123") {
		fmt.Println("Вы зашли в систему как: " + TestUser.Username)
	} else {
		fmt.Println("Вы ввели неверный пароль")
	}
}
