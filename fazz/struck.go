package fazz

import "fmt"


type Bio struct {
	fullname string
	skillSet []SkillSet
}

type SkillSet struct {
	name string
}
func main() {
	
 bio := []Bio {
	{
		fullname: "habib",
			skillSet: []SkillSet{
			{
				name: "JavaScript",
			},
			{
				name: "CSS",
			},
			{
				name: "PHP",
			},
			{
				name: "HTML",
			},
		},
	},
	{
		fullname: "Joki",
			skillSet: []SkillSet{
			{
				name: "JavaScript",
			},
			{
				name: "CSS",
			},
			{
				name: "PHP",
			},
			{
				name: "HTML",
			},
		},
	},
	{
		fullname: "Budi",
			skillSet: []SkillSet{
			{
				name: "JavaScript",
			},
			{
				name: "CSS",
			},
			{
				name: "PHP",
			},
			{
				name: "HTML",
			},
		},
	},
	{
		fullname: "Bambang",
			skillSet: []SkillSet{
			{
				name: "JavaScript",
			},
			{
				name: "CSS",
			},
			{
				name: "PHP",
			},
			{
				name: "HTML",
			},
		},
	},
 }
 fmt.Println(bio[4].skillSet[2].name)
}
