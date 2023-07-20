// Package friends_types implements all the types used by the Friends protocol
package friends_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// RelationshipData contains data relating to a friend
type RelationshipData struct {
	nex.Structure
	PID            uint32
	StrName        string
	ByRelationship uint8
	UIDetails      uint32
	ByStatus       uint8
}

// ExtractFromStream extracts a RelationshipData structure from a stream
func (relationshipData *RelationshipData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	relationshipData.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.PID. %s", err.Error())
	}

	relationshipData.StrName, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.StrName. %s", err.Error())
	}

	relationshipData.ByRelationship, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.ByRelationship. %s", err.Error())
	}

	relationshipData.UIDetails, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.UIDetails. %s", err.Error())
	}

	relationshipData.ByStatus, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.ByStatus. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RelationshipData
func (relationshipData *RelationshipData) Copy() nex.StructureInterface {
	copied := NewRelationshipData()

	copied.PID = relationshipData.PID
	copied.StrName = relationshipData.StrName
	copied.ByRelationship = relationshipData.ByRelationship
	copied.UIDetails = relationshipData.UIDetails
	copied.ByStatus = relationshipData.ByStatus

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (relationshipData *RelationshipData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*RelationshipData)

	if relationshipData.PID != other.PID {
		return false
	}

	if relationshipData.StrName != other.StrName {
		return false
	}

	if relationshipData.ByRelationship != other.ByRelationship {
		return false
	}

	if relationshipData.UIDetails != other.UIDetails {
		return false
	}

	if relationshipData.ByStatus != other.ByStatus {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (relationshipData *RelationshipData) String() string {
	return relationshipData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (relationshipData *RelationshipData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RelationshipData{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, relationshipData.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, relationshipData.PID))
	b.WriteString(fmt.Sprintf("%sStrName: %q,\n", indentationValues, relationshipData.StrName))
	b.WriteString(fmt.Sprintf("%sByRelationship: %d,\n", indentationValues, relationshipData.ByRelationship))
	b.WriteString(fmt.Sprintf("%sUIDetails: %d,\n", indentationValues, relationshipData.UIDetails))
	b.WriteString(fmt.Sprintf("%sByStatus: %q\n", indentationValues, relationshipData.ByStatus))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRelationshipData returns a new RelationshipData
func NewRelationshipData() *RelationshipData {
	return &RelationshipData{}
}
