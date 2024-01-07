// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemAttribute holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAttribute struct {
	types.Structure
	Name  string
	Value string
}

// ExtractFrom extracts the ServiceItemAttribute from the given readable
func (serviceItemAttribute *ServiceItemAttribute) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemAttribute.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemAttribute header. %s", err.Error())
	}

	err = serviceItemAttribute.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAttribute.Name from stream. %s", err.Error())
	}

	err = serviceItemAttribute.Value.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAttribute.Value from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemAttribute to the given writable
func (serviceItemAttribute *ServiceItemAttribute) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemAttribute.Name.WriteTo(contentWritable)
	serviceItemAttribute.Value.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemAttribute.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemAttribute
func (serviceItemAttribute *ServiceItemAttribute) Copy() types.RVType {
	copied := NewServiceItemAttribute()

	copied.StructureVersion = serviceItemAttribute.StructureVersion

	copied.Name = serviceItemAttribute.Name
	copied.Value = serviceItemAttribute.Value

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAttribute *ServiceItemAttribute) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAttribute); !ok {
		return false
	}

	other := o.(*ServiceItemAttribute)

	if serviceItemAttribute.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemAttribute.Name.Equals(other.Name) {
		return false
	}

	if !serviceItemAttribute.Value.Equals(other.Value) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemAttribute.StructureVersion))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, serviceItemAttribute.Name))
	b.WriteString(fmt.Sprintf("%sValue: %q,\n", indentationValues, serviceItemAttribute.Value))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAttribute returns a new ServiceItemAttribute
func NewServiceItemAttribute() *ServiceItemAttribute {
	return &ServiceItemAttribute{}
}
