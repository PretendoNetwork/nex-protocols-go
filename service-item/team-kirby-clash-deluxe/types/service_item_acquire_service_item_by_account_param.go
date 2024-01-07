// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemAcquireServiceItemByAccountParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAcquireServiceItemByAccountParam struct {
	types.Structure
	ReferenceIDForAcquisition string
	ReferenceIDForRightBinary string
	UseType                   *types.PrimitiveU8
	LimitationType            *types.PrimitiveU32
	LimitationValue           *types.PrimitiveU32
	RightBinary               []byte
	LogMessage                string
	UniqueID                  *types.PrimitiveU32
	Platform                  *types.PrimitiveU8
}

// ExtractFrom extracts the ServiceItemAcquireServiceItemByAccountParam from the given readable
func (serviceItemAcquireServiceItemByAccountParam *ServiceItemAcquireServiceItemByAccountParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemAcquireServiceItemByAccountParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemAcquireServiceItemByAccountParam header. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemByAccountParam.UseType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.UseType from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemByAccountParam.LimitationType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.LimitationType from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemByAccountParam.LimitationValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.LimitationValue from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemByAccountParam.RightBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.RightBinary from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemByAccountParam.LogMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.LogMessage from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemByAccountParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemAcquireServiceItemByAccountParam.StructureVersion >= 1 {
	err = 	serviceItemAcquireServiceItemByAccountParam.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the ServiceItemAcquireServiceItemByAccountParam to the given writable
func (serviceItemAcquireServiceItemByAccountParam *ServiceItemAcquireServiceItemByAccountParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition.WriteTo(contentWritable)
	serviceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary.WriteTo(contentWritable)
	serviceItemAcquireServiceItemByAccountParam.UseType.WriteTo(contentWritable)
	serviceItemAcquireServiceItemByAccountParam.LimitationType.WriteTo(contentWritable)
	serviceItemAcquireServiceItemByAccountParam.LimitationValue.WriteTo(contentWritable)
	stream.WriteQBuffer(serviceItemAcquireServiceItemByAccountParam.RightBinary)
	serviceItemAcquireServiceItemByAccountParam.LogMessage.WriteTo(contentWritable)
	serviceItemAcquireServiceItemByAccountParam.UniqueID.WriteTo(contentWritable)

	if serviceItemAcquireServiceItemByAccountParam.StructureVersion >= 1 {
		serviceItemAcquireServiceItemByAccountParam.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemAcquireServiceItemByAccountParam
func (serviceItemAcquireServiceItemByAccountParam *ServiceItemAcquireServiceItemByAccountParam) Copy() types.RVType {
	copied := NewServiceItemAcquireServiceItemByAccountParam()

	copied.StructureVersion = serviceItemAcquireServiceItemByAccountParam.StructureVersion

	copied.ReferenceIDForAcquisition = serviceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition
	copied.ReferenceIDForRightBinary = serviceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary
	copied.UseType = serviceItemAcquireServiceItemByAccountParam.UseType
	copied.LimitationType = serviceItemAcquireServiceItemByAccountParam.LimitationType
	copied.LimitationValue = serviceItemAcquireServiceItemByAccountParam.LimitationValue
	copied.RightBinary = serviceItemAcquireServiceItemByAccountParam.RightBinary
	copied.LogMessage = serviceItemAcquireServiceItemByAccountParam.LogMessage
	copied.UniqueID = serviceItemAcquireServiceItemByAccountParam.UniqueID
	copied.Platform = serviceItemAcquireServiceItemByAccountParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAcquireServiceItemByAccountParam *ServiceItemAcquireServiceItemByAccountParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAcquireServiceItemByAccountParam); !ok {
		return false
	}

	other := o.(*ServiceItemAcquireServiceItemByAccountParam)

	if serviceItemAcquireServiceItemByAccountParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition.Equals(other.ReferenceIDForAcquisition) {
		return false
	}

	if !serviceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary.Equals(other.ReferenceIDForRightBinary) {
		return false
	}

	if !serviceItemAcquireServiceItemByAccountParam.UseType.Equals(other.UseType) {
		return false
	}

	if !serviceItemAcquireServiceItemByAccountParam.LimitationType.Equals(other.LimitationType) {
		return false
	}

	if !serviceItemAcquireServiceItemByAccountParam.LimitationValue.Equals(other.LimitationValue) {
		return false
	}

	if !serviceItemAcquireServiceItemByAccountParam.RightBinary.Equals(other.RightBinary) {
		return false
	}

	if !serviceItemAcquireServiceItemByAccountParam.LogMessage.Equals(other.LogMessage) {
		return false
	}

	if !serviceItemAcquireServiceItemByAccountParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !serviceItemAcquireServiceItemByAccountParam.Platform.Equals(other.Platform) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemAcquireServiceItemByAccountParam *ServiceItemAcquireServiceItemByAccountParam) String() string {
	return serviceItemAcquireServiceItemByAccountParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemAcquireServiceItemByAccountParam *ServiceItemAcquireServiceItemByAccountParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAcquireServiceItemByAccountParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReferenceIDForAcquisition: %q,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition))
	b.WriteString(fmt.Sprintf("%sReferenceIDForRightBinary: %q,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary))
	b.WriteString(fmt.Sprintf("%sUseType: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.UseType))
	b.WriteString(fmt.Sprintf("%sLimitationType: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.LimitationType))
	b.WriteString(fmt.Sprintf("%sLimitationValue: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.LimitationValue))
	b.WriteString(fmt.Sprintf("%sRightBinary: %x,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.RightBinary))
	b.WriteString(fmt.Sprintf("%sLogMessage: %q,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.LogMessage))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.UniqueID))

	if serviceItemAcquireServiceItemByAccountParam.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAcquireServiceItemByAccountParam returns a new ServiceItemAcquireServiceItemByAccountParam
func NewServiceItemAcquireServiceItemByAccountParam() *ServiceItemAcquireServiceItemByAccountParam {
	return &ServiceItemAcquireServiceItemByAccountParam{}
}
