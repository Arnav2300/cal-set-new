package services

import (
	"fmt"
	"user-management-service/api/repository"
)

func SingupService(user repository.User) {
	//get email,username, password and role
	//check if email is already present in db, if yes then fail
	//hash the password
	//write into db
	//return status created message
	fmt.Print(user)
}
