package main

import (
	"bytes"
	"testing"
)

func TestPerson_EmbeddedAddress_PromotesCity(t *testing.T) {
	person := Person{Address: Address{City: "Auckland"}}

	if person.City != "Auckland" {
		t.Errorf("expected Auckland, got %s", person.City)
	}
}

func TestPerson_EmbeddedAddress_PromotesLocation(t *testing.T) {
	person := Person{Address: Address{City: "Auckland"}}

	if person.Location() != "Auckland" {
		t.Errorf("expected Auckland, got %s", person.Location())
	}
}

func TestWritePromotionDemo_PrintsEquivalentSelectors(t *testing.T) {
	var output bytes.Buffer

	writePromotionDemo(&output)

	want := "promoted field: Auckland\n" +
		"explicit field: Auckland\n" +
		"promoted method: Auckland\n" +
		"explicit method: Auckland\n"
	if output.String() != want {
		t.Errorf("expected %q, got %q", want, output.String())
	}
}
