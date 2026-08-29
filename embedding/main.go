package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	writePromotionDemo(os.Stdout)
}

func writePromotionDemo(output io.Writer) {
	person := Person{
		Name: "Amir",
		Address: Address{
			City: "Auckland",
		},
	}

	fmt.Fprintf(output, "promoted field: %s\n", person.City)
	fmt.Fprintf(output, "explicit field: %s\n", person.Address.City)
	fmt.Fprintf(output, "promoted method: %s\n", person.Location())
	fmt.Fprintf(output, "explicit method: %s\n", person.Address.Location())
}
