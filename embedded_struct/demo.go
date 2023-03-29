package main

import "fmt"

type Person struct {
	Name    string
	Sername string
}

func (p *Person) FullName() string {
	return p.Name + " " + p.Name
}

type Employee struct {
	Person
	Id     string
	Office string
}

func (e *Employee) Detail() string {
	return "ID : " + e.Id + "Office" + e.Office + ". Fullname" + e.FullName()
}
func (e *Employee) IsSameOffice(other *Employee) bool {
	return e.Office == other.Office
}

type Programmer struct {
	Employee
	Language []string
}

func (p *Programmer) Detail() string {
	return "Programmer : " + p.Employee.Detail()
}

type Tester struct {
	Employee
	Tools []string
}

func (t *Tester) Detail() string {
	return "Tester : " + t.Employee.Detail()
}
func main() {
	david := Person{
		Name:    "David",
		Sername: "Wright",
	}
	empDavid := Employee{
		Person: david,
		Id:     "123",
		Office: "Thailand",
	}
	progDavid := Programmer{
		Employee: empDavid,
		Language: []string{"golang", "java"},
	}
	fmt.Printf("%+v\n", progDavid)

	eavid := Person{
		Name:    "Eavid",
		Sername: "Wright",
	}
	empEavid := Employee{
		Person: eavid,
		Id:     "345",
		Office: "Thailand",
	}
	testerEavid := Programmer{
		Employee: empEavid,
		Language: []string{"Robot"},
	}
	fmt.Printf("%+v\n", testerEavid)
	fmt.Println(progDavid.IsSameOffice(&empEavid))
	fmt.Println(progDavid.IsSameOffice(&(testerEavid.Employee)))
	fmt.Println(progDavid.FullName())
	fmt.Println(progDavid.Detail())

	isSameOffice := (*Employee).IsSameOffice
	fmt.Println(isSameOffice(&empDavid, &empEavid))
}
