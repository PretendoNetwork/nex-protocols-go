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

// GatheringStats holds stats about a Gathering
type GatheringStats struct {
	types.Structure
	PIDParticipant *types.PID
	UIFlags        *types.PrimitiveU32
	LstValues      *types.List[*types.PrimitiveF32]
}

// ExtractFrom extracts the GatheringStats from the given readable
func (gatheringStats *GatheringStats) ExtractFrom(readable types.Readable) error {
	var err error

	if err = gatheringStats.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GatheringStats header. %s", err.Error())
	}

	err = gatheringStats.PIDParticipant.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringStats.PIDParticipant. %s", err.Error())
	}

	err = gatheringStats.UIFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringStats.UIFlags. %s", err.Error())
	}

	err = gatheringStats.LstValues.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringStats.LstValues. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GatheringStats to the given writable
func (gatheringStats *GatheringStats) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gatheringStats.PIDParticipant.WriteTo(contentWritable)
	gatheringStats.UIFlags.WriteTo(contentWritable)
	gatheringStats.LstValues.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gatheringStats.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GatheringStats
func (gatheringStats *GatheringStats) Copy() types.RVType {
	copied := NewGatheringStats()

	copied.StructureVersion = gatheringStats.StructureVersion

	copied.PIDParticipant = gatheringStats.PIDParticipant.Copy()
	copied.UIFlags = gatheringStats.UIFlags
	copied.LstValues = make(*types.List[*types.PrimitiveF32], len(gatheringStats.LstValues))

	copy(copied.LstValues, gatheringStats.LstValues)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (gatheringStats *GatheringStats) Equals(o types.RVType) bool {
	if _, ok := o.(*GatheringStats); !ok {
		return false
	}

	other := o.(*GatheringStats)

	if gatheringStats.StructureVersion != other.StructureVersion {
		return false
	}

	if !gatheringStats.PIDParticipant.Equals(other.PIDParticipant) {
		return false
	}

	if !gatheringStats.UIFlags.Equals(other.UIFlags) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, gatheringStats.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPIDParticipant: %s,\n", indentationValues, gatheringStats.PIDParticipant.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUIFlags: %d,\n", indentationValues, gatheringStats.UIFlags))
	b.WriteString(fmt.Sprintf("%sLstValues: %v\n", indentationValues, gatheringStats.LstValues))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGatheringStats returns a new GatheringStats
func NewGatheringStats() *GatheringStats {
	return &GatheringStats{}
}
