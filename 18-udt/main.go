package main

import "fmt"

// main is the entry point of the program.
// It demonstrates the usage of various struct types and their initialization.
func main() {

	var p1 Person
	p1.Name = "Jiten"
	p1.Id = 100
	p1.Email = "Jitenp@Outlook.com"
	p1.Address = "Bangalore"
	p1.Mobile = "9618558500"

	fmt.Println(p1)

	p2 := Person{Id: 101, Name: "Rahul", Email: "Raghul@abcd.com", Mobile: "123123", Address: "Hyd"}

	fmt.Println(p2)

	p3 := new(Person)
	p3.Name = "Jiten"
	p3.Id = 100
	p3.Email = "Jitenp@Outlook.com"
	p3.Address = "Bangalore"
	p3.Mobile = "9618558500"
	fmt.Println(p3)

	p4 := &Person{Id: 101, Name: "Rahul"}
	fmt.Println(p4)

	e1 := Employee{Id: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "123123", Address: Address{City: "Bangalore", Country: "India", PinCode: "560086"}}

	a2 := Address{City: "Bangalore", Country: "India", PinCode: "560086"}
	e2 := Employee{Id: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "123123", Address: a2}

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e1.Address.City)

	s1 := Student{Id: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "123123", Address: Address{Name: "Rajaji Nagar", City: "Bangalore", Country: "India", PinCode: "560086"}}

	// if city is in Address which is a promoted field

	fmt.Println(s1.City)
	fmt.Println(s1.Address.Name)

	cc1 := ColourCode{int: 100, string: "Red", integer: 12312}
	fmt.Println(cc1.string)
	fmt.Println(cc1)

	cm1 := Company{Name: "Rakuten", Address: Address{City: "Bangalore", Country: "India"}, MoreInfo: struct {
		Website string
		GSTNo   string
	}{Website: "www.Rakuten.com", GSTNo: "12312312312"}}

	fmt.Println(cm1)

	struct1 := struct {
		Name string
		Age  uint8
	}{
		Name: "Jiten",
		Age:  38,
	}

	fmt.Println(struct1)

	type struct2 = struct {
		Name string
		Age  uint8
	}

	s2 := struct2{Name: "Jiten", Age: 38}
	fmt.Println(s2)
}

type integer = int // create alias

type Person struct {
	Id      int
	Name    string
	Email   string
	Mobile  string
	Address string
	Married bool
	Age     uint8
}

type Address struct {
	Name, City, Country, PinCode string
}

type Employee struct {
	Id          int
	Name        string
	Email       string
	Mobile      string
	Address     Address // embedding , composition
	SocialMedia []string
}

// Promoted fields

type Student struct {
	Id      int
	Name    string
	Email   string
	Mobile  string
	Address // promoted field
}

// no special field names but just fields
type ColourCode struct {
	int
	int2
	integer
	string
}

type int2 = int

type Company struct {
	Name string
	Address
	MoreInfo struct { // inline struct
		Website string
		GSTNo   string
	}
}
