package fluent

func NewBuilder[T any]() *T {
	return new(T)
}

func (sample *SampleStruct) WithProperty1(value string) *SampleStruct {
	sample.property1 = value
	return sample
}

func (sample *SampleStruct) WithProperty2(value int) *SampleStruct {
	sample.property2 = value
	return sample
}

func (sample *SampleStruct) WithProperty3(value bool) *SampleStruct {
	sample.property3 = value
	return sample
}
