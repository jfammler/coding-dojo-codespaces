package main

import "fmt"



type skill struct {
	name string
	level int
}

type person struct {
	ID string
    name string
    skills  [100]skill
}

type project struct {
	ID string
    name string
    skills  [100]skill
}

func newPerson(ID string, name string, skills []skill) *person {
	p := person{ID: ID, name: name}
	copy(p.skills[:], skills)
	return &p
}


func main() {
	fmt.Printf("test")

	persons := make([]*person, 0)
	persons = append(persons, newPerson("1", "A", []skill{{"Java", 5}, {"Python", 3}, {"C++", 2}}))
	persons = append(persons, newPerson("2", "B", []skill{{"Java", 5}, {"Python", 3}, }))
	persons = append(persons, newPerson("3", "C", []skill{{"Java", 5}, {"Python", 3}}))
	persons = append(persons, newPerson("4", "D", []skill{{"Java", 5}, {"Python", 3}}))


	fmt.Printf("First person: %v\n", persons[0].name)
	
}






