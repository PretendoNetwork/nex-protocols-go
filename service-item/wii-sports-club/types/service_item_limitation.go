// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemLimitation holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemLimitation struct {
	types.Structure
	LimitationType  *types.PrimitiveU32
	LimitationValue *types.PrimitiveU32
}

// ExtractFrom extracts the ServiceItemLimitation from the given readable
func (serviceItemLimitation *ServiceItemLimitation) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemLimitation.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemLimitation header. %s", err.Error())
	}

	err = serviceItemLimitation.LimitationType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLimitation.LimitationType from stream. %s", err.Error())
	}

	err = serviceItemLimitation.LimitationValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLimitation.LimitationValue from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemLimitation to the given writable
func (serviceItemLimitation *ServiceItemLimitation) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemLimitation.LimitationType.WriteTo(contentWritable)
	serviceItemLimitation.LimitationValue.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemLimitation.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemLimitation
func (serviceItemLimitation *ServiceItemLimitation) Copy() types.RVType {
	copied := NewServiceItemLimitation()

	copied.StructureVersion = serviceItemLimitation.StructureVersion

	copied.LimitationType = serviceItemLimitation.LimitationType
	copied.LimitationValue = serviceItemLimitation.LimitationValue

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemLimitation *ServiceItemLimitation) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemLimitation); !ok {
		return false
	}

	other := o.(*ServiceItemLimitation)

	if serviceItemLimitation.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemLimitation.LimitationType.Equals(other.LimitationType) {
		return false
	}

	if !serviceItemLimitation.LimitationValue.Equals(other.LimitationValue) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemLimitation.StructureVersion))
	b.WriteString(fmt.Sprintf("%sLimitationType: %d,\n", indentationValues, serviceItemLimitation.LimitationType))
	b.WriteString(fmt.Sprintf("%sLimitationValue: %d,\n", indentationValues, serviceItemLimitation.LimitationValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemLimitation returns a new ServiceItemLimitation
func NewServiceItemLimitation() *ServiceItemLimitation {
	return &ServiceItemLimitation{}
}
