// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemAccountRight holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAccountRight struct {
	nex.Structure
	PID           uint32
	Limitation    *ServiceItemLimitation
	RightBinaries []*ServiceItemRightBinary
}

// ExtractFromStream extracts a ServiceItemAccountRight structure from a stream
func (serviceItemAccountRight *ServiceItemAccountRight) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemAccountRight.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.PID from stream. %s", err.Error())
	}

	limitation, err := stream.ReadStructure(NewServiceItemLimitation())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.Limitation from stream. %s", err.Error())
	}

	serviceItemAccountRight.Limitation = limitation.(*ServiceItemLimitation)

	rightBinaries, err := stream.ReadListStructure(NewServiceItemRightBinary())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.RightBinaries from stream. %s", err.Error())
	}

	serviceItemAccountRight.RightBinaries = rightBinaries.([]*ServiceItemRightBinary)

	return nil
}

// Bytes encodes the ServiceItemAccountRight and returns a byte array
func (serviceItemAccountRight *ServiceItemAccountRight) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemAccountRight.PID)
	stream.WriteStructure(serviceItemAccountRight.Limitation)
	stream.WriteListStructure(serviceItemAccountRight.RightBinaries)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemAccountRight
func (serviceItemAccountRight *ServiceItemAccountRight) Copy() nex.StructureInterface {
	copied := NewServiceItemAccountRight()

	copied.SetStructureVersion(serviceItemAccountRight.StructureVersion())

	copied.PID = serviceItemAccountRight.PID
	copied.Limitation = serviceItemAccountRight.Limitation.Copy().(*ServiceItemLimitation)
	copied.RightBinaries = make([]*ServiceItemRightBinary, len(serviceItemAccountRight.RightBinaries))

	for i := 0; i < len(serviceItemAccountRight.RightBinaries); i++ {
		copied.RightBinaries[i] = serviceItemAccountRight.RightBinaries[i].Copy().(*ServiceItemRightBinary)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAccountRight *ServiceItemAccountRight) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemAccountRight)

	if serviceItemAccountRight.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemAccountRight.PID != other.PID {
		return false
	}

	if !serviceItemAccountRight.Limitation.Equals(other.Limitation) {
		return false
	}

	if len(serviceItemAccountRight.RightBinaries) != len(other.RightBinaries) {
		return false
	}

	for i := 0; i < len(serviceItemAccountRight.RightBinaries); i++ {
		if !serviceItemAccountRight.RightBinaries[i].Equals(other.RightBinaries[i]) {
			return false
		}
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
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAccountRight{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemAccountRight.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, serviceItemAccountRight.PID))

	if serviceItemAccountRight.Limitation != nil {
		b.WriteString(fmt.Sprintf("%sLimitation: %s\n", indentationValues, serviceItemAccountRight.Limitation.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sLimitation: nil\n", indentationValues))
	}

	if len(serviceItemAccountRight.RightBinaries) == 0 {
		b.WriteString(fmt.Sprintf("%sRightBinaries: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRightBinaries: [\n", indentationValues))

		for i := 0; i < len(serviceItemAccountRight.RightBinaries); i++ {
			str := serviceItemAccountRight.RightBinaries[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemAccountRight.RightBinaries)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAccountRight returns a new ServiceItemAccountRight
func NewServiceItemAccountRight() *ServiceItemAccountRight {
	return &ServiceItemAccountRight{}
}
