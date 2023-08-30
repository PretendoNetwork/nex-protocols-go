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

// MatchmakeSessionSearchCriteria holds information about a matchmaking search
type MatchmakeSessionSearchCriteria struct {
	nex.Structure
	Attribs                  []string
	GameMode                 string
	MinParticipants          string
	MaxParticipants          string
	MatchmakeSystemType      string
	VacantOnly               bool
	ExcludeLocked            bool
	ExcludeNonHostPID        bool
	SelectionMethod          uint32           // NEX v3.0.0+
	VacantParticipants       uint16           // NEX v3.4.0+
	MatchmakeParam           *MatchmakeParam  // NEX v3.6.0+
	ExcludeUserPasswordSet   bool             // NEX v3.7.0+
	ExcludeSystemPasswordSet bool             // NEX v3.7.0+
	ReferGID                 uint32           // NEX v3.8.0+
	CodeWord                 string           // NEX v4.0.0+
	ResultRange              *nex.ResultRange // NEX v4.0.0+
}

// ExtractFromStream extracts a Gathering structure from a stream
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) ExtractFromStream(stream *nex.StreamIn) error {
	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	var err error

	matchmakeSessionSearchCriteria.Attribs, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.Attribs. %s", err.Error())
	}

	matchmakeSessionSearchCriteria.GameMode, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.GameMode. %s", err.Error())
	}

	matchmakeSessionSearchCriteria.MinParticipants, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MinParticipants. %s", err.Error())
	}

	matchmakeSessionSearchCriteria.MaxParticipants, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MaxParticipants. %s", err.Error())
	}

	matchmakeSessionSearchCriteria.MatchmakeSystemType, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MatchmakeSystemType. %s", err.Error())
	}

	matchmakeSessionSearchCriteria.VacantOnly, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.VacantOnly. %s", err.Error())
	}

	matchmakeSessionSearchCriteria.ExcludeLocked, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeLocked. %s", err.Error())
	}

	matchmakeSessionSearchCriteria.ExcludeNonHostPID, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeNonHostPID. %s", err.Error())
	}

	if matchmakingVersion.GreaterOrEqual("3.0.0") {
		matchmakeSessionSearchCriteria.SelectionMethod, err = stream.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.SelectionMethod. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
		matchmakeSessionSearchCriteria.VacantParticipants, err = stream.ReadUInt16LE()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.VacantParticipants. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.6.0") {
		matchmakeParam, err := stream.ReadStructure(NewMatchmakeParam())
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MatchmakeParam. %s", err.Error())
		}

		matchmakeSessionSearchCriteria.MatchmakeParam = matchmakeParam.(*MatchmakeParam)
	}

	if matchmakingVersion.GreaterOrEqual("3.7.0") {
		matchmakeSessionSearchCriteria.ExcludeUserPasswordSet, err = stream.ReadBool()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeUserPasswordSet. %s", err.Error())
		}

		matchmakeSessionSearchCriteria.ExcludeSystemPasswordSet, err = stream.ReadBool()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeSystemPasswordSet. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.8.0") {
		matchmakeSessionSearchCriteria.ReferGID, err = stream.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ReferGID. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("4.0.0") {
		matchmakeSessionSearchCriteria.CodeWord, err = stream.ReadString()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.CodeWord. %s", err.Error())
		}

		resultRange, err := stream.ReadStructure(nex.NewResultRange())
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ResultRange. %s", err.Error())
		}

		matchmakeSessionSearchCriteria.ResultRange = resultRange.(*nex.ResultRange)
	}

	return nil
}

// Bytes encodes the Gathering and returns a byte array
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) Bytes(stream *nex.StreamOut) []byte {
	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	stream.WriteListString(matchmakeSessionSearchCriteria.Attribs)
	stream.WriteString(matchmakeSessionSearchCriteria.GameMode)
	stream.WriteString(matchmakeSessionSearchCriteria.MinParticipants)
	stream.WriteString(matchmakeSessionSearchCriteria.MaxParticipants)
	stream.WriteString(matchmakeSessionSearchCriteria.MatchmakeSystemType)
	stream.WriteBool(matchmakeSessionSearchCriteria.VacantOnly)
	stream.WriteBool(matchmakeSessionSearchCriteria.ExcludeLocked)
	stream.WriteBool(matchmakeSessionSearchCriteria.ExcludeNonHostPID)

	if matchmakingVersion.GreaterOrEqual("3.0.0") {
		stream.WriteUInt32LE(matchmakeSessionSearchCriteria.SelectionMethod)
	}

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
		stream.WriteUInt16LE(matchmakeSessionSearchCriteria.VacantParticipants)
	}

	if matchmakingVersion.GreaterOrEqual("3.6.0") {
		stream.WriteStructure(matchmakeSessionSearchCriteria.MatchmakeParam)
	}

	if matchmakingVersion.GreaterOrEqual("3.7.0") {
		stream.WriteBool(matchmakeSessionSearchCriteria.ExcludeUserPasswordSet)
		stream.WriteBool(matchmakeSessionSearchCriteria.ExcludeSystemPasswordSet)
	}

	if matchmakingVersion.GreaterOrEqual("3.8.0") {
		stream.WriteUInt32LE(matchmakeSessionSearchCriteria.ReferGID)
	}

	if matchmakingVersion.GreaterOrEqual("4.0.0") {
		stream.WriteString(matchmakeSessionSearchCriteria.CodeWord)
		stream.WriteStructure(matchmakeSessionSearchCriteria.ResultRange)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of Gathering
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) Copy() nex.StructureInterface {
	copied := NewMatchmakeSessionSearchCriteria()

	copied.SetStructureVersion(matchmakeSessionSearchCriteria.StructureVersion())

	copied.Attribs = make([]string, len(matchmakeSessionSearchCriteria.Attribs))

	copy(copied.Attribs, matchmakeSessionSearchCriteria.Attribs)

	copied.GameMode = matchmakeSessionSearchCriteria.GameMode
	copied.MinParticipants = matchmakeSessionSearchCriteria.MinParticipants
	copied.MaxParticipants = matchmakeSessionSearchCriteria.MaxParticipants
	copied.MatchmakeSystemType = matchmakeSessionSearchCriteria.MatchmakeSystemType
	copied.VacantOnly = matchmakeSessionSearchCriteria.VacantOnly
	copied.ExcludeLocked = matchmakeSessionSearchCriteria.ExcludeLocked
	copied.ExcludeNonHostPID = matchmakeSessionSearchCriteria.ExcludeNonHostPID
	copied.SelectionMethod = matchmakeSessionSearchCriteria.SelectionMethod
	copied.VacantParticipants = matchmakeSessionSearchCriteria.VacantParticipants

	if matchmakeSessionSearchCriteria.MatchmakeParam != nil {
		copied.MatchmakeParam = matchmakeSessionSearchCriteria.MatchmakeParam.Copy().(*MatchmakeParam)
	}

	copied.ExcludeUserPasswordSet = matchmakeSessionSearchCriteria.ExcludeUserPasswordSet
	copied.ExcludeSystemPasswordSet = matchmakeSessionSearchCriteria.ExcludeSystemPasswordSet
	copied.ReferGID = matchmakeSessionSearchCriteria.ReferGID
	copied.CodeWord = matchmakeSessionSearchCriteria.CodeWord

	if matchmakeSessionSearchCriteria.ResultRange != nil {
		copied.ResultRange = matchmakeSessionSearchCriteria.ResultRange.Copy().(*nex.ResultRange)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeSessionSearchCriteria)

	if matchmakeSessionSearchCriteria.StructureVersion() != other.StructureVersion() {
		return false
	}

	if len(matchmakeSessionSearchCriteria.Attribs) != len(other.Attribs) {
		return false
	}

	for i := 0; i < len(matchmakeSessionSearchCriteria.Attribs); i++ {
		if matchmakeSessionSearchCriteria.Attribs[i] != other.Attribs[i] {
			return false
		}
	}

	if matchmakeSessionSearchCriteria.GameMode != other.GameMode {
		return false
	}

	if matchmakeSessionSearchCriteria.MinParticipants != other.MinParticipants {
		return false
	}

	if matchmakeSessionSearchCriteria.MaxParticipants != other.MaxParticipants {
		return false
	}

	if matchmakeSessionSearchCriteria.MatchmakeSystemType != other.MatchmakeSystemType {
		return false
	}

	if matchmakeSessionSearchCriteria.VacantOnly != other.VacantOnly {
		return false
	}

	if matchmakeSessionSearchCriteria.ExcludeLocked != other.ExcludeLocked {
		return false
	}

	if matchmakeSessionSearchCriteria.ExcludeNonHostPID != other.ExcludeNonHostPID {
		return false
	}

	if matchmakeSessionSearchCriteria.SelectionMethod != other.SelectionMethod {
		return false
	}

	if matchmakeSessionSearchCriteria.VacantParticipants != other.VacantParticipants {
		return false
	}

	if matchmakeSessionSearchCriteria.MatchmakeParam != nil && other.MatchmakeParam == nil {
		return false
	}

	if matchmakeSessionSearchCriteria.MatchmakeParam == nil && other.MatchmakeParam != nil {
		return false
	}

	if matchmakeSessionSearchCriteria.MatchmakeParam != nil && other.MatchmakeParam != nil {
		if !matchmakeSessionSearchCriteria.MatchmakeParam.Equals(other.MatchmakeParam) {
			return false
		}
	}

	if matchmakeSessionSearchCriteria.ExcludeUserPasswordSet != other.ExcludeUserPasswordSet {
		return false
	}

	if matchmakeSessionSearchCriteria.ExcludeSystemPasswordSet != other.ExcludeSystemPasswordSet {
		return false
	}

	if matchmakeSessionSearchCriteria.ReferGID != other.ReferGID {
		return false
	}

	if matchmakeSessionSearchCriteria.CodeWord != other.CodeWord {
		return false
	}

	if matchmakeSessionSearchCriteria.ResultRange != nil && other.ResultRange == nil {
		return false
	}

	if matchmakeSessionSearchCriteria.ResultRange == nil && other.ResultRange != nil {
		return false
	}

	if matchmakeSessionSearchCriteria.ResultRange != nil && other.ResultRange != nil {
		if !matchmakeSessionSearchCriteria.ResultRange.Equals(other.ResultRange) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) String() string {
	return matchmakeSessionSearchCriteria.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeSessionSearchCriteria{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, matchmakeSessionSearchCriteria.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sAttribs: %v,\n", indentationValues, matchmakeSessionSearchCriteria.Attribs))
	b.WriteString(fmt.Sprintf("%sGameMode: %q,\n", indentationValues, matchmakeSessionSearchCriteria.GameMode))
	b.WriteString(fmt.Sprintf("%sMinParticipants: %q,\n", indentationValues, matchmakeSessionSearchCriteria.MinParticipants))
	b.WriteString(fmt.Sprintf("%sMaxParticipants: %q,\n", indentationValues, matchmakeSessionSearchCriteria.MaxParticipants))
	b.WriteString(fmt.Sprintf("%sMatchmakeSystemType: %q,\n", indentationValues, matchmakeSessionSearchCriteria.MatchmakeSystemType))
	b.WriteString(fmt.Sprintf("%sVacantOnly: %t,\n", indentationValues, matchmakeSessionSearchCriteria.VacantOnly))
	b.WriteString(fmt.Sprintf("%sExcludeLocked: %t,\n", indentationValues, matchmakeSessionSearchCriteria.ExcludeLocked))
	b.WriteString(fmt.Sprintf("%sExcludeNonHostPID: %t,\n", indentationValues, matchmakeSessionSearchCriteria.ExcludeNonHostPID))
	b.WriteString(fmt.Sprintf("%sSelectionMethod: %d,\n", indentationValues, matchmakeSessionSearchCriteria.SelectionMethod))
	b.WriteString(fmt.Sprintf("%sVacantParticipants: %d,\n", indentationValues, matchmakeSessionSearchCriteria.VacantParticipants))

	if matchmakeSessionSearchCriteria.MatchmakeParam != nil {
		b.WriteString(fmt.Sprintf("%sMatchmakeParam: %s,\n", indentationValues, matchmakeSessionSearchCriteria.MatchmakeParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMatchmakeParam: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sExcludeUserPasswordSet: %t,\n", indentationValues, matchmakeSessionSearchCriteria.ExcludeUserPasswordSet))
	b.WriteString(fmt.Sprintf("%sExcludeSystemPasswordSet: %t,\n", indentationValues, matchmakeSessionSearchCriteria.ExcludeSystemPasswordSet))
	b.WriteString(fmt.Sprintf("%sReferGID: %d,\n", indentationValues, matchmakeSessionSearchCriteria.ReferGID))
	b.WriteString(fmt.Sprintf("%sCodeWord: %q\n", indentationValues, matchmakeSessionSearchCriteria.CodeWord))

	if matchmakeSessionSearchCriteria.ResultRange != nil {
		b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, matchmakeSessionSearchCriteria.ResultRange.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResultRange: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeSessionSearchCriteria returns a new MatchmakeSessionSearchCriteria
func NewMatchmakeSessionSearchCriteria() *MatchmakeSessionSearchCriteria {
	return &MatchmakeSessionSearchCriteria{}
}
