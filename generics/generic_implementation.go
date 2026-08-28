package generics

type I interface {
	M(string) string
}

type P[T string] struct {
}

func (p P[T]) M(name T) T {
	return "Hello " + name
}
