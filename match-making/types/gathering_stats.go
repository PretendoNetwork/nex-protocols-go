// Package match_making_types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package match_making_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GatheringStats holds stats about a Gathering
type GatheringStats struct {
	nex.Structure
	PIDParticipant uint32
	UIFlags        uint32
	LstValues      []float32
}

// ExtractFromStream extracts a GatheringStats structure from a stream
func (gatheringStats *GatheringStats) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	gatheringStats.PIDParticipant, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringStats.PIDParticipant. %s", err.Error())
	}

	gatheringStats.UIFlags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringStats.UIFlags. %s", err.Error())
	}

	gatheringStats.LstValues, err = stream.ReadListFloat32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringStats.LstValues. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GatheringStats and returns a byte array
func (gatheringStats *GatheringStats) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(gatheringStats.PIDParticipant)
	stream.WriteUInt32LE(gatheringStats.UIFlags)
	stream.WriteListFloat32LE(gatheringStats.LstValues)

	return stream.Bytes()
}

// Copy returns a new copied instance of GatheringStats
func (gatheringStats *GatheringStats) Copy() nex.StructureInterface {
	copied := NewGatheringStats()

	copied.PIDParticipant = gatheringStats.PIDParticipant
	copied.UIFlags = gatheringStats.UIFlags
	copied.LstValues = make([]float32, len(gatheringStats.LstValues))

	copy(copied.LstValues, gatheringStats.LstValues)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (gatheringStats *GatheringStats) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GatheringStats)

	if gatheringStats.PIDParticipant != other.PIDParticipant {
		return false
	}

	if gatheringStats.UIFlags != other.UIFlags {
		return false
	}

	if len(gatheringStats.LstValues) != len(other.LstValues) {
		return false
	}

	for i := 0; i < len(gatheringStats.LstValues); i++ {
		if gatheringStats.LstValues[i] != other.LstValues[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (gatheringStats *GatheringStats) String() string {
	return gatheringStats.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (gatheringStats *GatheringStats) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GatheringStats{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, gatheringStats.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPIDParticipant: %d,\n", indentationValues, gatheringStats.PIDParticipant))
	b.WriteString(fmt.Sprintf("%sUIFlags: %d,\n", indentationValues, gatheringStats.UIFlags))
	b.WriteString(fmt.Sprintf("%sLstValues: %v\n", indentationValues, gatheringStats.LstValues))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGatheringStats returns a new GatheringStats
func NewGatheringStats() *GatheringStats {
	return &GatheringStats{}
}
