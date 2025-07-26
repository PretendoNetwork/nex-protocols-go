package constants

// * Miscellaneous constants

const (
	// MaxStringLength is the max length a UserMessage/TextMessage string may be
	//
	// NOTE: We don't know the real name of this constant, if there is one. This
	// name is based on the DataStore constant `MaxNameLength`
	MaxStringLength uint32 = 256

	// MaxBinarySize is the max size a BinaryMessage body may be
	//
	// NOTE: We don't know the real name of this constant, if there is one. This
	// name is based on the DataStore constant `MaxMetaBinSize`
	MaxBinarySize uint32 = 512
)
