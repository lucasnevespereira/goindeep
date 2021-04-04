package main

import "fmt"

type student struct {
	 name string
	 grade int
	 attendance int
}

func filter(students []student, f func(student) bool) []student {
	var accepted []student
	for _, s := range students {

		if f(s) {
			accepted = append(accepted, s)
		}
	}

	return accepted
}

func genFilterFunc(schoolName string) func(student) bool {
	if schoolName == "B" {
		return func(s student) bool {
			if s.grade >= 5 {
				return true
			}

			return false
		}
	}
	return func(s student) bool { return true }
}


func main() {
	s1 := student {
		name: "Bob",
		grade: 4,
		attendance: 8,
	}

	s2 := student {
		name: "Alice",
		grade: 8,
		attendance: 4,
	}

	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade >= 5 {
			return true
		}
		return false
	})

	fmt.Printf("Accepted Students: %v \n", f) // [{Alice 8 4}]


	// Create a function generator
	// Usecase => Example, for school A use func X instead of func Y.
	filterFuncForSchoolA := genFilterFunc("A")
	f = filter(s, filterFuncForSchoolA)
	fmt.Printf("Accepted Students in school A: %v \n", f) //  [{Bob 4 8} {Alice 8 4}]

	// School B
	filterFuncForSchoolB := genFilterFunc("B")
	f = filter(s, filterFuncForSchoolB)
	fmt.Printf("Accepted Students in school B: %v \n", f) //  [{Alice 8 4}]

}
