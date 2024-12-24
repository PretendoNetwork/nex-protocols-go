// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemAcquireServiceItemByAccountParam is a type within the ServiceItem protocol
type ServiceItemAcquireServiceItemByAccountParam struct {
	types.Structure
	ReferenceIDForAcquisition types.String
	ReferenceIDForRightBinary types.String
	UseType                   types.UInt8
	LimitationType            types.UInt32
	LimitationValue           types.UInt32
	RightBinary               types.QBuffer
	LogMessage                types.String
	UniqueID                  types.UInt32
	Platform                  types.UInt8
}

// WriteTo writes the ServiceItemAcquireServiceItemByAccountParam to the given writable
func (siasibap ServiceItemAcquireServiceItemByAccountParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siasibap.ReferenceIDForAcquisition.WriteTo(contentWritable)
	siasibap.ReferenceIDForRightBinary.WriteTo(contentWritable)
	siasibap.UseType.WriteTo(contentWritable)
	siasibap.LimitationType.WriteTo(contentWritable)
	siasibap.LimitationValue.WriteTo(contentWritable)
	siasibap.RightBinary.WriteTo(contentWritable)
	siasibap.LogMessage.WriteTo(contentWritable)
	siasibap.UniqueID.WriteTo(contentWritable)
	siasibap.Platform.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siasibap.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemAcquireServiceItemByAccountParam from the given readable
func (siasibap *ServiceItemAcquireServiceItemByAccountParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = siasibap.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam header. %s", err.Error())
	}

	err = siasibap.ReferenceIDForAcquisition.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition. %s", err.Error())
	}

	err = siasibap.ReferenceIDForRightBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary. %s", err.Error())
	}

	err = siasibap.UseType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.UseType. %s", err.Error())
	}

	err = siasibap.LimitationType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.LimitationType. %s", err.Error())
	}

	err = siasibap.LimitationValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.LimitationValue. %s", err.Error())
	}

	err = siasibap.RightBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.RightBinary. %s", err.Error())
	}

	err = siasibap.LogMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.LogMessage. %s", err.Error())
	}

	err = siasibap.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.UniqueID. %s", err.Error())
	}

	err = siasibap.Platform.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.Platform. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemAcquireServiceItemByAccountParam
func (siasibap ServiceItemAcquireServiceItemByAccountParam) Copy() types.RVType {
	copied := NewServiceItemAcquireServiceItemByAccountParam()

	copied.StructureVersion = siasibap.StructureVersion
	copied.ReferenceIDForAcquisition = siasibap.ReferenceIDForAcquisition.Copy().(types.String)
	copied.ReferenceIDForRightBinary = siasibap.ReferenceIDForRightBinary.Copy().(types.String)
	copied.UseType = siasibap.UseType.Copy().(types.UInt8)
	copied.LimitationType = siasibap.LimitationType.Copy().(types.UInt32)
	copied.LimitationValue = siasibap.LimitationValue.Copy().(types.UInt32)
	copied.RightBinary = siasibap.RightBinary.Copy().(types.QBuffer)
	copied.LogMessage = siasibap.LogMessage.Copy().(types.String)
	copied.UniqueID = siasibap.UniqueID.Copy().(types.UInt32)
	copied.Platform = siasibap.Platform.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given ServiceItemAcquireServiceItemByAccountParam contains the same data as the current ServiceItemAcquireServiceItemByAccountParam
func (siasibap ServiceItemAcquireServiceItemByAccountParam) Equals(o types.RVType) bool {
	if _, ok := o.(ServiceItemAcquireServiceItemByAccountParam); !ok {
		return false
	}

	other := o.(ServiceItemAcquireServiceItemByAccountParam)

	if siasibap.StructureVersion != other.StructureVersion {
		return false
	}

	if !siasibap.ReferenceIDForAcquisition.Equals(other.ReferenceIDForAcquisition) {
		return false
	}

	if !siasibap.ReferenceIDForRightBinary.Equals(other.ReferenceIDForRightBinary) {
		return false
	}

	if !siasibap.UseType.Equals(other.UseType) {
		return false
	}

	if !siasibap.LimitationType.Equals(other.LimitationType) {
		return false
	}

	if !siasibap.LimitationValue.Equals(other.LimitationValue) {
		return false
	}

	if !siasibap.RightBinary.Equals(other.RightBinary) {
		return false
	}

	if !siasibap.LogMessage.Equals(other.LogMessage) {
		return false
	}

	if !siasibap.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return siasibap.Platform.Equals(other.Platform)
}

// CopyRef copies the current value of the ServiceItemAcquireServiceItemByAccountParam
// and returns a pointer to the new copy
func (siasibap ServiceItemAcquireServiceItemByAccountParam) CopyRef() types.RVTypePtr {
	copied := siasibap.Copy().(ServiceItemAcquireServiceItemByAccountParam)
	return &copied
}

// Deref takes a pointer to the ServiceItemAcquireServiceItemByAccountParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (siasibap *ServiceItemAcquireServiceItemByAccountParam) Deref() types.RVType {
	return *siasibap
}

// String returns the string representation of the ServiceItemAcquireServiceItemByAccountParam
func (siasibap ServiceItemAcquireServiceItemByAccountParam) String() string {
	return siasibap.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemAcquireServiceItemByAccountParam using the provided indentation level
func (siasibap ServiceItemAcquireServiceItemByAccountParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAcquireServiceItemByAccountParam{\n")
	b.WriteString(fmt.Sprintf("%sReferenceIDForAcquisition: %s,\n", indentationValues, siasibap.ReferenceIDForAcquisition))
	b.WriteString(fmt.Sprintf("%sReferenceIDForRightBinary: %s,\n", indentationValues, siasibap.ReferenceIDForRightBinary))
	b.WriteString(fmt.Sprintf("%sUseType: %s,\n", indentationValues, siasibap.UseType))
	b.WriteString(fmt.Sprintf("%sLimitationType: %s,\n", indentationValues, siasibap.LimitationType))
	b.WriteString(fmt.Sprintf("%sLimitationValue: %s,\n", indentationValues, siasibap.LimitationValue))
	b.WriteString(fmt.Sprintf("%sRightBinary: %s,\n", indentationValues, siasibap.RightBinary))
	b.WriteString(fmt.Sprintf("%sLogMessage: %s,\n", indentationValues, siasibap.LogMessage))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, siasibap.UniqueID))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, siasibap.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAcquireServiceItemByAccountParam returns a new ServiceItemAcquireServiceItemByAccountParam
func NewServiceItemAcquireServiceItemByAccountParam() ServiceItemAcquireServiceItemByAccountParam {
	return ServiceItemAcquireServiceItemByAccountParam{
		ReferenceIDForAcquisition: types.NewString(""),
		ReferenceIDForRightBinary: types.NewString(""),
		UseType:                   types.NewUInt8(0),
		LimitationType:            types.NewUInt32(0),
		LimitationValue:           types.NewUInt32(0),
		RightBinary:               types.NewQBuffer(nil),
		LogMessage:                types.NewString(""),
		UniqueID:                  types.NewUInt32(0),
		Platform:                  types.NewUInt8(0),
	}

}
