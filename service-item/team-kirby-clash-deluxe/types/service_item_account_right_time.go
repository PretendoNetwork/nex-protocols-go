// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemAccountRightTime holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAccountRightTime struct {
	types.Structure
	*ServiceItemAccountRight
}

// ExtractFrom extracts the ServiceItemAccountRightTime from the given readable
func (serviceItemAccountRightTime *ServiceItemAccountRightTime) ExtractFrom(readable types.Readable) error {
	return nil
}

// WriteTo writes the ServiceItemAccountRightTime to the given writable
func (serviceItemAccountRightTime *ServiceItemAccountRightTime) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemAccountRightTime
func (serviceItemAccountRightTime *ServiceItemAccountRightTime) Copy() types.RVType {
	copied := NewServiceItemAccountRightTime()

	copied.StructureVersion = serviceItemAccountRightTime.StructureVersion

	copied.ServiceItemAccountRight = serviceItemAccountRightTime.ServiceItemAccountRight.Copy().(*ServiceItemAccountRight)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAccountRightTime *ServiceItemAccountRightTime) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAccountRightTime); !ok {
		return false
	}

	other := o.(*ServiceItemAccountRightTime)

	if serviceItemAccountRightTime.StructureVersion != other.StructureVersion {
		return false
	}

	return serviceItemAccountRightTime.ParentType().Equals(other.ParentType())
}

// String returns a string representation of the struct
func (serviceItemAccountRightTime *ServiceItemAccountRightTime) String() string {
	return serviceItemAccountRightTime.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemAccountRightTime *ServiceItemAccountRightTime) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAccountRightTime{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemAccountRightTime.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemAccountRightTime.StructureVersion))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAccountRightTime returns a new ServiceItemAccountRightTime
func NewServiceItemAccountRightTime() *ServiceItemAccountRightTime {
	serviceItemAccountRightTime := &ServiceItemAccountRightTime{}

	serviceItemAccountRightTime.ServiceItemAccountRight = NewServiceItemAccountRight()
	serviceItemAccountRightTime.SetParentType(serviceItemAccountRightTime.ServiceItemAccountRight)

	return serviceItemAccountRightTime
}
