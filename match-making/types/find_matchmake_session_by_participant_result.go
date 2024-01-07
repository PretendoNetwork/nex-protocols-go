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

// FindMatchmakeSessionByParticipantResult holds parameters for a matchmake session
type FindMatchmakeSessionByParticipantResult struct {
	types.Structure
	PrincipalID *types.PID
	Session     *MatchmakeSession
}

// ExtractFrom extracts the FindMatchmakeSessionByParticipantResult from the given readable
func (findMatchmakeSessionByParticipantResult *FindMatchmakeSessionByParticipantResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = findMatchmakeSessionByParticipantResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read FindMatchmakeSessionByParticipantResult header. %s", err.Error())
	}

	err = findMatchmakeSessionByParticipantResult.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantResult.PrincipalID. %s", err.Error())
	}

	err = findMatchmakeSessionByParticipantResult.Session.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantResult.Session. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FindMatchmakeSessionByParticipantResult
func (findMatchmakeSessionByParticipantResult *FindMatchmakeSessionByParticipantResult) Copy() types.RVType {
	copied := NewFindMatchmakeSessionByParticipantResult()

	copied.StructureVersion = findMatchmakeSessionByParticipantResult.StructureVersion

	copied.PrincipalID = findMatchmakeSessionByParticipantResult.PrincipalID.Copy()

	copied.Session = findMatchmakeSessionByParticipantResult.Session.Copy().(*MatchmakeSession)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (findMatchmakeSessionByParticipantResult *FindMatchmakeSessionByParticipantResult) Equals(o types.RVType) bool {
	if _, ok := o.(*FindMatchmakeSessionByParticipantResult); !ok {
		return false
	}

	other := o.(*FindMatchmakeSessionByParticipantResult)

	if findMatchmakeSessionByParticipantResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !findMatchmakeSessionByParticipantResult.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if findMatchmakeSessionByParticipantResult.Session != nil && other.Session != nil {
		if findMatchmakeSessionByParticipantResult.Session.Equals(other.Session) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (findMatchmakeSessionByParticipantResult *FindMatchmakeSessionByParticipantResult) String() string {
	return findMatchmakeSessionByParticipantResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (findMatchmakeSessionByParticipantResult *FindMatchmakeSessionByParticipantResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FindMatchmakeSessionByParticipantResult{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, findMatchmakeSessionByParticipantResult.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, findMatchmakeSessionByParticipantResult.PrincipalID.FormatToString(indentationLevel+1)))

	if findMatchmakeSessionByParticipantResult.Session != nil {
		b.WriteString(fmt.Sprintf("%sSession: %s\n", indentationValues, findMatchmakeSessionByParticipantResult.Session.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sSession: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFindMatchmakeSessionByParticipantResult returns a new FindMatchmakeSessionByParticipantResult
func NewFindMatchmakeSessionByParticipantResult() *FindMatchmakeSessionByParticipantResult {
	return &FindMatchmakeSessionByParticipantResult{}
}
