package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// ResultFlag tells the server what fields to populate in responses
// to object searches
type ResultFlag uint8

const (
	// ResultFlagTags means the object tags should be populated
	ResultFlagTags ResultFlag = 0x1

	// ResultFlagRatings means the object ratings should be populated
	ResultFlagRatings ResultFlag = 0x2

	// ResultFlagMetaBinary means the object MetaBinary should be populated
	ResultFlagMetaBinary ResultFlag = 0x4

	// ResultFlagPermittedIDs means the object permissions should be populated
	ResultFlagPermittedIDs ResultFlag = 0x8
)

// WriteTo writes the ResultFlag to the given writable
func (rf ResultFlag) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(rf))
}

// ExtractFrom extracts the ResultFlag value from the given readable
func (rf *ResultFlag) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*rf = ResultFlag(value)
	return nil
}
