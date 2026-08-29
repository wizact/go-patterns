package functionaloptions

import "testing"

func TestNewSampleStruct_NoOptions_ReturnsDefaults(t *testing.T) {
	want := SampleStruct{
		property1: "default",
		property2: 1,
		property3: true,
	}

	got := NewSampleStruct()

	if *got != want {
		t.Errorf("expected %+v, got %+v", want, *got)
	}
}

func TestNewSampleStruct_RepeatedOption_LastWins(t *testing.T) {
	got := NewSampleStruct(
		WithProperty1("first"),
		WithProperty1("second"),
	)

	if got.property1 != "second" {
		t.Errorf("expected second, got %s", got.property1)
	}
}

func TestNewSampleStruct_WithProperty2_OverridesDefault(t *testing.T) {
	got := NewSampleStruct(WithProperty2(42))

	if got.property2 != 42 {
		t.Errorf("expected 42, got %d", got.property2)
	}
}

func TestNewSampleStruct_WithProperty3_OverridesDefault(t *testing.T) {
	got := NewSampleStruct(WithProperty3(false))

	if got.property3 {
		t.Error("expected false, got true")
	}
}
