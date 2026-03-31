// In Golang, a struct (short for "structure") is a composite data type that
// groups together variables under a single name.
// Each variable within a struct is called a field. Structs are used to
// create complex data types that can represent real-world entities or concepts.

// Structs are the closest thing to classes in Go, but they do not have methods or
// inheritance like classes in other languages.
// However, you can define methods on structs to provide behavior and functionality.

// Structs are a collection of fields, and each field has a name and a type.
// Structs are defined using the "type" keyword followed by the name of the
// struct and the fields it contains.
// For example, to define a struct for a "Person" with fields for name and age,
// you would write:
package main

import "fmt"

func Tittle(text string) {
	leftDashes := (80 - len(text)) / 2
	rightDashes := 80 - len(text) - leftDashes

	left := ""
	for i := 0; i < leftDashes; i++ {
		left += "-"
	}
	right := ""
	for i := 0; i < rightDashes; i++ {
		right += "-"
	}
	fmt.Println(left + text + right)
}

// Structs can also be nested, meaning that a struct can contain another struct as a field.
type chain struct {
	base
	value int
	next  *chain
}

type base struct {
}

type Person struct {
	Name     string
	LastName string
	Age      int
	Email    string
}

func (x *chain) sayOk() {
	fmt.Println("OK from chain")
}
func (x *chain) sayHello() {
	fmt.Println("Hello from chain")
}
func (b *base) sayHello() {
	fmt.Println("Hello from base")
}
func (b *base) sayOk() {
	fmt.Println("OK from base")
}

func main() {
	Tittle("Structs Part III")
	b1 := base{}
	b1.sayHello()
	b1.sayOk()
	c1 := chain{}
	c1.sayHello()
	c1.sayOk()

	// Accessing fields in a struct
	Tittle("Accessing fields in a struct")

	c1 = chain{value: 10}
	fmt.Println("Value in chain:", c1.value)

	c1.next = &chain{value: 20}
	fmt.Println("Value in next chain:", c1.next.value)
	c1.next.next = &chain{value: 30}                                       // Creating a new chain struct and assigning it to the next field of the current chain
	fmt.Println("Value in next next chain:", c1.next.next.value)           // Accessing the value field of the next next chain struct through the current chain struct
	c1.next.next.next = &chain{value: 40}                                  // Creating a new chain struct and assigning it to the next field of the next next chain
	fmt.Println("Value in next next next chain:", c1.next.next.next.value) // Accessing the value field of the next next next chain struct through the current chain struct

	// Structs can also contain slices, which are dynamic arrays in Go.
	type Company struct {
		Name      string
		Employees []Person
	}

	company := Company{
		Name: "Tech Co",
		Employees: []Person{
			{Name: "Alice", LastName: "Smith", Age: 30, Email: "alice.smith@techco.com"},
			{Name: "Bob", LastName: "Johnson", Age: 35, Email: "bob.johnson@techco.com"},
		},
	}
	Tittle("Structs with slices")

	fmt.Println("Company:", company)
	fmt.Println("Employees:", company.Employees)
	fmt.Println("Employee:", company.Employees[0])
	fmt.Println("Employee Name:", company.Employees[0].Name)
	fmt.Println("Employee Last Name:", company.Employees[0].LastName)
	fmt.Println("Employee Age:", company.Employees[0].Age)
	fmt.Println("Employee Email:", company.Employees[0].Email)

	//Structs with maps
	type Department struct {
		Name      string
		Employees map[string]Person
	}

	department := Department{
		Name: "Engineering",
		Employees: map[string]Person{
			"Alice": {Name: "Alice", LastName: "Smith", Age: 30, Email: "alice.smith@techco.com"},
			"Bob":   {Name: "Bob", LastName: "Johnson", Age: 35, Email: "bob.johnson@techco.com"},
		},
	}
	Tittle("Structs with maps")

	fmt.Println("Department:", department)
	fmt.Println("Employees:", department.Employees)
	fmt.Println("Employee:", department.Employees["Alice"])
	fmt.Println("Employee Name:", department.Employees["Alice"].Name)
	fmt.Println("Employee Last Name:", department.Employees["Alice"].LastName)
	fmt.Println("Employee Age:", department.Employees["Alice"].Age)
	fmt.Println("Employee Email:", department.Employees["Alice"].Email)
}
