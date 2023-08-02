// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemAcquireServiceItemByAccountParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAcquireServiceItemByAccountParam struct {
	nex.Structure
	ReferenceIDForAcquisition string
	ReferenceIDForRightBinary string
	UseType uint8
	LimitationType uint32
	LimitationValue uint32
	RightBinary []byte
	LogMessage string
	UniqueID uint32
}

// ExtractFromStream extracts a ServiceItemAcquireServiceItemByAccountParam structure from a stream
func (serviceItemAcquireServiceItemByAccountParam *ServiceItemAcquireServiceItemByAccountParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemByAccountParam.UseType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.UseType from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemByAccountParam.LimitationType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.LimitationType from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemByAccountParam.LimitationValue, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.LimitationValue from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemByAccountParam.RightBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.RightBinary from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemByAccountParam.LogMessage, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.LogMessage from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemByAccountParam.UniqueID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemByAccountParam.UniqueID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemAcquireServiceItemByAccountParam and returns a byte array
func (serviceItemAcquireServiceItemByAccountParam *ServiceItemAcquireServiceItemByAccountParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition)
	stream.WriteString(serviceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary)
	stream.WriteUInt8(serviceItemAcquireServiceItemByAccountParam.UseType)
	stream.WriteUInt32LE(serviceItemAcquireServiceItemByAccountParam.LimitationType)
	stream.WriteUInt32LE(serviceItemAcquireServiceItemByAccountParam.LimitationValue)
	stream.WriteQBuffer(serviceItemAcquireServiceItemByAccountParam.RightBinary)
	stream.WriteString(serviceItemAcquireServiceItemByAccountParam.LogMessage)
	stream.WriteUInt32LE(serviceItemAcquireServiceItemByAccountParam.UniqueID)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemAcquireServiceItemByAccountParam
func (serviceItemAcquireServiceItemByAccountParam *ServiceItemAcquireServiceItemByAccountParam) Copy() nex.StructureInterface {
	copied := NewServiceItemAcquireServiceItemByAccountParam()

	copied.ReferenceIDForAcquisition = serviceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition
	copied.ReferenceIDForRightBinary = serviceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary
	copied.UseType = serviceItemAcquireServiceItemByAccountParam.UseType
	copied.LimitationType = serviceItemAcquireServiceItemByAccountParam.LimitationType
	copied.LimitationValue = serviceItemAcquireServiceItemByAccountParam.LimitationValue
	copied.RightBinary = serviceItemAcquireServiceItemByAccountParam.RightBinary
	copied.LogMessage = serviceItemAcquireServiceItemByAccountParam.LogMessage
	copied.UniqueID = serviceItemAcquireServiceItemByAccountParam.UniqueID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAcquireServiceItemByAccountParam *ServiceItemAcquireServiceItemByAccountParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemAcquireServiceItemByAccountParam)

	if serviceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition != other.ReferenceIDForAcquisition {
		return false
	}

	if serviceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary != other.ReferenceIDForRightBinary {
		return false
	}

	if serviceItemAcquireServiceItemByAccountParam.UseType != other.UseType {
		return false
	}

	if serviceItemAcquireServiceItemByAccountParam.LimitationType != other.LimitationType {
		return false
	}

	if serviceItemAcquireServiceItemByAccountParam.LimitationValue != other.LimitationValue {
		return false
	}

	if !bytes.Equal(serviceItemAcquireServiceItemByAccountParam.RightBinary, other.RightBinary) {
		return false
	}

	if serviceItemAcquireServiceItemByAccountParam.LogMessage != other.LogMessage {
		return false
	}

	if serviceItemAcquireServiceItemByAccountParam.UniqueID != other.UniqueID {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReferenceIDForAcquisition: %q,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.ReferenceIDForAcquisition))
	b.WriteString(fmt.Sprintf("%sReferenceIDForRightBinary: %q,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.ReferenceIDForRightBinary))
	b.WriteString(fmt.Sprintf("%sUseType: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.UseType))
	b.WriteString(fmt.Sprintf("%sLimitationType: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.LimitationType))
	b.WriteString(fmt.Sprintf("%sLimitationValue: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.LimitationValue))
	b.WriteString(fmt.Sprintf("%sRightBinary: %x,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.RightBinary))
	b.WriteString(fmt.Sprintf("%sLogMessage: %q,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.LogMessage))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemAcquireServiceItemByAccountParam.UniqueID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAcquireServiceItemByAccountParam returns a new ServiceItemAcquireServiceItemByAccountParam
func NewServiceItemAcquireServiceItemByAccountParam() *ServiceItemAcquireServiceItemByAccountParam {
	return &ServiceItemAcquireServiceItemByAccountParam{}
}
