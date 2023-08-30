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
	CommunityType          uint32
	Password               string
	Attribs                []uint32
	ApplicationBuffer      []byte
	ParticipationStartDate *nex.DateTime
	ParticipationEndDate   *nex.DateTime
	MatchmakeSessionCount  uint32
	ParticipationCount     uint32
}

// ExtractFromStream extracts a PersistentGathering structure from a stream
func (persistentGathering *PersistentGathering) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	persistentGathering.CommunityType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.CommunityType. %s", err.Error())
	}

	persistentGathering.Password, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.Password. %s", err.Error())
	}

	persistentGathering.Attribs, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.Attribs. %s", err.Error())
	}

	persistentGathering.ApplicationBuffer, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.ApplicationBuffer. %s", err.Error())
	}

	persistentGathering.ParticipationStartDate, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.ParticipationStartDate. %s", err.Error())
	}

	persistentGathering.ParticipationEndDate, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.ParticipationEndDate. %s", err.Error())
	}

	persistentGathering.MatchmakeSessionCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.MatchmakeSessionCount. %s", err.Error())
	}

	persistentGathering.ParticipationCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.ParticipationCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PersistentGathering
func (persistentGathering *PersistentGathering) Copy() nex.StructureInterface {
	copied := NewPersistentGathering()

	copied.SetStructureVersion(persistentGathering.StructureVersion())

	copied.Gathering = persistentGathering.Gathering.Copy().(*Gathering)
	copied.SetParentType(copied.Gathering)
	copied.CommunityType = persistentGathering.CommunityType
	copied.Password = persistentGathering.Password
	copied.Attribs = make([]uint32, len(persistentGathering.Attribs))

	copy(copied.Attribs, persistentGathering.Attribs)

	copied.ApplicationBuffer = make([]byte, len(persistentGathering.ApplicationBuffer))

	copy(copied.ApplicationBuffer, persistentGathering.ApplicationBuffer)

	if persistentGathering.ParticipationStartDate != nil {
		copied.ParticipationStartDate = persistentGathering.ParticipationStartDate.Copy()
	}

	if persistentGathering.ParticipationEndDate != nil {
		copied.ParticipationEndDate = persistentGathering.ParticipationEndDate.Copy()
	}

	copied.MatchmakeSessionCount = persistentGathering.MatchmakeSessionCount
	copied.ParticipationCount = persistentGathering.ParticipationCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (persistentGathering *PersistentGathering) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PersistentGathering)

	if persistentGathering.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !persistentGathering.ParentType().Equals(other.ParentType()) {
		return false
	}

	if persistentGathering.CommunityType != other.CommunityType {
		return false
	}

	if persistentGathering.Password != other.Password {
		return false
	}

	if len(persistentGathering.Attribs) != len(other.Attribs) {
		return false
	}

	for i := 0; i < len(persistentGathering.Attribs); i++ {
		if persistentGathering.Attribs[i] != other.Attribs[i] {
			return false
		}
	}

	if !bytes.Equal(persistentGathering.ApplicationBuffer, other.ApplicationBuffer) {
		return false
	}

	if persistentGathering.ParticipationStartDate != nil && other.ParticipationStartDate == nil {
		return false
	}

	if persistentGathering.ParticipationStartDate == nil && other.ParticipationStartDate != nil {
		return false
	}

	if persistentGathering.ParticipationStartDate != nil && other.ParticipationStartDate != nil {
		if persistentGathering.ParticipationStartDate.Equals(other.ParticipationStartDate) {
			return false
		}
	}

	if persistentGathering.ParticipationEndDate != nil && other.ParticipationEndDate == nil {
		return false
	}

	if persistentGathering.ParticipationEndDate == nil && other.ParticipationEndDate != nil {
		return false
	}

	if persistentGathering.ParticipationEndDate != nil && other.ParticipationEndDate != nil {
		if persistentGathering.ParticipationEndDate.Equals(other.ParticipationEndDate) {
			return false
		}
	}

	if persistentGathering.MatchmakeSessionCount != other.MatchmakeSessionCount {
		return false
	}

	if persistentGathering.ParticipationCount != other.ParticipationCount {
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
	b.WriteString(fmt.Sprintf("%sCommunityType: %d,\n", indentationValues, persistentGathering.CommunityType))
	b.WriteString(fmt.Sprintf("%sPassword: %q,\n", indentationValues, persistentGathering.Password))
	b.WriteString(fmt.Sprintf("%sAttribs: %v,\n", indentationValues, persistentGathering.Attribs))
	b.WriteString(fmt.Sprintf("%sApplicationBuffer: %x,\n", indentationValues, persistentGathering.ApplicationBuffer))

	if persistentGathering.ParticipationStartDate != nil {
		b.WriteString(fmt.Sprintf("%sParticipationStartDate: %s,\n", indentationValues, persistentGathering.ParticipationStartDate.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sParticipationStartDate: nil,\n", indentationValues))
	}

	if persistentGathering.ParticipationEndDate != nil {
		b.WriteString(fmt.Sprintf("%sParticipationEndDate: %s,\n", indentationValues, persistentGathering.ParticipationEndDate.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sParticipationEndDate: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sMatchmakeSessionCount: %d,\n", indentationValues, persistentGathering.MatchmakeSessionCount))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %d\n", indentationValues, persistentGathering.ParticipationCount))
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
