// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MatchmakeRefereeStatsTarget contains the results of a round
type MatchmakeRefereeStatsTarget struct {
	nex.Structure
	*nex.Data
	PID      uint32
	Category uint32
}

// Bytes encodes the MatchmakeRefereeStatsTarget and returns a byte array
func (matchmakeRefereeStatsTarget *MatchmakeRefereeStatsTarget) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(matchmakeRefereeStatsTarget.PID)
	stream.WriteUInt32LE(matchmakeRefereeStatsTarget.Category)

	return stream.Bytes()
}

// ExtractFromStream extracts a MatchmakeRefereeStatsTarget structure from a stream
func (matchmakeRefereeStatsTarget *MatchmakeRefereeStatsTarget) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	matchmakeRefereeStatsTarget.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsTarget.PID. %s", err.Error())
	}

	matchmakeRefereeStatsTarget.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsTarget.Category. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStatsTarget
func (matchmakeRefereeStatsTarget *MatchmakeRefereeStatsTarget) Copy() nex.StructureInterface {
	copied := NewMatchmakeRefereeStatsTarget()

	copied.SetStructureVersion(matchmakeRefereeStatsTarget.StructureVersion())

	copied.Data = matchmakeRefereeStatsTarget.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.PID = matchmakeRefereeStatsTarget.PID
	copied.Category = matchmakeRefereeStatsTarget.Category

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereeStatsTarget *MatchmakeRefereeStatsTarget) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeRefereeStatsTarget)

	if matchmakeRefereeStatsTarget.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !matchmakeRefereeStatsTarget.ParentType().Equals(other.ParentType()) {
		return false
	}

	if matchmakeRefereeStatsTarget.PID != other.PID {
		return false
	}

	if matchmakeRefereeStatsTarget.Category != other.Category {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (matchmakeRefereeStatsTarget *MatchmakeRefereeStatsTarget) String() string {
	return matchmakeRefereeStatsTarget.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeRefereeStatsTarget *MatchmakeRefereeStatsTarget) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeStatsTarget{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, matchmakeRefereeStatsTarget.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, matchmakeRefereeStatsTarget.PID))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, matchmakeRefereeStatsTarget.Category))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeStatsTarget returns a new MatchmakeRefereeStatsTarget
func NewMatchmakeRefereeStatsTarget() *MatchmakeRefereeStatsTarget {
	return &MatchmakeRefereeStatsTarget{}
}
