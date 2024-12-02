package main

import "errors"

func checkNumber(num int) error {
	if num < 0 {
		return errors.New("El número no puede ser negativo")
	}
	return nil
}

func main() {

	if err := checkNumber(8); err != nil {
		panic(err)
	} else {
		println("El número es válido")
	}

}
