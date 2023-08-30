// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemAttribute holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemAttribute struct {
	nex.Structure
	Name  string
	Value string
}

// ExtractFromStream extracts a ServiceItemAttribute structure from a stream
func (serviceItemAttribute *ServiceItemAttribute) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemAttribute.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAttribute.Name from stream. %s", err.Error())
	}

	serviceItemAttribute.Value, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAttribute.Value from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemAttribute and returns a byte array
func (serviceItemAttribute *ServiceItemAttribute) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemAttribute.Name)
	stream.WriteString(serviceItemAttribute.Value)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemAttribute
func (serviceItemAttribute *ServiceItemAttribute) Copy() nex.StructureInterface {
	copied := NewServiceItemAttribute()

	copied.SetStructureVersion(serviceItemAttribute.StructureVersion())

	copied.Name = serviceItemAttribute.Name
	copied.Value = serviceItemAttribute.Value

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAttribute *ServiceItemAttribute) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemAttribute)

	if serviceItemAttribute.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemAttribute.Name != other.Name {
		return false
	}

	if serviceItemAttribute.Value != other.Value {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemAttribute *ServiceItemAttribute) String() string {
	return serviceItemAttribute.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemAttribute *ServiceItemAttribute) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAttribute{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemAttribute.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, serviceItemAttribute.Name))
	b.WriteString(fmt.Sprintf("%sValue: %q,\n", indentationValues, serviceItemAttribute.Value))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAttribute returns a new ServiceItemAttribute
func NewServiceItemAttribute() *ServiceItemAttribute {
	return &ServiceItemAttribute{}
}
