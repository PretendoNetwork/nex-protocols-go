// Package types implements all the types used by the Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// BasicAccountInfo contains data for creating a new NNID on the network
type BasicAccountInfo struct {
	nex.Structure
	PIDOwner uint32
	StrName  string
}

// ExtractFromStream extracts a BasicAccountInfo structure from a stream
func (basicAccountInfo *BasicAccountInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	basicAccountInfo.PIDOwner, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BasicAccountInfo.PIDOwner. %s", err.Error())
	}

	basicAccountInfo.StrName, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract BasicAccountInfo.StrName. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BasicAccountInfo
func (basicAccountInfo *BasicAccountInfo) Copy() nex.StructureInterface {
	copied := NewBasicAccountInfo()

	copied.SetStructureVersion(basicAccountInfo.StructureVersion())

	copied.PIDOwner = basicAccountInfo.PIDOwner
	copied.StrName = basicAccountInfo.StrName

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (basicAccountInfo *BasicAccountInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*BasicAccountInfo)

	if basicAccountInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if basicAccountInfo.PIDOwner != other.PIDOwner {
		return false
	}

	if basicAccountInfo.StrName != other.StrName {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (basicAccountInfo *BasicAccountInfo) String() string {
	return basicAccountInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (basicAccountInfo *BasicAccountInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BasicAccountInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, basicAccountInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPIDOwner: %d,\n", indentationValues, basicAccountInfo.PIDOwner))
	b.WriteString(fmt.Sprintf("%sStrName: %q\n", indentationValues, basicAccountInfo.StrName))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBasicAccountInfo returns a new BasicAccountInfo
func NewBasicAccountInfo() *BasicAccountInfo {
	return &BasicAccountInfo{}
}
