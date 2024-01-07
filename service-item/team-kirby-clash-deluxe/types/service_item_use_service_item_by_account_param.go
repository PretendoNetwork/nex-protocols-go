// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemUseServiceItemByAccountParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemUseServiceItemByAccountParam struct {
	types.Structure
	ReferenceIDForUse         string
	ReferenceIDForRightBinary string
	UseType                   *types.PrimitiveU8
	UseNumber                 *types.PrimitiveU8
	RightBinary               []byte
	LogMessage                string
	UniqueID                  *types.PrimitiveU32
	Platform                  *types.PrimitiveU8
}

// ExtractFrom extracts the ServiceItemUseServiceItemByAccountParam from the given readable
func (serviceItemUseServiceItemByAccountParam *ServiceItemUseServiceItemByAccountParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemUseServiceItemByAccountParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemUseServiceItemByAccountParam header. %s", err.Error())
	}

	err = serviceItemUseServiceItemByAccountParam.ReferenceIDForUse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.ReferenceIDForUse from stream. %s", err.Error())
	}

	err = serviceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary from stream. %s", err.Error())
	}

	err = serviceItemUseServiceItemByAccountParam.UseType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.UseType from stream. %s", err.Error())
	}

	err = serviceItemUseServiceItemByAccountParam.UseNumber.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.UseNumber from stream. %s", err.Error())
	}

	serviceItemUseServiceItemByAccountParam.RightBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.RightBinary from stream. %s", err.Error())
	}

	err = serviceItemUseServiceItemByAccountParam.LogMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.LogMessage from stream. %s", err.Error())
	}

	err = serviceItemUseServiceItemByAccountParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemUseServiceItemByAccountParam.StructureVersion >= 1 {
	err = 	serviceItemUseServiceItemByAccountParam.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the ServiceItemUseServiceItemByAccountParam to the given writable
func (serviceItemUseServiceItemByAccountParam *ServiceItemUseServiceItemByAccountParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemUseServiceItemByAccountParam.ReferenceIDForUse.WriteTo(contentWritable)
	serviceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary.WriteTo(contentWritable)
	serviceItemUseServiceItemByAccountParam.UseType.WriteTo(contentWritable)
	serviceItemUseServiceItemByAccountParam.UseNumber.WriteTo(contentWritable)
	stream.WriteQBuffer(serviceItemUseServiceItemByAccountParam.RightBinary)
	serviceItemUseServiceItemByAccountParam.LogMessage.WriteTo(contentWritable)
	serviceItemUseServiceItemByAccountParam.UniqueID.WriteTo(contentWritable)

	if serviceItemUseServiceItemByAccountParam.StructureVersion >= 1 {
		serviceItemUseServiceItemByAccountParam.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemUseServiceItemByAccountParam
func (serviceItemUseServiceItemByAccountParam *ServiceItemUseServiceItemByAccountParam) Copy() types.RVType {
	copied := NewServiceItemUseServiceItemByAccountParam()

	copied.StructureVersion = serviceItemUseServiceItemByAccountParam.StructureVersion

	copied.ReferenceIDForUse = serviceItemUseServiceItemByAccountParam.ReferenceIDForUse
	copied.ReferenceIDForRightBinary = serviceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary
	copied.UseType = serviceItemUseServiceItemByAccountParam.UseType
	copied.UseNumber = serviceItemUseServiceItemByAccountParam.UseNumber
	copied.RightBinary = serviceItemUseServiceItemByAccountParam.RightBinary
	copied.LogMessage = serviceItemUseServiceItemByAccountParam.LogMessage
	copied.UniqueID = serviceItemUseServiceItemByAccountParam.UniqueID
	copied.Platform = serviceItemUseServiceItemByAccountParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemUseServiceItemByAccountParam *ServiceItemUseServiceItemByAccountParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemUseServiceItemByAccountParam); !ok {
		return false
	}

	other := o.(*ServiceItemUseServiceItemByAccountParam)

	if serviceItemUseServiceItemByAccountParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemUseServiceItemByAccountParam.ReferenceIDForUse.Equals(other.ReferenceIDForUse) {
		return false
	}

	if !serviceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary.Equals(other.ReferenceIDForRightBinary) {
		return false
	}

	if !serviceItemUseServiceItemByAccountParam.UseType.Equals(other.UseType) {
		return false
	}

	if !serviceItemUseServiceItemByAccountParam.UseNumber.Equals(other.UseNumber) {
		return false
	}

	if !serviceItemUseServiceItemByAccountParam.RightBinary.Equals(other.RightBinary) {
		return false
	}

	if !serviceItemUseServiceItemByAccountParam.LogMessage.Equals(other.LogMessage) {
		return false
	}

	if !serviceItemUseServiceItemByAccountParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !serviceItemUseServiceItemByAccountParam.Platform.Equals(other.Platform) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemUseServiceItemByAccountParam *ServiceItemUseServiceItemByAccountParam) String() string {
	return serviceItemUseServiceItemByAccountParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemUseServiceItemByAccountParam *ServiceItemUseServiceItemByAccountParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemUseServiceItemByAccountParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemUseServiceItemByAccountParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReferenceIDForUse: %q,\n", indentationValues, serviceItemUseServiceItemByAccountParam.ReferenceIDForUse))
	b.WriteString(fmt.Sprintf("%sReferenceIDForRightBinary: %q,\n", indentationValues, serviceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary))
	b.WriteString(fmt.Sprintf("%sUseType: %d,\n", indentationValues, serviceItemUseServiceItemByAccountParam.UseType))
	b.WriteString(fmt.Sprintf("%sUseNumber: %d,\n", indentationValues, serviceItemUseServiceItemByAccountParam.UseNumber))
	b.WriteString(fmt.Sprintf("%sRightBinary: %x,\n", indentationValues, serviceItemUseServiceItemByAccountParam.RightBinary))
	b.WriteString(fmt.Sprintf("%sLogMessage: %q,\n", indentationValues, serviceItemUseServiceItemByAccountParam.LogMessage))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemUseServiceItemByAccountParam.UniqueID))

	if serviceItemUseServiceItemByAccountParam.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemUseServiceItemByAccountParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemUseServiceItemByAccountParam returns a new ServiceItemUseServiceItemByAccountParam
func NewServiceItemUseServiceItemByAccountParam() *ServiceItemUseServiceItemByAccountParam {
	return &ServiceItemUseServiceItemByAccountParam{}
}
