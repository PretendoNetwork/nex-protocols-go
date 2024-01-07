// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// SimpleCommunity holds basic info about a community
type SimpleCommunity struct {
	types.Structure
	GatheringID           *types.PrimitiveU32
	MatchmakeSessionCount *types.PrimitiveU32
}

// ExtractFrom extracts the SimpleCommunity from the given readable
func (simpleCommunity *SimpleCommunity) ExtractFrom(readable types.Readable) error {
	var err error

	if err = simpleCommunity.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read SimpleCommunity header. %s", err.Error())
	}

	err = simpleCommunity.GatheringID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleCommunity.GatheringID. %s", err.Error())
	}

	err = simpleCommunity.MatchmakeSessionCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleCommunity.MatchmakeSessionCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SimpleCommunity
func (simpleCommunity *SimpleCommunity) Copy() types.RVType {
	copied := NewSimpleCommunity()

	copied.StructureVersion = simpleCommunity.StructureVersion

	copied.GatheringID = simpleCommunity.GatheringID
	copied.MatchmakeSessionCount = simpleCommunity.MatchmakeSessionCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleCommunity *SimpleCommunity) Equals(o types.RVType) bool {
	if _, ok := o.(*SimpleCommunity); !ok {
		return false
	}

	other := o.(*SimpleCommunity)

	if simpleCommunity.StructureVersion != other.StructureVersion {
		return false
	}

	if !simpleCommunity.GatheringID.Equals(other.GatheringID) {
		return false
	}

	if !simpleCommunity.MatchmakeSessionCount.Equals(other.MatchmakeSessionCount) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, simpleCommunity.StructureVersion))
	b.WriteString(fmt.Sprintf("%sGatheringID: %d,\n", indentationValues, simpleCommunity.GatheringID))
	b.WriteString(fmt.Sprintf("%sMatchmakeSessionCount: %d\n", indentationValues, simpleCommunity.MatchmakeSessionCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleCommunity returns a new SimpleCommunity
func NewSimpleCommunity() *SimpleCommunity {
	return &SimpleCommunity{}
}
