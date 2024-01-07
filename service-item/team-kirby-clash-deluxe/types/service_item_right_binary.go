// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemRightBinary holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemRightBinary struct {
	types.Structure
	UseType     *types.PrimitiveU8
	RightBinary []byte
}

// ExtractFrom extracts the ServiceItemRightBinary from the given readable
func (serviceItemRightBinary *ServiceItemRightBinary) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemRightBinary.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemRightBinary header. %s", err.Error())
	}

	err = serviceItemRightBinary.UseType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightBinary.UseType from stream. %s", err.Error())
	}

	serviceItemRightBinary.RightBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightBinary.RightBinary from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemRightBinary to the given writable
func (serviceItemRightBinary *ServiceItemRightBinary) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemRightBinary.UseType.WriteTo(contentWritable)
	stream.WriteQBuffer(serviceItemRightBinary.RightBinary)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemRightBinary
func (serviceItemRightBinary *ServiceItemRightBinary) Copy() types.RVType {
	copied := NewServiceItemRightBinary()

	copied.StructureVersion = serviceItemRightBinary.StructureVersion

	copied.UseType = serviceItemRightBinary.UseType
	copied.RightBinary = serviceItemRightBinary.RightBinary

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightBinary *ServiceItemRightBinary) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightBinary); !ok {
		return false
	}

	other := o.(*ServiceItemRightBinary)

	if serviceItemRightBinary.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemRightBinary.UseType.Equals(other.UseType) {
		return false
	}

	if !serviceItemRightBinary.RightBinary.Equals(other.RightBinary) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemRightBinary.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUseType: %d,\n", indentationValues, serviceItemRightBinary.UseType))
	b.WriteString(fmt.Sprintf("%sRightBinary: %x,\n", indentationValues, serviceItemRightBinary.RightBinary))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightBinary returns a new ServiceItemRightBinary
func NewServiceItemRightBinary() *ServiceItemRightBinary {
	return &ServiceItemRightBinary{}
}
