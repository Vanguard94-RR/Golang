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

type Person struct {
	Name     string
	LastName string
	Age      int
	Email    string
	Phone    string
	Score    float64
	Active   bool
}

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

func main() {
	// To create an instance of a struct, you can use the struct literal syntax.
	// For example, to create a new Person struct, you would write:
	Tittle("Person 1")
	p1 := Person{
		Name:     "John",
		LastName: "Doe",
		Age:      30,
		Email:    "john.doe@example.com",
		Phone:    "123-456-7890",
		Score:    95.5,
		Active:   true,
	}
	fmt.Println(p1)

	// You can also create a struct instance without initializing the fields, and then set the fields later.
	Tittle("Person 2 - Empty struct")
	p2 := Person{}
	p2.Name = "Jane"
	p2.LastName = "Smith"
	p2.Age = 25
	p2.Email = "jane.smith@example.com"
	p2.Phone = "987-654-3210"
	p2.Score = 88.7
	p2.Active = false
	if p2.Active {
		fmt.Println(p2)
	} else {
		fmt.Println("Person is not active")
	}

	// You can also create a pointer to a struct and access its fields using the pointer.
	Tittle("Person 3 - Using Pointer")
	p3 := &Person{
		Name:     "Alice",
		LastName: "Johnson",
		Age:      28,
		Email:    "alice.johnson@example.com",
		Phone:    "555-555-5555",
		Score:    75.2,
		Active:   true,
	}
	fmt.Println(p3)
	fmt.Println(p3.Name)
	fmt.Println(p3.LastName)
	fmt.Println(p3.Age)
	fmt.Println(p3.Email)
	fmt.Println(p3.Phone)
	fmt.Println(p3.Score)
	fmt.Println(p3.Active)

	// Another struct usage example
	Tittle("Person 4 - Using Pointer and modifying values")
	p4 := &Person{
		Name:     "Bob",
		LastName: "Smith",
		Age:      35,
		Email:    "bob.smith@example.com",
		Phone:    "111-222-3333",
		Score:    90.0,
		Active:   true,
	}
	fmt.Println(p4)
	p4.Age = 36 // Modifying the age field of p4
	fmt.Println("Updated Age:", p4.Age)

	// You can also create a struct with anonymous fields, which are fields that do not have a name.
	// Anonymous fields are accessed using the type of the field as the name.
	Tittle("Person 5 - Anonymous Fields")
	type Employee struct {
		Person   // Anonymous field of type Person
		Position string
	}
	e1 := Employee{
		Person: Person{
			Name:     "Charlie",
			LastName: "Brown",
			Age:      40,
			Email:    "charlie.brown@example.com",
			Phone:    "444-444-4444",
			Score:    85.0,
			Active:   true,
		},
		Position: "Manager",
	}
	fmt.Println(e1)
	fmt.Println(e1.Name)     // Accessing the Name field from the embedded Person struct
	fmt.Println(e1.Position) // Accessing the Position field from the Employee struct

	// You can also create a struct with nested structs, which are structs that contain other structs as fields.
	Tittle("Person 6 - Nested Structs")
	type Company struct {
		Name      string
		Employees []Employee
	}
	c1 := Company{
		Name: "Tech Company",
		Employees: []Employee{
			e1,
			{
				Person: Person{
					Name:     "Dave",
					LastName: "Smith",
					Age:      30,
					Email:    "dave.smith@example.com",
					Phone:    "777-777-7777",
					Score:    90.0,
					Active:   true,
				},
				Position: "Engineer",
			},
		},
	}

	fmt.Println(c1)
	fmt.Println(c1.Name)
	fmt.Println(c1.Employees)
	Tittle("Employees")
	Tittle("Employee " + c1.Employees[0].Name + " Details")
	fmt.Println(c1.Employees[0])
	fmt.Println(c1.Employees[0].Name)
	fmt.Println(c1.Employees[0].LastName)
	fmt.Println(c1.Employees[0].Age)
	fmt.Println(c1.Employees[0].Email)
	fmt.Println(c1.Employees[0].Phone)
	fmt.Println(c1.Employees[0].Score)
	fmt.Println(c1.Employees[0].Active)
	fmt.Println(c1.Employees[0].Position)
	Tittle("Employee " + c1.Employees[1].Name + " Details")
	fmt.Println(c1.Employees[1])
	fmt.Println(c1.Employees[1].Name)
	fmt.Println(c1.Employees[1].LastName)
	fmt.Println(c1.Employees[1].Age)
	fmt.Println(c1.Employees[1].Email)
	fmt.Println(c1.Employees[1].Phone)
	fmt.Println(c1.Employees[1].Score)
	fmt.Println(c1.Employees[1].Active)
	fmt.Println(c1.Employees[1].Position)

	// You can also create a struct with tags, which are metadata that can be associated with struct fields.
	// Tags are used to provide additional information about the fields, such as how they should be serialized or validated.
	Tittle("Person 7 - Struct Tags")
	type User struct {
		Name     string `json:"name" validate:"required"`           // Struct tags for JSON serialization and validation
		Email    string `json:"email" validate:"required,email"`    // Struct tags for JSON serialization and validation
		Password string `json:"password" validate:"required,min=8"` // Struct tags for JSON serialization and validation
	}
	u1 := User{
		Name:     "Eve",
		Email:    "eve.smith@example.com",
		Password: "password123",
	}
	fmt.Println(u1)
	fmt.Println(u1.Name)
	fmt.Println(u1.Email)
	fmt.Println(u1.Password)
	//Tags are purely metadata—Go ignores them unless a library reads them via reflection.

	// Structs
}
