// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// NNAInfo contains information about a Nintendo Network Account
type NNAInfo struct {
	nex.Structure
	*nex.Data
	PrincipalBasicInfo *PrincipalBasicInfo
	Unknown1           uint8
	Unknown2           uint8
}

// Bytes encodes the NNAInfo and returns a byte array
func (nnaInfo *NNAInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(nnaInfo.PrincipalBasicInfo)
	stream.WriteUInt8(nnaInfo.Unknown1)
	stream.WriteUInt8(nnaInfo.Unknown2)

	return stream.Bytes()
}

// ExtractFromStream extracts a NNAInfo structure from a stream
func (nnaInfo *NNAInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nnaInfo.PrincipalBasicInfo, err = nex.StreamReadStructure(stream, NewPrincipalBasicInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.PrincipalBasicInfo. %s", err.Error())
	}

	nnaInfo.Unknown1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Unknown1. %s", err.Error())
	}

	nnaInfo.Unknown2, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Unknown2. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NNAInfo
func (nnaInfo *NNAInfo) Copy() nex.StructureInterface {
	copied := NewNNAInfo()

	copied.SetStructureVersion(nnaInfo.StructureVersion())

	if nnaInfo.ParentType() != nil {
		copied.Data = nnaInfo.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.SetParentType(copied.Data)

	copied.PrincipalBasicInfo = nnaInfo.PrincipalBasicInfo.Copy().(*PrincipalBasicInfo)
	copied.Unknown1 = nnaInfo.Unknown1
	copied.Unknown2 = nnaInfo.Unknown2

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nnaInfo *NNAInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NNAInfo)

	if nnaInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !nnaInfo.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !nnaInfo.PrincipalBasicInfo.Equals(other.PrincipalBasicInfo) {
		return false
	}

	if nnaInfo.Unknown1 != other.Unknown1 {
		return false
	}

	if nnaInfo.Unknown2 != other.Unknown2 {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (nnaInfo *NNAInfo) String() string {
	return nnaInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (nnaInfo *NNAInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NNAInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, nnaInfo.StructureVersion()))

	if nnaInfo.PrincipalBasicInfo != nil {
		b.WriteString(fmt.Sprintf("%sPrincipalBasicInfo: %s,\n", indentationValues, nnaInfo.PrincipalBasicInfo.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrincipalBasicInfo: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sUnknown1: %d,\n", indentationValues, nnaInfo.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d\n", indentationValues, nnaInfo.Unknown2))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNNAInfo returns a new NNAInfo
func NewNNAInfo() *NNAInfo {
	return &NNAInfo{}
}
