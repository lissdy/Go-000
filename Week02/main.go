package main

import (
	svc "week02/service"
	"fmt"
)

func main(){
	_, err := svc.FindUserById("1")
	if err!=nil{
		fmt.Errorf("Error Happened %w", err)
		// Printing the error message
		fmt.Println(err.Error())
	}
}
