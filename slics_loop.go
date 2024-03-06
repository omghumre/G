package main

import "fmt"
import "sort"

func main() {
	// numbers := []int{1, 9, 8, 4, 5, 6}

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	// making slice from array

	// intSlice := numbers[2:5]
	// fmt.Println(intSlice)
	
		// sort.Ints(numbers) //            Ints is int slice, used to sort slice
		// fmt.Println(numbers)

slices:= []string {"i","the","best","policy"}	
sort.Strings(slices)
for _,ele := range slices{
	fmt.Printf("%s",ele)
}

status := sort.StringsAreSorted(slices)
println(status)
}
