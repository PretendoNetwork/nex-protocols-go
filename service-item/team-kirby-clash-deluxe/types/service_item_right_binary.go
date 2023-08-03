// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemRightBinary holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemRightBinary struct {
	nex.Structure
	UseType     uint8
	RightBinary []byte
}

// ExtractFromStream extracts a ServiceItemRightBinary structure from a stream
func (serviceItemRightBinary *ServiceItemRightBinary) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemRightBinary.UseType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightBinary.UseType from stream. %s", err.Error())
	}

	serviceItemRightBinary.RightBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightBinary.RightBinary from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemRightBinary and returns a byte array
func (serviceItemRightBinary *ServiceItemRightBinary) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(serviceItemRightBinary.UseType)
	stream.WriteQBuffer(serviceItemRightBinary.RightBinary)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemRightBinary
func (serviceItemRightBinary *ServiceItemRightBinary) Copy() nex.StructureInterface {
	copied := NewServiceItemRightBinary()

	copied.UseType = serviceItemRightBinary.UseType
	copied.RightBinary = serviceItemRightBinary.RightBinary

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightBinary *ServiceItemRightBinary) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemRightBinary)

	if serviceItemRightBinary.UseType != other.UseType {
		return false
	}

	if !bytes.Equal(serviceItemRightBinary.RightBinary, other.RightBinary) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemRightBinary *ServiceItemRightBinary) String() string {
	return serviceItemRightBinary.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemRightBinary *ServiceItemRightBinary) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightBinary{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemRightBinary.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUseType: %d,\n", indentationValues, serviceItemRightBinary.UseType))
	b.WriteString(fmt.Sprintf("%sRightBinary: %x,\n", indentationValues, serviceItemRightBinary.RightBinary))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightBinary returns a new ServiceItemRightBinary
func NewServiceItemRightBinary() *ServiceItemRightBinary {
	return &ServiceItemRightBinary{}
}
