// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemAccountRightTime holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAccountRightTime struct {
	nex.Structure
	*ServiceItemAccountRight
}

// ExtractFromStream extracts a ServiceItemAccountRightTime structure from a stream
func (serviceItemAccountRightTime *ServiceItemAccountRightTime) ExtractFromStream(stream *nex.StreamIn) error {
	return nil
}

// Bytes encodes the ServiceItemAccountRightTime and returns a byte array
func (serviceItemAccountRightTime *ServiceItemAccountRightTime) Bytes(stream *nex.StreamOut) []byte {
	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemAccountRightTime
func (serviceItemAccountRightTime *ServiceItemAccountRightTime) Copy() nex.StructureInterface {
	copied := NewServiceItemAccountRightTime()

	copied.ServiceItemAccountRight = serviceItemAccountRightTime.ServiceItemAccountRight.Copy().(*ServiceItemAccountRight)
	copied.SetParentType(copied.ServiceItemAccountRight)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAccountRightTime *ServiceItemAccountRightTime) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemAccountRightTime)

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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemAccountRightTime.StructureVersion()))
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
