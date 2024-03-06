package main 
import "fmt"


func add(){
	n1 := 1
	n2 := 2
	fmt.Println("sum is",n1+n2)
}


func sub (n1 int, n2 int ) int {
	// fmt.Println(n1,"-",n2,"=",n1-n2)
	// return n1+n2
	fmt.Println("hello")
	return  n1-n2
}
func calculate(n1 int,n2 int) (int , int) {
	sum:= n1 + n2 
	diff := n1-n2
	return sum,diff 
} 

func square(n1 int) int {
	square := n1 * n1
	return square
}

const sum int = 9
func addnum(sum int ) int{
sum = sum +1
return sum 
}

func count(num int) int {
	if num <= 0 { return 0}
	count(num-1)
	println(num)
	return 0
}

func sumNum(num int) int {
	if num <= 0{ return 0 }
	return num + sumNum(num - 1)
}

func fact(num int) int {
	if num == 1 || num <= 0 { return 1}
	return num * fact(num - 1)
}

func fibo(num int, sum int) int {
	if num == 0 || num == 1 {
		sum += num
	} else {
		sum = fibo(num-1,sum) + fibo(num-2, sum)
	}
	return sum 
}

	

// var global = func(){
// 	fmt.Println("this is global")
// }

// func main(){
	// greet()
	// var n1 int
	// var n2 int 
	// fmt.Println("Enter two numbers")
	// fmt.Scanf("%d %d",&n1,&n2)
	// n1,n2 = calculate(n1,n2)
	// fmt.Println("Sum is",n1,"Sub is",n2)
	// fmt.Println(sum)
	// fmt.Println(addnum())
		
// }
	

// func squ(n1 int)int {
// 	ad(n1)
// return n1*n1
// }

// func ad(n1 int) int {
// 	squ(n1)
// 	return n1 + n1
// }


func main() {
												// Anonymous function
// var greet = func() {
// 	fmt.Println("I am greet")
// }
// wel := greet
// wel()
// global()
// greet()

// var sum = func(n1 , n2 int) int{
// 	result := n1 + n2
// 	return result
// }

// fmt.Printf("%T",sum)


}

