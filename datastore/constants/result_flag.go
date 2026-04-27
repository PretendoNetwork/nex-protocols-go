package constants

import (
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ResultFlag tells the server what fields to populate in responses
// to object searches
type ResultFlag uint8

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

// HasFlag checks if a given flag is set
func (rf ResultFlag) HasFlag(flag ResultFlag) bool {
	return rf&flag == flag
}

// HasFlag checks if all given flags are set
func (rf ResultFlag) HasFlags(flags ...ResultFlag) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if rf&flag != flag {
			return false
		}
	}

	return true
}

// String returns a human-readable representation of the ResultFlag bitmask.
// Multiple flags are joined with "|", e.g. "Tags|Ratings|MetaBinary".
// Returns "None" if no flags are set.
func (rf ResultFlag) String() string {
	if rf == 0 {
		return "None"
	}

	flags := []struct {
		flag ResultFlag
		name string
	}{
		{ResultFlagTags, "Tags"},
		{ResultFlagRatings, "Ratings"},
		{ResultFlagMetaBinary, "MetaBinary"},
		{ResultFlagPermittedIDs, "PermittedIDs"},
	}

	var parts []string
	for _, f := range flags {
		if rf&f.flag != 0 {
			parts = append(parts, f.name)
		}
	}

	return strings.Join(parts, "|")
}

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
