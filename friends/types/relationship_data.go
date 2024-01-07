// Package types implements all the types used by the Friends protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// RelationshipData contains data relating to a friend
type RelationshipData struct {
	types.Structure
	PID            *types.PrimitiveU32
	StrName        string
	ByRelationship *types.PrimitiveU8
	UIDetails      *types.PrimitiveU32
	ByStatus       *types.PrimitiveU8
}

// ExtractFrom extracts the RelationshipData from the given readable
func (relationshipData *RelationshipData) ExtractFrom(readable types.Readable) error {
	var err error

	if err = relationshipData.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read RelationshipData header. %s", err.Error())
	}

	err = relationshipData.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.PID. %s", err.Error())
	}

	err = relationshipData.StrName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.StrName. %s", err.Error())
	}

	err = relationshipData.ByRelationship.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.ByRelationship. %s", err.Error())
	}

	err = relationshipData.UIDetails.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.UIDetails. %s", err.Error())
	}

	err = relationshipData.ByStatus.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.ByStatus. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RelationshipData
func (relationshipData *RelationshipData) Copy() types.RVType {
	copied := NewRelationshipData()

	copied.StructureVersion = relationshipData.StructureVersion

	copied.PID = relationshipData.PID
	copied.StrName = relationshipData.StrName
	copied.ByRelationship = relationshipData.ByRelationship
	copied.UIDetails = relationshipData.UIDetails
	copied.ByStatus = relationshipData.ByStatus

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (relationshipData *RelationshipData) Equals(o types.RVType) bool {
	if _, ok := o.(*RelationshipData); !ok {
		return false
	}

	other := o.(*RelationshipData)

	if relationshipData.StructureVersion != other.StructureVersion {
		return false
	}

	if !relationshipData.PID.Equals(other.PID) {
		return false
	}

	if !relationshipData.StrName.Equals(other.StrName) {
		return false
	}

	if !relationshipData.ByRelationship.Equals(other.ByRelationship) {
		return false
	}

	if !relationshipData.UIDetails.Equals(other.UIDetails) {
		return false
	}

	if !relationshipData.ByStatus.Equals(other.ByStatus) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, relationshipData.StructureVersion))
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
