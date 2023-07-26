// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// PersistentGathering holds parameters for a matchmake session
type PersistentGathering struct {
	nex.Structure
	*Gathering
	M_CommunityType          uint32
	M_Password               string
	M_Attribs                []uint32
	M_ApplicationBuffer      []byte
	M_ParticipationStartDate *nex.DateTime
	M_ParticipationEndDate   *nex.DateTime
	M_MatchmakeSessionCount  uint32
	M_ParticipationCount     uint32
}

// ExtractFromStream extracts a PersistentGathering structure from a stream
func (persistentGathering *PersistentGathering) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	persistentGathering.M_CommunityType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.M_CommunityType. %s", err.Error())
	}

	persistentGathering.M_Password, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.M_Password. %s", err.Error())
	}

	persistentGathering.M_Attribs, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.M_Attribs. %s", err.Error())
	}

	persistentGathering.M_ApplicationBuffer, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.M_ApplicationBuffer. %s", err.Error())
	}

	persistentGathering.M_ParticipationStartDate, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.M_ParticipationStartDate. %s", err.Error())
	}

	persistentGathering.M_ParticipationEndDate, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.M_ParticipationEndDate. %s", err.Error())
	}

	persistentGathering.M_MatchmakeSessionCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.M_MatchmakeSessionCount. %s", err.Error())
	}

	persistentGathering.M_ParticipationCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.M_ParticipationCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PersistentGathering
func (persistentGathering *PersistentGathering) Copy() nex.StructureInterface {
	copied := NewPersistentGathering()

	copied.Gathering = persistentGathering.Gathering.Copy().(*Gathering)
	copied.SetParentType(copied.Gathering)
	copied.M_CommunityType = persistentGathering.M_CommunityType
	copied.M_Password = persistentGathering.M_Password
	copied.M_Attribs = make([]uint32, len(persistentGathering.M_Attribs))

	copy(copied.M_Attribs, persistentGathering.M_Attribs)

	copied.M_ApplicationBuffer = make([]byte, len(persistentGathering.M_ApplicationBuffer))

	copy(copied.M_ApplicationBuffer, persistentGathering.M_ApplicationBuffer)

	if persistentGathering.M_ParticipationStartDate != nil {
		copied.M_ParticipationStartDate = persistentGathering.M_ParticipationStartDate.Copy()
	}

	if persistentGathering.M_ParticipationEndDate != nil {
		copied.M_ParticipationEndDate = persistentGathering.M_ParticipationEndDate.Copy()
	}

	copied.M_MatchmakeSessionCount = persistentGathering.M_MatchmakeSessionCount
	copied.M_ParticipationCount = persistentGathering.M_ParticipationCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (persistentGathering *PersistentGathering) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PersistentGathering)

	if !persistentGathering.ParentType().Equals(other.ParentType()) {
		return false
	}

	if persistentGathering.M_CommunityType != other.M_CommunityType {
		return false
	}

	if persistentGathering.M_Password != other.M_Password {
		return false
	}

	if len(persistentGathering.M_Attribs) != len(other.M_Attribs) {
		return false
	}

	for i := 0; i < len(persistentGathering.M_Attribs); i++ {
		if persistentGathering.M_Attribs[i] != other.M_Attribs[i] {
			return false
		}
	}

	if !bytes.Equal(persistentGathering.M_ApplicationBuffer, other.M_ApplicationBuffer) {
		return false
	}

	if persistentGathering.M_ParticipationStartDate != nil && other.M_ParticipationStartDate == nil {
		return false
	}

	if persistentGathering.M_ParticipationStartDate == nil && other.M_ParticipationStartDate != nil {
		return false
	}

	if persistentGathering.M_ParticipationStartDate != nil && other.M_ParticipationStartDate != nil {
		if persistentGathering.M_ParticipationStartDate.Equals(other.M_ParticipationStartDate) {
			return false
		}
	}

	if persistentGathering.M_ParticipationEndDate != nil && other.M_ParticipationEndDate == nil {
		return false
	}

	if persistentGathering.M_ParticipationEndDate == nil && other.M_ParticipationEndDate != nil {
		return false
	}

	if persistentGathering.M_ParticipationEndDate != nil && other.M_ParticipationEndDate != nil {
		if persistentGathering.M_ParticipationEndDate.Equals(other.M_ParticipationEndDate) {
			return false
		}
	}

	if persistentGathering.M_MatchmakeSessionCount != other.M_MatchmakeSessionCount {
		return false
	}

	if persistentGathering.M_ParticipationCount != other.M_ParticipationCount {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (persistentGathering *PersistentGathering) String() string {
	return persistentGathering.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (persistentGathering *PersistentGathering) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PersistentGathering{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, persistentGathering.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, persistentGathering.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sM_CommunityType: %d,\n", indentationValues, persistentGathering.M_CommunityType))
	b.WriteString(fmt.Sprintf("%sM_Password: %q,\n", indentationValues, persistentGathering.M_Password))
	b.WriteString(fmt.Sprintf("%sM_Attribs: %v,\n", indentationValues, persistentGathering.M_Attribs))
	b.WriteString(fmt.Sprintf("%sM_ApplicationBuffer: %x,\n", indentationValues, persistentGathering.M_ApplicationBuffer))

	if persistentGathering.M_ParticipationStartDate != nil {
		b.WriteString(fmt.Sprintf("%sM_ParticipationStartDate: %s,\n", indentationValues, persistentGathering.M_ParticipationStartDate.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sM_ParticipationStartDate: nil,\n", indentationValues))
	}

	if persistentGathering.M_ParticipationEndDate != nil {
		b.WriteString(fmt.Sprintf("%sM_ParticipationEndDate: %s,\n", indentationValues, persistentGathering.M_ParticipationEndDate.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sM_ParticipationEndDate: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sM_MatchmakeSessionCount: %d,\n", indentationValues, persistentGathering.M_MatchmakeSessionCount))
	b.WriteString(fmt.Sprintf("%sM_ParticipationCount: %d\n", indentationValues, persistentGathering.M_ParticipationCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPersistentGathering returns a new PersistentGathering
func NewPersistentGathering() *PersistentGathering {
	persistentGathering := &PersistentGathering{}
	persistentGathering.Gathering = NewGathering()
	persistentGathering.SetParentType(persistentGathering.Gathering)

	return persistentGathering
}
