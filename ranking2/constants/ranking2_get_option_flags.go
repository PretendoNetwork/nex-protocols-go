package constants

// Ranking2GetOptionFlags determines what data is returned when requesting common data.
type Ranking2GetOptionFlags uint32

const (
	// Ranking2GetOptionFlagsNothing means that no extra data should be returned
	// in the common data.
	//
	// Note: The `userName` and `binaryData` fields seem to always be populated, regardless
	// of flags. These seem to exist solely to enable the `mii` field.
	Ranking2GetOptionFlagsNothing Ranking2GetOptionFlags = 0

	// Ranking2GetOptionFlagsMii means that Mii data should be returned
	// in the common data.
	//
	// Note: This is a guess based on some light test behavior. The real
	// name of this is unknown.
	Ranking2GetOptionFlagsMii Ranking2GetOptionFlags = 1
)
