package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PersistentGatheringType indicates the type of PersistentGathering
type PersistentGatheringType uint32

// WriteTo writes the PersistentGatheringType to the given writable
func (pgt PersistentGatheringType) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(pgt))
}

// ExtractFrom extracts the PersistentGatheringType value from the given readable
func (pgt *PersistentGatheringType) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*pgt = PersistentGatheringType(value)
	if !pgt.IsValid() {
		return fmt.Errorf("Value %d is out of range", *pgt)
	}

	return nil
}

// IsValid ensures the value of the PersistentGatheringType is within
// the expected range
func (pgt PersistentGatheringType) IsValid() bool {
	return pgt >= PersistentGatheringTypeOpen && pgt <= PersistentGatheringTypeOfficial
}

const (
	// PersistentGatheringTypeOpen indicates that the PersistentGathering is open to everyone
	PersistentGatheringTypeOpen PersistentGatheringType = iota

	// PersistentGatheringTypePasswordLocked indicates that the PersistentGathering requires a password
	PersistentGatheringTypePasswordLocked

	// PersistentGatheringTypeOfficial indicates that the PersistentGathering is official
	PersistentGatheringTypeOfficial
)
