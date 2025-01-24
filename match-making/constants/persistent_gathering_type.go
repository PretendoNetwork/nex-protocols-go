package constants

// PersistentGatheringType indicates the type of PersistentGathering
type PersistentGatheringType uint32

const (
	// PersistentGatheringTypeOpen indicates that the PersistentGathering is open to everyone
	PersistentGatheringTypeOpen PersistentGatheringType = iota

	// PersistentGatheringTypePasswordLocked indicates that the PersistentGathering requires a password
	PersistentGatheringTypePasswordLocked

	// PersistentGatheringTypeOfficial indicates that the PersistentGathering is official
	PersistentGatheringTypeOfficial
)
