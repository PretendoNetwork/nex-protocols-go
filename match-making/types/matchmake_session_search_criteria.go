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

// MatchmakeSessionSearchCriteria holds information about a matchmaking search
type MatchmakeSessionSearchCriteria struct {
	nex.Structure
	Attribs             []string
	GameMode            string
	MinParticipants     string
	MaxParticipants     string
	MatchmakeSystemType string
	VacantOnly          bool
	ExcludeLocked       bool
	ExcludeNonHostPID   bool
	SelectionMethod     uint32 // NEX v3.0.0+
	VacantParticipants  uint16 // NEX v3.4.0+
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

	if matchmakingVersion.Major >= 3 {
		matchmakeSessionSearchCriteria.SelectionMethod, err = stream.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.SelectionMethod. %s", err.Error())
		}
	}


	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 4 {
		matchmakeSessionSearchCriteria.VacantParticipants, err = stream.ReadUInt16LE()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.VacantParticipants. %s", err.Error())
		}
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

	if matchmakingVersion.Major >= 3 {
		stream.WriteUInt32LE(matchmakeSessionSearchCriteria.SelectionMethod)
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 4 {
		stream.WriteUInt16LE(matchmakeSessionSearchCriteria.VacantParticipants)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of Gathering
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) Copy() nex.StructureInterface {
	copied := NewMatchmakeSessionSearchCriteria()

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

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeSessionSearchCriteria)

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
	b.WriteString(fmt.Sprintf("%sVacantParticipants: %d\n", indentationValues, matchmakeSessionSearchCriteria.VacantParticipants))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeSessionSearchCriteria returns a new MatchmakeSessionSearchCriteria
func NewMatchmakeSessionSearchCriteria() *MatchmakeSessionSearchCriteria {
	return &MatchmakeSessionSearchCriteria{}
}
