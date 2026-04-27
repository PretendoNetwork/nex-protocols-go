package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MiiCharacterSet defines the character set (font/language) used by a Mii
type MiiCharacterSet uint8

// WriteTo writes the MiiCharacterSet to the given writable
func (mcs MiiCharacterSet) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(mcs))
}

// ExtractFrom extracts the MiiCharacterSet value from the given readable
func (mcs *MiiCharacterSet) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*mcs = MiiCharacterSet(value)
	return nil
}

// String returns a human-readable representation of the MiiCharacterSet.
func (mcs MiiCharacterSet) String() string {
	switch mcs {
	case MiiCharacterSetJUE:
		return "JUE"
	case MiiCharacterSetCHN:
		return "CHN"
	case MiiCharacterSetKOR:
		return "KOR"
	case MiiCharacterSetTWN:
		return "TWN"
	default:
		return fmt.Sprintf("MiiCharacterSet(%d)", int(mcs))
	}
}

const (
	// MiiCharacterSetJPN means that the Mii uses the JPN + USA + EUR character set
	MiiCharacterSetJUE MiiCharacterSet = iota

	// MiiCharacterSetCHN means that the Mii uses the Chinese character set
	MiiCharacterSetCHN

	// MiiCharacterSetKOR means that the Mii uses the Korean character set
	MiiCharacterSetKOR

	// MiiCharacterSetTWN means that the Mii uses the Taiwanese character set
	MiiCharacterSetTWN
)
