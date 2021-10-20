package main

import "fmt"

type Person struct {
	name string
	age  int
	address string
}

// default functions

func (obj *Person) setDefault() {
	obj.name, obj.address, obj.age = "null", "null", -1
}

// getters

func (obj *Person) getName() string {
	return obj.name
} 

func (obj *Person) getAddress() string {
	return obj.address
} 

func (obj *Person) getAge() int {
	return obj.age
} 

// setters

func (obj *Person) setName(name string) {
	obj.name = name
} 

func (obj *Person) setAddress(address string) {
	obj.address = address
} 

func (obj *Person) setAge(age int) {
	obj.age = age
}

func main() {
	var v Person
	v.setDefault()

	v.setName("Ritik")
	v.setAddress("Pep")
	v.setAge(20)
	fmt.Println(v)

	g := Person{v.getName(), v.getAge(), v.getAddress()}
	fmt.Println(g)
}