package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (ele Person) intro(arg string) string { //Personのメソッド
	return arg + " I am" + " " + ele.name
}

//ex4:継承に似た機能として構造体の埋め込みが可能
//構造体UserにPersonを組み込む。
type User struct {
	Person
}

func (ele User) intro(arg string) string { //Userのメソッド
	return "User No." + arg + " " + ele.name
}

func main() {
	bob := Person{"Bob", 85}

	var user1 User
	//組み込みによりnameの定義が可能
	user1.name = "Bob"

	fmt.Println(bob.intro("Hello!")) //=> Hello! I am Bob
	fmt.Println(user1.intro("1"))    //=> User No.1 Bob
}
