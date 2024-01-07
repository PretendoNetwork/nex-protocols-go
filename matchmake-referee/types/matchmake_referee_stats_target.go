// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// MatchmakeRefereeStatsTarget contains the results of a round
type MatchmakeRefereeStatsTarget struct {
	types.Structure
	*types.Data
	PID      *types.PID
	Category *types.PrimitiveU32
}

// WriteTo writes the MatchmakeRefereeStatsTarget to the given writable
func (matchmakeRefereeStatsTarget *MatchmakeRefereeStatsTarget) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	matchmakeRefereeStatsTarget.PID.WriteTo(contentWritable)
	matchmakeRefereeStatsTarget.Category.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	matchmakeRefereeStatsTarget.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeStatsTarget from the given readable
func (matchmakeRefereeStatsTarget *MatchmakeRefereeStatsTarget) ExtractFrom(readable types.Readable) error {
	var err error

	if err = matchmakeRefereeStatsTarget.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read MatchmakeRefereeStatsTarget header. %s", err.Error())
	}

	err = matchmakeRefereeStatsTarget.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsTarget.PID. %s", err.Error())
	}

	err = matchmakeRefereeStatsTarget.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsTarget.Category. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStatsTarget
func (matchmakeRefereeStatsTarget *MatchmakeRefereeStatsTarget) Copy() types.RVType {
	copied := NewMatchmakeRefereeStatsTarget()

	copied.StructureVersion = matchmakeRefereeStatsTarget.StructureVersion

	copied.Data = matchmakeRefereeStatsTarget.Data.Copy().(*types.Data)

	copied.PID = matchmakeRefereeStatsTarget.PID.Copy()
	copied.Category = matchmakeRefereeStatsTarget.Category

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereeStatsTarget *MatchmakeRefereeStatsTarget) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereeStatsTarget); !ok {
		return false
	}

	other := o.(*MatchmakeRefereeStatsTarget)

	if matchmakeRefereeStatsTarget.StructureVersion != other.StructureVersion {
		return false
	}

	if !matchmakeRefereeStatsTarget.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !matchmakeRefereeStatsTarget.PID.Equals(other.PID) {
		return false
	}

	if !matchmakeRefereeStatsTarget.Category.Equals(other.Category) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, matchmakeRefereeStatsTarget.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, matchmakeRefereeStatsTarget.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, matchmakeRefereeStatsTarget.Category))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeStatsTarget returns a new MatchmakeRefereeStatsTarget
func NewMatchmakeRefereeStatsTarget() *MatchmakeRefereeStatsTarget {
	return &MatchmakeRefereeStatsTarget{}
}
