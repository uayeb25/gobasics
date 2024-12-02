package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	capitales := map[string]string{
		"Mexico":      "CDMX",
		"Honduras":    "Tegucigalpa",
		"Guatemala":   "Guatemala",
		"El Salvador": "San Salvador",
		"Nicaragua":   "Managua",
	}

	pais := "Costa Rica"

	if capital, ok := capitales[pais]; ok {
		fmt.Println("La capital de", pais, "es", capital)
	} else {
		fmt.Println("No se encontro la capital de", pais)
	}
}
