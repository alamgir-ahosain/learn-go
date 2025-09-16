package main

import "fmt"

//interface Parent1 with a method to printDetails
type Parent1 interface {
	printDetails()
}

type Parent2 interface {
	cgpaIncrease(float64) float64
}

// struct child that will implement both interfaces
type child struct {
	Id   string
	Name string
	Cgpa float64
}

//receiver function / method, / behavior
func (obj child) printDetails() {
	fmt.Print(obj.Id, " ", obj.Name, " ", obj.Cgpa, "\n")

}
func (obj child) cgpaIncrease(cgpa float64) float64 {
	obj.Cgpa += cgpa
	return obj.Cgpa
}
func main() {

	// Declare a variable of type Parent1 and assign a child instance
	var u1 Parent1
	u1 = child{Id: "CE21012", Name: "Alamgir", Cgpa: 3.44}
	u1.printDetails() // Calls printDetails via interface Parent1

	var u2 Parent2
	u2 = child{Id: "CE21012", Name: "Alamgir Hosain", Cgpa: 3.44}
	fmt.Println(u2.cgpaIncrease(0.2))

}
