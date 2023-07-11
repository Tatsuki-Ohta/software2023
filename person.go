package main
import (
    "fmt"
)
type Person struct {
   name string 
   age int
}

func main(){
//ex2:複数の初期化方法が存在する
    //初期化方法①:変数定義後にフィールドを設定する
    var mike Person
    mike.name = "Mike"
    mike.age = 23

    //初期化方法②: {} で順番にフィールドの値を渡す
    var bob = Person{"Bob", 35}

    //初期化方法③:フィールド名を ： で指定する方法
    var sam = Person{age:89, name: "Sam"}

    fmt.Println(mike, bob, sam) //=> {Mike 23} {Bob 35} {Sam 89}
}
