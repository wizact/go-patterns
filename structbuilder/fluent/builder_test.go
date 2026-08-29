package fluent

import "testing"

func TestNewBuilder_ChainedSetters_ConfiguresSampleStruct(t *testing.T) {
	want := SampleStruct{
		property1: "test_value",
		property2: 42,
		property3: true,
	}

	got := NewBuilder[SampleStruct]().
		WithProperty1("test_value").
		WithProperty2(42).
		WithProperty3(true)

	if *got != want {
		t.Errorf("expected %+v, got %+v", want, *got)
	}
}
