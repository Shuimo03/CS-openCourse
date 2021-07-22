package main

import (
	"fmt"
)

type User struct {

	Name string
	Age int
	Sex string
}

/**
	simple DB, save User information
	id ==> context, for example:
	1: [Name: "Dorothy",Age: 65,Sex: "woman"]
 */
var userDB = map[int]User{
	1:  {Name: "Dorothy",Age: 65,Sex: "woman"},
	2:  {Name:"Eloise",Age: 65,Sex: "man"},
	3: 	{Name:"Florence",Age: 65,Sex: "man"},
	4: 	{Name:"Pearl",Age: 65,Sex: "woman"},
}

// 查询用户

func queryUser(id int) (User,error){

	if userInfo,ok := userDB[id]; ok{
		return userInfo,nil
	}

	return User{},fmt.Errorf("UserDB dont have input %d Id",id)
}

func main() {
	var input int
	fmt.Scanln(&input)

	userInfo, err := queryUser(input)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("name: %s, age: %d  sex: %s \n", userInfo.Name, userInfo.Age,userInfo.Sex)
}