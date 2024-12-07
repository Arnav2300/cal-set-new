package services

import (
	"fmt"
	"user-management-service/api/dto"
)

func SingupService(user dto.SignupDTO) {
	//check if email is already present in db, if yes then fail

	//hash the password
	//write into db
	//return status created message
	fmt.Print(user)
}
