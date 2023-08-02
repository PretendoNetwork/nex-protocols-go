// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemLimitation holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemLimitation struct {
	nex.Structure
	LimitationType uint32
	LimitationValue uint32
}

// ExtractFromStream extracts a ServiceItemLimitation structure from a stream
func (serviceItemLimitation *ServiceItemLimitation) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemLimitation.LimitationType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLimitation.LimitationType from stream. %s", err.Error())
	}

	serviceItemLimitation.LimitationValue, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLimitation.LimitationValue from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemLimitation and returns a byte array
func (serviceItemLimitation *ServiceItemLimitation) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemLimitation.LimitationType)
	stream.WriteUInt32LE(serviceItemLimitation.LimitationValue)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemLimitation
func (serviceItemLimitation *ServiceItemLimitation) Copy() nex.StructureInterface {
	copied := NewServiceItemLimitation()

	copied.LimitationType = serviceItemLimitation.LimitationType
	copied.LimitationValue = serviceItemLimitation.LimitationValue

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemLimitation *ServiceItemLimitation) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemLimitation)

	if serviceItemLimitation.LimitationType != other.LimitationType {
		return false
	}

	if serviceItemLimitation.LimitationValue != other.LimitationValue {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemLimitation *ServiceItemLimitation) String() string {
	return serviceItemLimitation.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemLimitation *ServiceItemLimitation) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemLimitation{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemLimitation.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sLimitationType: %d,\n", indentationValues, serviceItemLimitation.LimitationType))
	b.WriteString(fmt.Sprintf("%sLimitationValue: %d,\n", indentationValues, serviceItemLimitation.LimitationValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemLimitation returns a new ServiceItemLimitation
func NewServiceItemLimitation() *ServiceItemLimitation {
	return &ServiceItemLimitation{}
}
