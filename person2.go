package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

//ex3:構造体内にメソッドを定義できる
//普通の関数と違うのはレシーバ引数(下記の「(ele Person)」)の部分だけ。
func (ele Person) intro(arg string) string {
	return arg + " I am" + " " + ele.name + "."
}

func main() {
	bob := Person{"Bob", 85}
	fmt.Println(bob.intro("Hello!")) //=>Hello! I am Bob
}
