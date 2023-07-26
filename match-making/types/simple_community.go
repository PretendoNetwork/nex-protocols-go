// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// SimpleCommunity holds basic info about a community
type SimpleCommunity struct {
	nex.Structure
	GatheringID           uint32
	MatchmakeSessionCount uint32
}

// ExtractFromStream extracts a SimpleCommunity structure from a stream
func (simpleCommunity *SimpleCommunity) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	simpleCommunity.GatheringID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleCommunity.GatheringID. %s", err.Error())
	}

	simpleCommunity.MatchmakeSessionCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleCommunity.MatchmakeSessionCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SimpleCommunity
func (simpleCommunity *SimpleCommunity) Copy() nex.StructureInterface {
	copied := NewSimpleCommunity()

	copied.GatheringID = simpleCommunity.GatheringID
	copied.MatchmakeSessionCount = simpleCommunity.MatchmakeSessionCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleCommunity *SimpleCommunity) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SimpleCommunity)

	if simpleCommunity.GatheringID != other.GatheringID {
		return false
	}

	if simpleCommunity.MatchmakeSessionCount != other.MatchmakeSessionCount {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (simpleCommunity *SimpleCommunity) String() string {
	return simpleCommunity.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (simpleCommunity *SimpleCommunity) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeSession{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, simpleCommunity.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sGatheringID: %d,\n", indentationValues, simpleCommunity.GatheringID))
	b.WriteString(fmt.Sprintf("%sMatchmakeSessionCount: %d\n", indentationValues, simpleCommunity.MatchmakeSessionCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleCommunity returns a new SimpleCommunity
func NewSimpleCommunity() *SimpleCommunity {
	return &SimpleCommunity{}
}
