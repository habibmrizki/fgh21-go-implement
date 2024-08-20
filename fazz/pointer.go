package fazz

import "fmt"

// type person struct {
// 	name string
// 	lastname string
// }

// func getInfo(person person)  {
// 	person.lastname = "rizki"
// }

type Data struct {
	name string
}

func main() {
	// var person = person {
	// 	name: "Habib",
	// 	lastname: "hghg",
	// }
	// getInfo(person)
	// fmt.Println(person)


	// arr := []string{"a", "b", "c"}
	// lisData := &arr[0]
	// *lisData = "test"
	// fmt.Println(arr)
	// fmt.Println(*lisData)

	data := Data{
		name: "Fazz",
		
	}
	newData := &data.name
	*newData = "Track"
	fmt.Println(data)
	fmt.Println(*newData)

}



