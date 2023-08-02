// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemUseServiceItemByAccountParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemUseServiceItemByAccountParam struct {
	nex.Structure
	ReferenceIDForUse string
	ReferenceIDForRightBinary string
	UseType uint8
	UseNumber uint8
	RightBinary []byte
	LogMessage string
	UniqueID uint32
}

// ExtractFromStream extracts a ServiceItemUseServiceItemByAccountParam structure from a stream
func (serviceItemUseServiceItemByAccountParam *ServiceItemUseServiceItemByAccountParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemUseServiceItemByAccountParam.ReferenceIDForUse, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.ReferenceIDForUse from stream. %s", err.Error())
	}

	serviceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary from stream. %s", err.Error())
	}

	serviceItemUseServiceItemByAccountParam.UseType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.UseType from stream. %s", err.Error())
	}

	serviceItemUseServiceItemByAccountParam.UseNumber, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.UseNumber from stream. %s", err.Error())
	}

	serviceItemUseServiceItemByAccountParam.RightBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.RightBinary from stream. %s", err.Error())
	}

	serviceItemUseServiceItemByAccountParam.LogMessage, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.LogMessage from stream. %s", err.Error())
	}

	serviceItemUseServiceItemByAccountParam.UniqueID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.UniqueID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemUseServiceItemByAccountParam and returns a byte array
func (serviceItemUseServiceItemByAccountParam *ServiceItemUseServiceItemByAccountParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemUseServiceItemByAccountParam.ReferenceIDForUse)
	stream.WriteString(serviceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary)
	stream.WriteUInt8(serviceItemUseServiceItemByAccountParam.UseType)
	stream.WriteUInt8(serviceItemUseServiceItemByAccountParam.UseNumber)
	stream.WriteQBuffer(serviceItemUseServiceItemByAccountParam.RightBinary)
	stream.WriteString(serviceItemUseServiceItemByAccountParam.LogMessage)
	stream.WriteUInt32LE(serviceItemUseServiceItemByAccountParam.UniqueID)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemUseServiceItemByAccountParam
func (serviceItemUseServiceItemByAccountParam *ServiceItemUseServiceItemByAccountParam) Copy() nex.StructureInterface {
	copied := NewServiceItemUseServiceItemByAccountParam()

	copied.ReferenceIDForUse = serviceItemUseServiceItemByAccountParam.ReferenceIDForUse
	copied.ReferenceIDForRightBinary = serviceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary
	copied.UseType = serviceItemUseServiceItemByAccountParam.UseType
	copied.UseNumber = serviceItemUseServiceItemByAccountParam.UseNumber
	copied.RightBinary = serviceItemUseServiceItemByAccountParam.RightBinary
	copied.LogMessage = serviceItemUseServiceItemByAccountParam.LogMessage
	copied.UniqueID = serviceItemUseServiceItemByAccountParam.UniqueID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemUseServiceItemByAccountParam *ServiceItemUseServiceItemByAccountParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemUseServiceItemByAccountParam)

	if serviceItemUseServiceItemByAccountParam.ReferenceIDForUse != other.ReferenceIDForUse {
		return false
	}

	if serviceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary != other.ReferenceIDForRightBinary {
		return false
	}

	if serviceItemUseServiceItemByAccountParam.UseType != other.UseType {
		return false
	}

	if serviceItemUseServiceItemByAccountParam.UseNumber != other.UseNumber {
		return false
	}

	if !bytes.Equal(serviceItemUseServiceItemByAccountParam.RightBinary, other.RightBinary) {
		return false
	}

	if serviceItemUseServiceItemByAccountParam.LogMessage != other.LogMessage {
		return false
	}

	if serviceItemUseServiceItemByAccountParam.UniqueID != other.UniqueID {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemUseServiceItemByAccountParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReferenceIDForUse: %q,\n", indentationValues, serviceItemUseServiceItemByAccountParam.ReferenceIDForUse))
	b.WriteString(fmt.Sprintf("%sReferenceIDForRightBinary: %q,\n", indentationValues, serviceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary))
	b.WriteString(fmt.Sprintf("%sUseType: %d,\n", indentationValues, serviceItemUseServiceItemByAccountParam.UseType))
	b.WriteString(fmt.Sprintf("%sUseNumber: %d,\n", indentationValues, serviceItemUseServiceItemByAccountParam.UseNumber))
	b.WriteString(fmt.Sprintf("%sRightBinary: %x,\n", indentationValues, serviceItemUseServiceItemByAccountParam.RightBinary))
	b.WriteString(fmt.Sprintf("%sLogMessage: %q,\n", indentationValues, serviceItemUseServiceItemByAccountParam.LogMessage))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemUseServiceItemByAccountParam.UniqueID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemUseServiceItemByAccountParam returns a new ServiceItemUseServiceItemByAccountParam
func NewServiceItemUseServiceItemByAccountParam() *ServiceItemUseServiceItemByAccountParam {
	return &ServiceItemUseServiceItemByAccountParam{}
}
