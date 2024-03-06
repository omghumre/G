package main

import "fmt"

func main() {
	subjectMarks := map[string]float64{"Golang": 85, "java": 80, "python": 75}
	fmt.Println(subjectMarks)

	subjectMarks["Golang"] = 80
	fmt.Println(subjectMarks)

	//roll no , paper1,paper2,paper3 marks

	/////// adding element in map

	subjectMarks["cpp"] = 90
	fmt.Println(subjectMarks)

	/////// delete element from map

	delete(subjectMarks,"cpp")
	fmt.Println(subjectMarks)

	////// for loop 

	for sub,_ := range subjectMarks {
		fmt.Println(sub)
	}

}
