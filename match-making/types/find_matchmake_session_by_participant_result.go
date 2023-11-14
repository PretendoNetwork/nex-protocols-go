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

// FindMatchmakeSessionByParticipantResult holds parameters for a matchmake session
type FindMatchmakeSessionByParticipantResult struct {
	nex.Structure
	PrincipalID *nex.PID
	Session     *MatchmakeSession
}

// ExtractFromStream extracts a FindMatchmakeSessionByParticipantResult structure from a stream
func (findMatchmakeSessionByParticipantResult *FindMatchmakeSessionByParticipantResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	findMatchmakeSessionByParticipantResult.PrincipalID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantResult.PrincipalID. %s", err.Error())
	}

	session, err := stream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantResult.Session. %s", err.Error())
	}

	findMatchmakeSessionByParticipantResult.Session = session.(*MatchmakeSession)

	return nil
}

// Copy returns a new copied instance of FindMatchmakeSessionByParticipantResult
func (findMatchmakeSessionByParticipantResult *FindMatchmakeSessionByParticipantResult) Copy() nex.StructureInterface {
	copied := NewFindMatchmakeSessionByParticipantResult()

	copied.SetStructureVersion(findMatchmakeSessionByParticipantResult.StructureVersion())

	copied.PrincipalID = findMatchmakeSessionByParticipantResult.PrincipalID.Copy()

	if findMatchmakeSessionByParticipantResult.Session != nil {
		copied.Session = findMatchmakeSessionByParticipantResult.Session.Copy().(*MatchmakeSession)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (findMatchmakeSessionByParticipantResult *FindMatchmakeSessionByParticipantResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FindMatchmakeSessionByParticipantResult)

	if findMatchmakeSessionByParticipantResult.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !findMatchmakeSessionByParticipantResult.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if findMatchmakeSessionByParticipantResult.Session != nil && other.Session == nil {
		return false
	}

	if findMatchmakeSessionByParticipantResult.Session == nil && other.Session != nil {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, findMatchmakeSessionByParticipantResult.StructureVersion()))
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
