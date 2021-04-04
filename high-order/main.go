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

	fmt.Printf("Accepted Students: %v \n", f)
}
