// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemPostRightBinaryByAccountParam is a type within the ServiceItem protocol
type ServiceItemPostRightBinaryByAccountParam struct {
	types.Structure
	ReferenceID types.String
	UseType     types.UInt8
	RightBinary types.QBuffer
	LogMessage  types.String
	UniqueID    types.UInt32
	Platform    types.UInt8
}

// WriteTo writes the ServiceItemPostRightBinaryByAccountParam to the given writable
func (siprbbap ServiceItemPostRightBinaryByAccountParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siprbbap.ReferenceID.WriteTo(contentWritable)
	siprbbap.UseType.WriteTo(contentWritable)
	siprbbap.RightBinary.WriteTo(contentWritable)
	siprbbap.LogMessage.WriteTo(contentWritable)
	siprbbap.UniqueID.WriteTo(contentWritable)
	siprbbap.Platform.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siprbbap.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemPostRightBinaryByAccountParam from the given readable
func (siprbbap *ServiceItemPostRightBinaryByAccountParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = siprbbap.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam header. %s", err.Error())
	}

	err = siprbbap.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.ReferenceID. %s", err.Error())
	}

	err = siprbbap.UseType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.UseType. %s", err.Error())
	}

	err = siprbbap.RightBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.RightBinary. %s", err.Error())
	}

	err = siprbbap.LogMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.LogMessage. %s", err.Error())
	}

	err = siprbbap.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.UniqueID. %s", err.Error())
	}

	err = siprbbap.Platform.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.Platform. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemPostRightBinaryByAccountParam
func (siprbbap ServiceItemPostRightBinaryByAccountParam) Copy() types.RVType {
	copied := NewServiceItemPostRightBinaryByAccountParam()

	copied.StructureVersion = siprbbap.StructureVersion
	copied.ReferenceID = siprbbap.ReferenceID.Copy().(types.String)
	copied.UseType = siprbbap.UseType.Copy().(types.UInt8)
	copied.RightBinary = siprbbap.RightBinary.Copy().(types.QBuffer)
	copied.LogMessage = siprbbap.LogMessage.Copy().(types.String)
	copied.UniqueID = siprbbap.UniqueID.Copy().(types.UInt32)
	copied.Platform = siprbbap.Platform.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given ServiceItemPostRightBinaryByAccountParam contains the same data as the current ServiceItemPostRightBinaryByAccountParam
func (siprbbap ServiceItemPostRightBinaryByAccountParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPostRightBinaryByAccountParam); !ok {
		return false
	}

	other := o.(*ServiceItemPostRightBinaryByAccountParam)

	if siprbbap.StructureVersion != other.StructureVersion {
		return false
	}

	if !siprbbap.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !siprbbap.UseType.Equals(other.UseType) {
		return false
	}

	if !siprbbap.RightBinary.Equals(other.RightBinary) {
		return false
	}

	if !siprbbap.LogMessage.Equals(other.LogMessage) {
		return false
	}

	if !siprbbap.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return siprbbap.Platform.Equals(other.Platform)
}

// CopyRef copies the current value of the ServiceItemPostRightBinaryByAccountParam
// and returns a pointer to the new copy
func (siprbbap ServiceItemPostRightBinaryByAccountParam) CopyRef() types.RVTypePtr {
	copied := siprbbap.Copy().(ServiceItemPostRightBinaryByAccountParam)
	return &copied
}

// Deref takes a pointer to the ServiceItemPostRightBinaryByAccountParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (siprbbap *ServiceItemPostRightBinaryByAccountParam) Deref() types.RVType {
	return *siprbbap
}

// String returns the string representation of the ServiceItemPostRightBinaryByAccountParam
func (siprbbap ServiceItemPostRightBinaryByAccountParam) String() string {
	return siprbbap.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemPostRightBinaryByAccountParam using the provided indentation level
func (siprbbap ServiceItemPostRightBinaryByAccountParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPostRightBinaryByAccountParam{\n")
	b.WriteString(fmt.Sprintf("%sReferenceID: %s,\n", indentationValues, siprbbap.ReferenceID))
	b.WriteString(fmt.Sprintf("%sUseType: %s,\n", indentationValues, siprbbap.UseType))
	b.WriteString(fmt.Sprintf("%sRightBinary: %s,\n", indentationValues, siprbbap.RightBinary))
	b.WriteString(fmt.Sprintf("%sLogMessage: %s,\n", indentationValues, siprbbap.LogMessage))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, siprbbap.UniqueID))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, siprbbap.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPostRightBinaryByAccountParam returns a new ServiceItemPostRightBinaryByAccountParam
func NewServiceItemPostRightBinaryByAccountParam() ServiceItemPostRightBinaryByAccountParam {
	return ServiceItemPostRightBinaryByAccountParam{
		ReferenceID: types.NewString(""),
		UseType:     types.NewUInt8(0),
		RightBinary: types.NewQBuffer(nil),
		LogMessage:  types.NewString(""),
		UniqueID:    types.NewUInt32(0),
		Platform:    types.NewUInt8(0),
	}

}
