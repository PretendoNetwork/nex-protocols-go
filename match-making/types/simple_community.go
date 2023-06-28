package match_making_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// SimpleCommunity holds basic info about a community
type SimpleCommunity struct {
	nex.Structure
	M_GatheringID           uint32
	M_MatchmakeSessionCount uint32
}

// ExtractFromStream extracts a SimpleCommunity structure from a stream
func (simpleCommunity *SimpleCommunity) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	simpleCommunity.M_GatheringID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleCommunity.M_GatheringID. %s", err.Error())
	}

	simpleCommunity.M_MatchmakeSessionCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleCommunity.M_MatchmakeSessionCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SimpleCommunity
func (simpleCommunity *SimpleCommunity) Copy() nex.StructureInterface {
	copied := NewSimpleCommunity()

	copied.M_GatheringID = simpleCommunity.M_GatheringID
	copied.M_MatchmakeSessionCount = simpleCommunity.M_MatchmakeSessionCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleCommunity *SimpleCommunity) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SimpleCommunity)

	if simpleCommunity.M_GatheringID != other.M_GatheringID {
		return false
	}

	if simpleCommunity.M_MatchmakeSessionCount != other.M_MatchmakeSessionCount {
		return false
	}

	return true
}

// NewSimpleCommunity returns a new SimpleCommunity
func NewSimpleCommunity() *SimpleCommunity {
	return &SimpleCommunity{}
}
