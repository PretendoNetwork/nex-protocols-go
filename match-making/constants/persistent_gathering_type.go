package constants

// PersistentGatheringType indicates the type of PersistentGathering
type PersistentGatheringType uint32

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
