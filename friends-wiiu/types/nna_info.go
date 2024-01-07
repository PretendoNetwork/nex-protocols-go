// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// NNAInfo contains information about a Nintendo Network Account
type NNAInfo struct {
	types.Structure
	*types.Data
	PrincipalBasicInfo *PrincipalBasicInfo
	Unknown1           *types.PrimitiveU8
	Unknown2           *types.PrimitiveU8
}

// WriteTo writes the NNAInfo to the given writable
func (nnaInfo *NNAInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nnaInfo.PrincipalBasicInfo.WriteTo(contentWritable)
	nnaInfo.Unknown1.WriteTo(contentWritable)
	nnaInfo.Unknown2.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nnaInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NNAInfo from the given readable
func (nnaInfo *NNAInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = nnaInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read NNAInfo header. %s", err.Error())
	}

	err = nnaInfo.PrincipalBasicInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.PrincipalBasicInfo. %s", err.Error())
	}

	err = nnaInfo.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Unknown1. %s", err.Error())
	}

	err = nnaInfo.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Unknown2. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NNAInfo
func (nnaInfo *NNAInfo) Copy() types.RVType {
	copied := NewNNAInfo()

	copied.StructureVersion = nnaInfo.StructureVersion

	copied.Data = nnaInfo.Data.Copy().(*types.Data)

	copied.PrincipalBasicInfo = nnaInfo.PrincipalBasicInfo.Copy().(*PrincipalBasicInfo)
	copied.Unknown1 = nnaInfo.Unknown1
	copied.Unknown2 = nnaInfo.Unknown2

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nnaInfo *NNAInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*NNAInfo); !ok {
		return false
	}

	other := o.(*NNAInfo)

	if nnaInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !nnaInfo.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !nnaInfo.PrincipalBasicInfo.Equals(other.PrincipalBasicInfo) {
		return false
	}

	if !nnaInfo.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !nnaInfo.Unknown2.Equals(other.Unknown2) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, nnaInfo.StructureVersion))

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
