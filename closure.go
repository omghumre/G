package main 
import ("fmt")

func  calculate() func() int {
	// println("above num")
	num := 1
	// println("below num")
	return func() int {
		num += 2
		return num
	}
}

func  greet() func() string {
	name := "john"
	return func() string {
		name += " Hi"
		return name
	}
}

func displayNumber() func() int {
	num := 0
return func() int {
	num++
	return num 
}

}
func main() {
	// var odd3 = greet()
	// fmt.Println(odd3())
	// odd1 := calculate()
	// fmt.Println(odd1())
	// fmt.Println(calculate()())
	// fmt.Println(calculate()())
	// fmt.Println(calculate()())
	// fmt.Println(odd1())
	// odd1 = calculate()
	// fmt.Println(odd1())
	// println()
	// odd2 := calculate()
	// fmt.Println(odd2())

	
var fib func(n int) int
fib = func(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1)+fib(n-2)
} 
fmt.Println(fib(7))

}