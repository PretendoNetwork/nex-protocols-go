// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemPostRightBinaryByAccountParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemPostRightBinaryByAccountParam struct {
	types.Structure
	ReferenceID string
	UseType     *types.PrimitiveU8
	RightBinary []byte
	LogMessage  string
	UniqueID    *types.PrimitiveU32
	Platform    *types.PrimitiveU8
}

// ExtractFrom extracts the ServiceItemPostRightBinaryByAccountParam from the given readable
func (serviceItemPostRightBinaryByAccountParam *ServiceItemPostRightBinaryByAccountParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemPostRightBinaryByAccountParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemPostRightBinaryByAccountParam header. %s", err.Error())
	}

	err = serviceItemPostRightBinaryByAccountParam.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.ReferenceID from stream. %s", err.Error())
	}

	err = serviceItemPostRightBinaryByAccountParam.UseType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.UseType from stream. %s", err.Error())
	}

	serviceItemPostRightBinaryByAccountParam.RightBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.RightBinary from stream. %s", err.Error())
	}

	err = serviceItemPostRightBinaryByAccountParam.LogMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.LogMessage from stream. %s", err.Error())
	}

	err = serviceItemPostRightBinaryByAccountParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemPostRightBinaryByAccountParam.StructureVersion >= 1 {
	err = 	serviceItemPostRightBinaryByAccountParam.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the ServiceItemPostRightBinaryByAccountParam to the given writable
func (serviceItemPostRightBinaryByAccountParam *ServiceItemPostRightBinaryByAccountParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemPostRightBinaryByAccountParam.ReferenceID.WriteTo(contentWritable)
	serviceItemPostRightBinaryByAccountParam.UseType.WriteTo(contentWritable)
	stream.WriteQBuffer(serviceItemPostRightBinaryByAccountParam.RightBinary)
	serviceItemPostRightBinaryByAccountParam.LogMessage.WriteTo(contentWritable)
	serviceItemPostRightBinaryByAccountParam.UniqueID.WriteTo(contentWritable)

	if serviceItemPostRightBinaryByAccountParam.StructureVersion >= 1 {
		serviceItemPostRightBinaryByAccountParam.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemPostRightBinaryByAccountParam
func (serviceItemPostRightBinaryByAccountParam *ServiceItemPostRightBinaryByAccountParam) Copy() types.RVType {
	copied := NewServiceItemPostRightBinaryByAccountParam()

	copied.StructureVersion = serviceItemPostRightBinaryByAccountParam.StructureVersion

	copied.ReferenceID = serviceItemPostRightBinaryByAccountParam.ReferenceID
	copied.UseType = serviceItemPostRightBinaryByAccountParam.UseType
	copied.RightBinary = serviceItemPostRightBinaryByAccountParam.RightBinary
	copied.LogMessage = serviceItemPostRightBinaryByAccountParam.LogMessage
	copied.UniqueID = serviceItemPostRightBinaryByAccountParam.UniqueID
	copied.Platform = serviceItemPostRightBinaryByAccountParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPostRightBinaryByAccountParam *ServiceItemPostRightBinaryByAccountParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPostRightBinaryByAccountParam); !ok {
		return false
	}

	other := o.(*ServiceItemPostRightBinaryByAccountParam)

	if serviceItemPostRightBinaryByAccountParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemPostRightBinaryByAccountParam.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !serviceItemPostRightBinaryByAccountParam.UseType.Equals(other.UseType) {
		return false
	}

	if !serviceItemPostRightBinaryByAccountParam.RightBinary.Equals(other.RightBinary) {
		return false
	}

	if !serviceItemPostRightBinaryByAccountParam.LogMessage.Equals(other.LogMessage) {
		return false
	}

	if !serviceItemPostRightBinaryByAccountParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !serviceItemPostRightBinaryByAccountParam.Platform.Equals(other.Platform) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemPostRightBinaryByAccountParam *ServiceItemPostRightBinaryByAccountParam) String() string {
	return serviceItemPostRightBinaryByAccountParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemPostRightBinaryByAccountParam *ServiceItemPostRightBinaryByAccountParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPostRightBinaryByAccountParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.ReferenceID))
	b.WriteString(fmt.Sprintf("%sUseType: %d,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.UseType))
	b.WriteString(fmt.Sprintf("%sRightBinary: %x,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.RightBinary))
	b.WriteString(fmt.Sprintf("%sLogMessage: %q,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.LogMessage))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.UniqueID))

	if serviceItemPostRightBinaryByAccountParam.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPostRightBinaryByAccountParam returns a new ServiceItemPostRightBinaryByAccountParam
func NewServiceItemPostRightBinaryByAccountParam() *ServiceItemPostRightBinaryByAccountParam {
	return &ServiceItemPostRightBinaryByAccountParam{}
}
