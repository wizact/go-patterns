package functionaloptions

type Option func(*SampleStruct)

func NewSampleStruct(options ...Option) *SampleStruct {
	sample := &SampleStruct{
		property1: "default",
		property2: 1,
		property3: true,
	}

	for _, option := range options {
		option(sample)
	}

	return sample
}

func WithProperty1(value string) Option {
	return func(sample *SampleStruct) {
		sample.property1 = value
	}
}

func WithProperty2(value int) Option {
	return func(sample *SampleStruct) {
		sample.property2 = value
	}
}

func WithProperty3(value bool) Option {
	return func(sample *SampleStruct) {
		sample.property3 = value
	}
}
