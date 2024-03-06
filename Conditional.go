package main 
import (
	"fmt"
	// "os"
)
func main(){
	// if 20 > 18 {
	// 	fmt.Printf("Hello")
	// }
	// if true { fmt.Printf("Hello") ; println("hello")}
	// if ! (20<18) {fmt.Printf("This is num")}

	// a := 10
	// b := 20
	// c := 30

	// if a,b,c := 10,19,20; a < b && b < c{
	// 	fmt.Println("a is less")
	// } 										// we can define variable before the if codition 
	// print(b)
	// var name string 
	// var age int
	// var v string 
	// var err string 
// 	fmt.Printf("Enter the value ") 
// 	if v,err := fmt.Scan(&name, &age); err != nil{
// 		// fmt.Scan return t20 value , first the input value and second the error if occured
// 		fmt.Println(err)
// 		fmt.Println(v)
// 		os.Exit(1)
// 	}
// 	print("correct")
// 	print(v)

// var a int /
// fmt.Scanln(&a)
// var b int 
// fmt.Scanln(&b)
// var c int 
// fmt.Scanln(&a,&b,&c)

// if a > b && a > c {fmt.Println(a)}
// if b > a && b > c {println(b)}
// if c > a && c > b {println(c)}
// }

// if a > b {
// 	print("if")
// } else {
// 	print("else")
// }

var num int;
fmt.Println("ENtre a number : ")
if _,err:= fmt.Scanln(&num);num < 0 {
	print("number less than 0",err)
}
if(num % 2 == 0) {
	println("It is even")
} else {
	println("It is odd")
}


}