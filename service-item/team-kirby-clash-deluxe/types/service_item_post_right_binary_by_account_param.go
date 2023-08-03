// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemPostRightBinaryByAccountParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemPostRightBinaryByAccountParam struct {
	nex.Structure
	ReferenceID string
	UseType     uint8
	RightBinary []byte
	LogMessage  string
	UniqueID    uint32
}

// ExtractFromStream extracts a ServiceItemPostRightBinaryByAccountParam structure from a stream
func (serviceItemPostRightBinaryByAccountParam *ServiceItemPostRightBinaryByAccountParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemPostRightBinaryByAccountParam.ReferenceID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.ReferenceID from stream. %s", err.Error())
	}

	serviceItemPostRightBinaryByAccountParam.UseType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.UseType from stream. %s", err.Error())
	}

	serviceItemPostRightBinaryByAccountParam.RightBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.RightBinary from stream. %s", err.Error())
	}

	serviceItemPostRightBinaryByAccountParam.LogMessage, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.LogMessage from stream. %s", err.Error())
	}

	serviceItemPostRightBinaryByAccountParam.UniqueID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPostRightBinaryByAccountParam.UniqueID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemPostRightBinaryByAccountParam and returns a byte array
func (serviceItemPostRightBinaryByAccountParam *ServiceItemPostRightBinaryByAccountParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemPostRightBinaryByAccountParam.ReferenceID)
	stream.WriteUInt8(serviceItemPostRightBinaryByAccountParam.UseType)
	stream.WriteQBuffer(serviceItemPostRightBinaryByAccountParam.RightBinary)
	stream.WriteString(serviceItemPostRightBinaryByAccountParam.LogMessage)
	stream.WriteUInt32LE(serviceItemPostRightBinaryByAccountParam.UniqueID)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemPostRightBinaryByAccountParam
func (serviceItemPostRightBinaryByAccountParam *ServiceItemPostRightBinaryByAccountParam) Copy() nex.StructureInterface {
	copied := NewServiceItemPostRightBinaryByAccountParam()

	copied.ReferenceID = serviceItemPostRightBinaryByAccountParam.ReferenceID
	copied.UseType = serviceItemPostRightBinaryByAccountParam.UseType
	copied.RightBinary = serviceItemPostRightBinaryByAccountParam.RightBinary
	copied.LogMessage = serviceItemPostRightBinaryByAccountParam.LogMessage
	copied.UniqueID = serviceItemPostRightBinaryByAccountParam.UniqueID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPostRightBinaryByAccountParam *ServiceItemPostRightBinaryByAccountParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemPostRightBinaryByAccountParam)

	if serviceItemPostRightBinaryByAccountParam.ReferenceID != other.ReferenceID {
		return false
	}

	if serviceItemPostRightBinaryByAccountParam.UseType != other.UseType {
		return false
	}

	if !bytes.Equal(serviceItemPostRightBinaryByAccountParam.RightBinary, other.RightBinary) {
		return false
	}

	if serviceItemPostRightBinaryByAccountParam.LogMessage != other.LogMessage {
		return false
	}

	if serviceItemPostRightBinaryByAccountParam.UniqueID != other.UniqueID {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.ReferenceID))
	b.WriteString(fmt.Sprintf("%sUseType: %d,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.UseType))
	b.WriteString(fmt.Sprintf("%sRightBinary: %x,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.RightBinary))
	b.WriteString(fmt.Sprintf("%sLogMessage: %q,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.LogMessage))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemPostRightBinaryByAccountParam.UniqueID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPostRightBinaryByAccountParam returns a new ServiceItemPostRightBinaryByAccountParam
func NewServiceItemPostRightBinaryByAccountParam() *ServiceItemPostRightBinaryByAccountParam {
	return &ServiceItemPostRightBinaryByAccountParam{}
}
