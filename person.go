package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Saludo() string {
	return "Hola, mi nombre es " + p.Name + " y tengo " + fmt.Sprint(p.Age) + " años."
}

func (p *Person) Cumpleaños() {
	p.Age++
}

func main() {
	fmt.Println("Hello, World!")
	p1 := &Person{Name: "Uayeb", Age: 35}
	message := p1.Saludo()

	fmt.Println(message)
	p1.Cumpleaños()
	message = p1.Saludo()

	fmt.Println(message)

}
