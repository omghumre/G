// Adding two no
package main

import (
	"fmt"
	"math"
)

// const (
//
//	A int = 1
//	B     = 3.14
//	C     = "Hi"
//
// )
var a int = 10

func main() {
	// const LENGTH int = 255
	// const WIDTH int = 5

	// var area int
	// area = LENGTH * WIDTH
	// fmt.Printf("area of rectange: %d", area)

	// var a uint8 = 255
	// a = a + 5
	// fmt.Println(a)

	// var x string = "jone"
	// fmt.Printf("%s\n", x)
	// fmt.Printf("%T", x)

	// var q bool
	// fmt.Println(q)

	// var a, b, c = 3, 4.6, "gg"
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T", c)
	// package main; import "fmt"; func main(){fmt.Println("hello");}
	// a, b, c := 1, 1.2, 4
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T", c)

	// var(
	// 	x int
	// 	y = 1
	// 	z string = "hello"
	// )'
	// fmt.Printf("%T\n", x)
	// fmt.Printf("%T\n", y)
	// fmt.Printf("%T", z)

	// fmt.Printf("%d\n%d\n%s",x,y,z)

	// fmt.Println(a)
	// var name,lname string;
	// fmt.Scanf("%s %s",&name,&lname)
	// fmt.Println("Hello",name,lname)
	// fmt.Printf("%s",name)

	// fmt.Scanln(&name)
	// fmt.Println(name)
	// var a,b int = 10,20
	// fmt.Printf("%d %d",a,b)
	// fmt.Printf("%d",a+b)
	// fmt.Print(a)
	// fmt.Print(b)

	// var name, lname, branch, clg = "om", "ghumre", "aiml", "rcoem"
	// fmt.Println("Name\t\t:\t", name)
	// fmt.Println("Last Name\t:\t", lname)
	// fmt.Println("Branch\t\t:\t", branch)
	// fmt.Println("College Name\t:\t", clg)
	// println()
	// fmt.Print("Name :", name)
	// fmt.Println("\tLast Name :", lname)
	// fmt.Print("Branch :", branch)
	// fmt.Println("\tCollege Name :", clg)
	// println()
	// fmt.Println("value of a is '", a, "'")
	// println()

	// var a int = 10
	// var z float64= 1.1
	// fmt.Printf("%v %v %v", a, z, name)

	// println(15 == 017) // 017 is octal number which is 15 in decimal
	// println(15 == 0o17)

	var val int = 10
	val++
	// fmt.Printf("%d\n",val)
	val--

	val = 15
	// fmt.Printf("\nBinary of %d is %0b",val,val)
	// println(val)
	val = 3e5 // it means 3 x 10^5
	// aEn = a x 10^n
	// println(val)

	// println(15 == 0xf)
	// fmt.Printf("exp is %t",15==0xf)

	// var i = 5
	// fmt.Printf("%s\n",i)
	// fmt.Printf("%s\n",j)
	// // fmt.Printf("%s %s",i,j)
	// fmt.Print(i,j)
	// 	println()
	// 	println()
	// var txt = "Hello"
	// fmt.Printf("%s",txt)
	// fmt.Printf("\n\n%+d",i) //%+d se sign print karega
	// fmt.Printf("\n\n%o",i)

	// var txt = "Hello"
	// integer := 23
	// // fmt.Printf("\n\n%-8s",txt) //space print after string
	// fmt.Printf("\n\n% x",txt)
	// fmt.Print("%T\n%T",integer,&integer)

	// number,floatNum := 238,1234.5678987654
	// fmt.Printf("%d\n",number)
	// fmt.Printf(".2f : %.2f\n",floatNum)

	// text := "rajesh"
	// fmt.Printf("%*s\n",40,text)
	// fmt.Printf("%010s",text)
	println()
	// println()
	// ch := 'a'
	// st := "a"
	// fmt.Printf("%T", ch) // gives int32
	// fmt.Printf("%T", st)
	// 	println()

	// var v byte = 'z'
	// fmt.Printf("%d", v)

	// var v float64 = -19.25
	// fmt.Printf("absolute of %d is %f",v,math.Abs(v))
	
	// n1 := 3.0
	// n2 := 8.0
	// large := math.Max(n1,n2)
	// fmt.Printf("Largest : %f",large)         Abs() and Max() accept float values only

	var n1 float64 = 10.258748574

	fmt.Printf("%f",math.Floor(n1))
	fmt.Printf("%f\n",n1)
	fmt.Printf("%g\n",n1)               // %g print the exact decimal
}
