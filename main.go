package main

import "fmt"


func main() {
	 matrix := [][][]int{
		{
			{1,2,3,4},
			{1,2,3,4},},
		{},
		{},
		{
			{1,2,10,4},
			{1,2,3,4},
		} ,
	}
	fmt.Println(matrix[3][0][2])
	fmt.Println(matrix[0][1][2])
 
}



