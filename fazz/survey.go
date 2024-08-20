package fazz

import "fmt"


	type survey struct {
		name string
		age int
		gender string
		isSmoker bool
		cigarVariant  []string
	}

	func Survey() {
	surveys := []survey{
		survey{
			name: "habib",
			age: 20,
			gender: "Laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Malboro",
				"Esse",
				"Berry Pop",
				"Marcopolo",
			},
		},
       	survey{
			name: "habib",
			age: 20,
			gender: "Laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Malboro",
				"Esse",
				"Berry Pop",
				"Marcopolo",
			},
		},
		survey{
			name: "habib",
			age: 20,
			gender: "Laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Malboro",
				"Esse",
				"Berry Pop",
				"Marcopolo",
			},
		},
		survey{
			name: "habib",
			age: 20,
			gender: "Laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Malboro",
				"Esse",
				"Berry Pop",
				"Marcopolo",
			}, 
		},
	}

   
	for i := 0; i < len(surveys); i++ {
		fmt.Println("Name:", surveys[0].name)
		fmt.Println("Age:", surveys[0].age)
		fmt.Println("Gender:", surveys[0].gender)
		fmt.Println("is respondent:", surveys[0].isSmoker)
		fmt.Println("cigar variant:", surveys[0].cigarVariant)
	}
}