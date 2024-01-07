// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemAccountRight holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAccountRight struct {
	types.Structure
	PID           *types.PID
	Limitation    *ServiceItemLimitation
	RightBinaries []*ServiceItemRightBinary
}

// ExtractFrom extracts the ServiceItemAccountRight from the given readable
func (serviceItemAccountRight *ServiceItemAccountRight) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemAccountRight.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemAccountRight header. %s", err.Error())
	}

	err = serviceItemAccountRight.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.PID from stream. %s", err.Error())
	}

	err = serviceItemAccountRight.Limitation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.Limitation from stream. %s", err.Error())
	}

	rightBinaries, err := nex.StreamReadListStructure(stream, NewServiceItemRightBinary())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.RightBinaries from stream. %s", err.Error())
	}

	serviceItemAccountRight.RightBinaries = rightBinaries

	return nil
}

// WriteTo writes the ServiceItemAccountRight to the given writable
func (serviceItemAccountRight *ServiceItemAccountRight) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemAccountRight.PID.WriteTo(contentWritable)
	serviceItemAccountRight.Limitation.WriteTo(contentWritable)
	serviceItemAccountRight.RightBinaries.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemAccountRight.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemAccountRight
func (serviceItemAccountRight *ServiceItemAccountRight) Copy() types.RVType {
	copied := NewServiceItemAccountRight()

	copied.StructureVersion = serviceItemAccountRight.StructureVersion

	copied.PID = serviceItemAccountRight.PID.Copy()
	copied.Limitation = serviceItemAccountRight.Limitation.Copy().(*ServiceItemLimitation)
	copied.RightBinaries = make([]*ServiceItemRightBinary, len(serviceItemAccountRight.RightBinaries))

	for i := 0; i < len(serviceItemAccountRight.RightBinaries); i++ {
		copied.RightBinaries[i] = serviceItemAccountRight.RightBinaries[i].Copy().(*ServiceItemRightBinary)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAccountRight *ServiceItemAccountRight) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAccountRight); !ok {
		return false
	}

	other := o.(*ServiceItemAccountRight)

	if serviceItemAccountRight.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemAccountRight.PID.Equals(other.PID) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemAccountRight.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, serviceItemAccountRight.PID.FormatToString(indentationLevel+1)))

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
