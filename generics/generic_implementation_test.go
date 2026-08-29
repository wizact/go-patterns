package generics

import (
	"testing"
)

func TestGenericImplementation_ConcreteInstantiation(t *testing.T) {
	expected := "Hello Amir"
	p := P[string]{}
	actual := p.M("Amir")

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGenericImplementation_InterfaceImplementation(t *testing.T) {
	expected := "Hello Amir"

	var p P[string]
	var i I = p

	actual := i.M("Amir")
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
