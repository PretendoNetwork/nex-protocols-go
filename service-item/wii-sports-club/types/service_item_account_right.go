// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemAccountRight holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemAccountRight struct {
	nex.Structure
	PID        *nex.PID
	Limitation *ServiceItemLimitation
}

// ExtractFromStream extracts a ServiceItemAccountRight structure from a stream
func (serviceItemAccountRight *ServiceItemAccountRight) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemAccountRight.PID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.PID from stream. %s", err.Error())
	}

	limitation, err := stream.ReadStructure(NewServiceItemLimitation())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.Limitation from stream. %s", err.Error())
	}

	serviceItemAccountRight.Limitation = limitation.(*ServiceItemLimitation)

	return nil
}

// Bytes encodes the ServiceItemAccountRight and returns a byte array
func (serviceItemAccountRight *ServiceItemAccountRight) Bytes(stream *nex.StreamOut) []byte {
	stream.WritePID(serviceItemAccountRight.PID)
	stream.WriteStructure(serviceItemAccountRight.Limitation)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemAccountRight
func (serviceItemAccountRight *ServiceItemAccountRight) Copy() nex.StructureInterface {
	copied := NewServiceItemAccountRight()

	copied.SetStructureVersion(serviceItemAccountRight.StructureVersion())

	copied.PID = serviceItemAccountRight.PID.Copy()
	copied.Limitation = serviceItemAccountRight.Limitation.Copy().(*ServiceItemLimitation)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAccountRight *ServiceItemAccountRight) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemAccountRight)

	if serviceItemAccountRight.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !serviceItemAccountRight.PID.Equals(other.PID) {
		return false
	}

	if !serviceItemAccountRight.Limitation.Equals(other.Limitation) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemAccountRight *ServiceItemAccountRight) String() string {
	return serviceItemAccountRight.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemAccountRight *ServiceItemAccountRight) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAccountRight{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemAccountRight.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, serviceItemAccountRight.PID.FormatToString(indentationLevel+1)))

	if serviceItemAccountRight.Limitation != nil {
		b.WriteString(fmt.Sprintf("%sLimitation: %s\n", indentationValues, serviceItemAccountRight.Limitation.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sLimitation: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAccountRight returns a new ServiceItemAccountRight
func NewServiceItemAccountRight() *ServiceItemAccountRight {
	return &ServiceItemAccountRight{}
}
