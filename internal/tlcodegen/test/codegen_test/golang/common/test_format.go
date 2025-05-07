package common

type MappingSuccessBytes struct {
	// expected bytes input and output
	Bytes string
	// expected TL2 bytes input and output
	BytesTL2 string
	// is tl1 data boxed
	IsTLBytesBoxed bool
}

type MappingTestSamplesBytes struct {
	// TL name for type to test
	TestingType string
	// json values which must success
	Successes []MappingSuccessBytes
}

type AllTestsBytes struct {
	Tests map[string]MappingTestSamplesBytes
}
